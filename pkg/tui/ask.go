package tui

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/haydenheroux/lolpro/pkg/model"

	"github.com/charmbracelet/bubbles/list"
)

func AskString(prompt, placeholder string) string {
	return ask(prompt, placeholder)
}

func AskInt(prompt, placeholder string) int {
	response := AskString(prompt, placeholder)

	n, _ := strconv.Atoi(response)

	return n
}

type teamItem struct {
	*model.Team
}

func (item teamItem) String() string { return item.Team.Name }

func (item teamItem) FilterValue() string { return "" }

func PickTeam(prompt string, teams []*model.Team) *model.Team {
	items := make([]list.Item, 0, len(teams))
	for _, team := range teams {
		items = append(items, teamItem{team})
	}

	item := pick(prompt, items)
	return item.(teamItem).Team
}

type regionItem struct {
	model.Region
}

func (item regionItem) String() string { return string(item.Region) }

func (item regionItem) FilterValue() string { return "" }

func PickRegion(prompt string) model.Region {
	items := make([]list.Item, 0, len(model.Regions))
	for _, region := range model.Regions {
		items = append(items, regionItem{region})
	}

	item := pick(prompt, items)
	return item.(regionItem).Region
}

func AskDuration() time.Duration {
	response := AskString("Duration?", "")

	parts := strings.Split(response, ":")

	minutes, _ := strconv.Atoi(parts[0])
	seconds, _ := strconv.Atoi(parts[1])

	return time.Duration(minutes*int(time.Minute) + seconds*int(time.Second))
}

type matchItem struct {
	*model.Match
}

func (item matchItem) String() string {
	return fmt.Sprintf("%s vs. %s", item.Match.BlueTeam.Name, item.Match.RedTeam.Name)
}

func (item matchItem) FilterValue() string { return "" }

func PickMatch(prompt string, matches []*model.Match) *model.Match {
	items := make([]list.Item, 0, len(matches))
	for _, match := range matches {
		items = append(items, matchItem{match})
	}

	item := pick(prompt, items)
	return item.(matchItem).Match
}

type playerItem struct {
	*model.Player
}

func (item playerItem) String() string { return item.Player.Name }

func (item playerItem) FilterValue() string { return "" }

func PickPlayer(prompt string, players []*model.Player) *model.Player {
	items := make([]list.Item, 0, len(players))
	for _, player := range players {
		items = append(items, playerItem{player})
	}

	item := pick(prompt, items)
	return item.(playerItem).Player
}

func PickWinnerLoser(blue, red *model.Team) (*model.Team, *model.Team) {
	teams := []*model.Team{blue, red}

	team := PickTeam("Winner?", teams)

	if team.ID == teams[0].ID {
		return blue, red
	} else {
		return red, blue
	}
}
