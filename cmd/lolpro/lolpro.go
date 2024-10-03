package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
					{
						Name:   "data",
						Usage:  "create a match data entry",
						Action: createMatchData,
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func askString(prompt string) string {
	var response string

	fmt.Print(prompt)
	fmt.Scanln(&response)

	response = strings.TrimSpace(response)

	return response
}

func askInt(prompt string) int {
	response := askString(prompt)

	n, _ := strconv.Atoi(response)

	return n
}

func askTeam() *model.Team {
	// TODO
	teams, _ := db.GetTeams()

	var sb strings.Builder

	for index, team := range teams {
		sb.WriteString(fmt.Sprintf("%d: %s\n", index, team.Name))
	}

	sb.WriteString("Team: ")

	prompt := sb.String()

	index := askInt(prompt)

	return teams[index]
}

func askRegion() model.Region {
	var sb strings.Builder

	for index, region := range model.Regions {
		sb.WriteString(fmt.Sprintf("%d: %s\n", index, region))
	}

	sb.WriteString("Region: ")

	prompt := sb.String()

	index := askInt(prompt)

	return model.Regions[index]
}

func askDuration() time.Duration {
	response := askString("Duration: ")

	parts := strings.Split(response, ":")

	minutes, _ := strconv.Atoi(parts[0])
	seconds, _ := strconv.Atoi(parts[1])

	return time.Duration(minutes*int(time.Minute) + seconds*int(time.Second))
}

func askMatch() *model.Match {
	// TODO
	matches, _ := db.GetMatches()

	var sb strings.Builder

	for index, match := range matches {
		// TOOD Hacky workaround to fiends being zero-valued
		blue, _ := db.GetTeam(match.BlueTeamID)
		red, _ := db.GetTeam(match.RedTeamID)

		sb.WriteString(fmt.Sprintf("%d: %s vs %s\n", index, blue.Name, red.Name))
	}

	sb.WriteString("Match: ")

	prompt := sb.String()

	index := askInt(prompt)

	return matches[index]
}

func askPlayer(match *model.Match) *model.Player {
	// TOOD Hacky workaround to fiends being zero-valued
	blue, _ := db.GetTeam(match.BlueTeamID)
	red, _ := db.GetTeam(match.RedTeamID)

	players := make([]*model.Player, 0)

	for _, player := range blue.Players {
		players = append(players, &player)
	}

	for _, player := range red.Players {
		players = append(players, &player)
	}

	var sb strings.Builder

	for index, player := range players {
		sb.WriteString(fmt.Sprintf("%d: %s\n", index, player.Name))
	}

	sb.WriteString("Player: ")

	prompt := sb.String()

	index := askInt(prompt)

	return players[index]
}

func createTeam(c *cli.Context) error {
	teamName := askString("Team.Name: ")
	region := askRegion()

	test := model.Team{
		Name:    teamName,
		Region:  region,
		Players: []model.Player{},
	}

	db.SaveTeam(&test)

	return nil
}

func createPlayer(c *cli.Context) error {
	team := askTeam()

	playerName := askString("Player.Name: ")
	residency := askRegion()

	player := model.Player{
		Name:      playerName,
		Residency: residency,
	}

	db.SavePlayer(&player)

	team.Players = append(team.Players, player)

	db.SaveTeam(team)

	return nil
}

func createMatch(c *cli.Context) error {
	blueTeam := askTeam()
	redTeam := askTeam()

	duration := askDuration()

	match := model.Match{
		BlueTeam: *blueTeam,
		RedTeam:  *redTeam,
		Duration: duration,
	}

	db.SaveMatch(&match)

	return nil
}

func createMatchData(c *cli.Context) error {
	match := askMatch()
	player := askPlayer(match)

	kills := askInt("Kills: ")
	deaths := askInt("Deaths: ")
	assists := askInt("Assists: ")
	damageDealt := askInt("DamageDealt: ")
	goldEarned := askInt("GoldEarned: ")
	creepScore := askInt("CreepScore: ")
	laneGoldDifference := askInt("LaneGoldDifference: ")
	soloKills := askInt("SoloKills: ")

	matchData := model.PlayerMatchData{
		Player:             *player,
		Match:              *match,
		Kills:              uint(kills),
		Deaths:             uint(deaths),
		Assists:            uint(assists),
		DamageDealt:        uint(damageDealt),
		GoldEarned:         uint(goldEarned),
		CreepScore:         uint(creepScore),
		LaneGoldDifference: uint(laneGoldDifference),
		SoloKills:          uint(soloKills),
	}

	db.SaveMatchData(&matchData)

	return nil
}
