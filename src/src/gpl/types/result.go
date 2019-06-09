package gql

import "github.com/graphql-go/graphql"

//Properties ... Get properties of a team
var Properties = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Team",
		Fields: graphql.Fields{
			"Last10WinningLocalMatchs": &graphql.Field{
				Type: graphql.Int,
			},
			"Last10TiedingLocalMatchs": &graphql.Field{
				Type: graphql.Int,
			},
			"Last10LosingLocalMatchs": &graphql.Field{
				Type: graphql.Int,
			},
			"Last10WinningAwayMatchs": &graphql.Field{
				Type: graphql.Int,
			},
			"Last10TiedingAwayMatchs": &graphql.Field{
				Type: graphql.Int,
			},
			"Last10LosingAwayMatchs": &graphql.Field{
				Type: graphql.Int,
			},
			"Last10StreackWinningLocal": &graphql.Field{
				Type: graphql.Int,
			},
			"Last10StreackNoLosingLocal": &graphql.Field{
				Type: graphql.Int,
			},
			"Last10StreackWinningAway": &graphql.Field{
				Type: graphql.Int,
			},
			"Last10StreackNoLosingAway": &graphql.Field{
				Type: graphql.Int,
			},
			"Last10AverageTuckedGoalsLocal": &graphql.Field{
				Type: graphql.Float,
			},
			"Last10AverageReceivedGoalsLocal": &graphql.Field{
				Type: graphql.Float,
			},
			"Last10AverageTuckedGoalsAway": &graphql.Field{
				Type: graphql.Float,
			},
			"Last10AverageReceivedGoalsAway": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)
