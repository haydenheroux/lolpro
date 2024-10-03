package main

import (
	"log"
	"os"
	"time"

	"github.com/haydenheroux/lolpro/pkg/database"
	"github.com/haydenheroux/lolpro/pkg/model"
	"github.com/urfave/cli/v2"
)

const databaseFile = "test.db"

var db *database.Database

func main() {
	var err error

	db, err = database.Create(databaseFile)
	if err != nil {
		log.Fatal(err)
	}

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "create",
				Usage: "create database object",
				Subcommands: []*cli.Command{
					{
						Name:   "team",
						Usage:  "create a team",
						Action: createTeam,
					},
					{
						Name:   "player",
						Usage:  "create a player",
						Action: createPlayer,
					},
					{
						Name:   "match",
						Usage:  "create a match",
						Action: createMatch,
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func createTeam(c *cli.Context) error {
	test := model.Team{
		Name:    "Marya",
		Region:  model.NorthAmerica,
		Players: []model.Player{},
	}

	db.SaveTeam(&test)

	return nil
}

func createPlayer(c *cli.Context) error {
	test := model.Player{
		Name:      "Mari",
		Residency: model.NorthAmerica,
	}

	db.SavePlayer(&test)

	return nil
}

func createMatch(c *cli.Context) error {
	test := model.Match{
		Duration: time.Duration(7*time.Minute + 20*time.Second),
	}

	db.SaveMatch(&test)

	return nil
}
