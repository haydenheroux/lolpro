package tui

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/haydenheroux/lolpro/pkg/model"
)

func AskString(prompt string) string {
	var response string

	fmt.Print(prompt)
	fmt.Scanln(&response)

	response = strings.TrimSpace(response)

	return response
}

func AskInt(prompt string) int {
	response := AskString(prompt)

	n, _ := strconv.Atoi(response)

	return n
}

func PickTeam(teams []*model.Team) *model.Team {
	var promptBuilder strings.Builder

	for index, team := range teams {
		promptBuilder.WriteString(fmt.Sprintf("%d: %s\n", index, team.Name))
	}

	promptBuilder.WriteString("Team: ")

	index := AskInt(promptBuilder.String())

	return teams[index]
}

func PickRegion() model.Region {
	var promptBuilder strings.Builder

	for index, region := range model.Regions {
		promptBuilder.WriteString(fmt.Sprintf("%d: %s\n", index, region))
	}

	promptBuilder.WriteString("Region: ")

	index := AskInt(promptBuilder.String())

	return model.Regions[index]
}

func AskDuration() time.Duration {
	response := AskString("Duration: ")

	parts := strings.Split(response, ":")

	minutes, _ := strconv.Atoi(parts[0])
	seconds, _ := strconv.Atoi(parts[1])

	return time.Duration(minutes*int(time.Minute) + seconds*int(time.Second))
}

func PickMatch(matches []*model.Match) *model.Match {
	var promptBuilder strings.Builder

	for index, match := range matches {
		promptBuilder.WriteString(fmt.Sprintf("%d: %s vs %s\n", index, match.BlueTeam.Name, match.RedTeam.Name))
	}

	promptBuilder.WriteString("Match: ")

	index := AskInt(promptBuilder.String())

	return matches[index]
}

func PickPlayer(players []*model.Player) *model.Player {
	var promptBuilder strings.Builder

	for index, player := range players {
		promptBuilder.WriteString(fmt.Sprintf("%d: %s\n", index, player.Name))
	}

	promptBuilder.WriteString("Player: ")

	index := AskInt(promptBuilder.String())

	return players[index]
}

func PickWinnerLoser(blue, red *model.Team) (*model.Team, *model.Team) {
	var promptBuilder strings.Builder

	promptBuilder.WriteString(fmt.Sprintf("%d: %s\n", 0, blue.Name))
	promptBuilder.WriteString(fmt.Sprintf("%d: %s\n", 1, red.Name))

	promptBuilder.WriteString("Winner: ")

	index := AskInt(promptBuilder.String())

	if index == 0 {
		return blue, red
	} else {
		return red, blue
	}
}
