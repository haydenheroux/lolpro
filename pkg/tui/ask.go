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
	response := ""

	for len(response) == 0 {
		response = strings.TrimSpace(ask(prompt, placeholder))
	}

	return response
}

func AskInt(prompt, placeholder string) int {
	n := 0

	var err error
	ok := false

	for !ok {
		response := AskString(prompt, placeholder)
		n, err = strconv.Atoi(response)

		ok = err == nil
	}

	return n
}

func AskDuration() time.Duration {
	var minutes, seconds int
	var err error

	ok := false

	for !ok {
		response := AskString("Duration?", "XX:XX")

		if !strings.Contains(response, ":") {
			continue
		}

		parts := strings.Split(response, ":")

		if len(parts) != 2 {
			continue
		}

		minutes, err = strconv.Atoi(parts[0])
		if err != nil {
			continue
		}

		seconds, err = strconv.Atoi(parts[1])
		if err != nil {
			continue
		}

		ok = true
	}

	return time.Duration(minutes*int(time.Minute) + seconds*int(time.Second))
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
