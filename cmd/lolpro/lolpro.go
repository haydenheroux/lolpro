package main

import (
	"log"
	"os"

	"github.com/haydenheroux/lolpro/pkg/database"
	"github.com/haydenheroux/lolpro/pkg/model"
	"github.com/haydenheroux/lolpro/pkg/tui"
	"github.com/urfave/cli/v2"
)

var db *database.Database

func main() {
	var err error

	dsn := os.Getenv("DB")

	if len(dsn) == 0 {
		dsn = "test.db"
	}

	db, err = database.Create(dsn)
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
func createTeam(c *cli.Context) error {
	teamName := tui.AskString("Team Name?", "")
	region := tui.PickRegion("Team Region?")

	test := model.Team{
		Name:    teamName,
		Region:  region,
		Players: []model.Player{},
	}

	db.SaveTeam(&test)

	return nil
}

func createPlayer(c *cli.Context) error {
	teams, _ := db.GetTeams()

	team := tui.PickTeam("Team?", teams)

	playerName := tui.AskString("Player Name?", "")
	residency := tui.PickRegion("Player Residency?")

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
	teams, _ := db.GetTeams()

	blueTeam := tui.PickTeam("Blue Team?", teams)

	// Filter out the blue team
	otherTeams := make([]*model.Team, 0, len(teams)-1)
	for _, team := range teams {
		if team.ID != blueTeam.ID {
			otherTeams = append(otherTeams, team)
		}
	}

	redTeam := tui.PickTeam("Red Team?", otherTeams)

	winner, loser := tui.PickWinnerLoser(blueTeam, redTeam)

	duration := tui.AskDuration()

	match := model.Match{
		BlueTeam:    *blueTeam,
		RedTeam:     *redTeam,
		WinningTeam: *winner,
		LosingTeam:  *loser,
		Duration:    duration,
	}

	db.SaveMatch(&match)

	return nil
}

func createMatchData(c *cli.Context) error {
	// TODO
	matches, _ := db.GetMatches()

	match := tui.PickMatch("Match?", matches)

	players := make([]*model.Player, 0)
	for _, player := range match.BlueTeam.Players {
		players = append(players, &player)
	}
	for _, player := range match.RedTeam.Players {
		players = append(players, &player)
	}

	player := tui.PickPlayer("Player?", players)

	kills := tui.AskInt("Kills: ", "")
	deaths := tui.AskInt("Deaths: ", "")
	assists := tui.AskInt("Assists: ", "")
	damageDealt := tui.AskInt("DamageDealt: ", "")
	goldEarned := tui.AskInt("GoldEarned: ", "")
	creepScore := tui.AskInt("CreepScore: ", "")
	laneGoldDifference := tui.AskInt("LaneGoldDifference: ", "")
	soloKills := tui.AskInt("SoloKills: ", "")

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
