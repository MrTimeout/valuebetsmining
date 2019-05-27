package data

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//Connection ... Struct used to connect to an url and download it
type Connection struct {
	Config
}

//Get ... Download the content of an endpoint
func (c Connection) Get() (string, error) {
	res, err := http.Get(fmt.Sprintf("%s/%s/%s.csv", c.Path, c.Year.GetYears()[0], c.Endpoint[0].Keys[0]))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	fmt.Printf("Body: %s", body)
	return string(body), nil
}

//GetAllByCountryDiv ... Download all the content of all years of a country/div
func (c Connection) GetAllByCountryDiv(country, div string) ([]string, error) {
	if resul, err := c.ExistsCountry(country); err != nil {
		return []string{}, err
	} else if !resul {
		return []string{}, errors.New("Errors getting country")
	}
	if resul, err := c.ExistsDivision(div); err != nil {
		return []string{}, err
	} else if !resul {
		return []string{}, errors.New("Errors getting division")
	}
	result := []string{}
	for _, value := range c.Year.GetYears() {
		res, err := http.Get(fmt.Sprintf("%s/%s/%s.csv", c.Path, value, div))
		if err != nil {
			return []string{}, err
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return []string{}, err
		}
		result = append(result, string(body))
	}
	return result, nil
}

//WriteWithParams ... Return the response of a get request
func (c Connection) WriteWithParams(year Year, country, div string) (string, error) {
	res, err := http.Get(fmt.Sprintf("%s/%d%d/%s.csv", c.Path, year.From, year.To, div))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(fmt.Sprintf("./leagues/%s/%s_%d%d.csv", country, div, year.From, year.To), body, 0644)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("./leagues/%s_%s_%d%d", country, div, c.Year.From, c.Year.To), nil
}

//RequestByCountryDivYear ... Download all the content of each year of a country/div and write it in a file with the name {{.Country}}_{{.Division}}.csv
func (c Connection) RequestByCountryDivYear(year Year, country, div string) (string, error) {
	if resul, err := c.ExistsCountry(country); err != nil {
		return "", err
	} else if !resul {
		return "", errors.New("Errors getting country")
	}
	if resul, err := c.ExistsDivision(div); err != nil {
		return "", err
	} else if !resul {
		return "", errors.New("Errors getting division")
	}
	return c.WriteWithParams(year, country, div)
}

//WriteByCountryDivYears ... Write in csv files by a range of years, country and division
func (c Connection) WriteByCountryDivYears(year Year, country, div string) error {
	if resul, err := c.ExistsCountry(country); err != nil {
		return err
	} else if !resul {
		return errors.New("Errors getting country")
	}
	if resul, err := c.ExistsDivision(div); err != nil {
		return err
	} else if !resul {
		return errors.New("Errors getting division")
	}
	err := os.MkdirAll(fmt.Sprintf("./leagues/%s", country), 0777)
	if err != nil {
		return err
	}
	for i := year.From; i < year.To; i++ {
		_, err := c.RequestByCountryDivYear(Year{From: i, To: i + 1}, country, div)
		if err != nil {
			return err
		}
	}
	return nil
}

//WriteAllByCountryDiv ... Download all the content of all years of a country/div and write it in a file with the name {{.Country}}_{{.Division}}.csv
func (c Connection) WriteAllByCountryDiv(country, div string) (string, error) {
	if resul, err := c.ExistsCountry(country); err != nil {
		return "", err
	} else if !resul {
		return "", errors.New("Errors getting country")
	}
	if resul, err := c.ExistsDivision(div); err != nil {
		return "", err
	} else if !resul {
		return "", errors.New("Errors getting division")
	}
	i := 0
	for _, value := range c.Year.GetYears() {
		res, err := http.Get(fmt.Sprintf("%s/%s/%s.csv", c.Path, value, div))
		if err != nil {
			return "", err
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", err
		}
		if i == 0 {
			err = ioutil.WriteFile(fmt.Sprintf("./leagues/%s_%s_%d%d.csv", country, div, c.Year.From, c.Year.To), body, 0644)
			if err != nil {
				return "", err
			}
		} else {
			f, err := os.OpenFile(fmt.Sprintf("./leagues/%s_%s_%d%d.csv", country, div, c.Year.From, c.Year.To), os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				return "", err
			}

			defer f.Close()

			if _, err = f.WriteString(string(body)); err != nil {
				return "", err
			}
		}
		i++
	}
	return fmt.Sprintf("./leagues/%s_%s_%d%d", country, div, c.Year.From, c.Year.To), nil
}
