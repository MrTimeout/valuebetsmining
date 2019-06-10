package ios

import (
	"github.com/txn2/txeh"
)

//AddHost ... Add host to the /etc/hosts file
func AddHost(ip, name string) error {
	hosts, err := txeh.NewHostsDefault()
	if err != nil {
		return err
	}
	hosts.AddHost(ip, name)
	hosts.Save()
	return nil
}
