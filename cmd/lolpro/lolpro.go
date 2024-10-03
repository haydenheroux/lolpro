package main

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/haydenheroux/lolpro/pkg/model"
	"github.com/urfave/cli/v2"
)

const databaseFile = "test.db"

var db *gorm.DB

func main() {
	var err error

	db, err = gorm.Open(sqlite.Open(databaseFile), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database")
	}

	db.AutoMigrate(&model.Team{})
	db.AutoMigrate(&model.Player{})
	db.AutoMigrate(&model.PlayerMatchData{})

	db.AutoMigrate(&model.Match{})

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

	db.Create(&test)

	return nil
}

func createPlayer(c *cli.Context) error {
	test := model.Player{
		Name:      "Mari",
		Residency: model.NorthAmerica,
	}

	db.Create(&test)

	return nil
}

func createMatch(c *cli.Context) error {
	test := model.Match{
		Duration: time.Duration(7*time.Minute + 20*time.Second),
	}

	db.Create(&test)

	return nil
}
