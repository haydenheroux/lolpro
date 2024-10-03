package model

import (
	"time"

	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	Name    string
	Region  Region
	Players []Player
}

type Player struct {
	gorm.Model
	Name      string
	Residency Region
	Role      Role
	TeamID    uint
}

type Match struct {
	gorm.Model
	BlueTeamID   uint
	BlueTeamData []PlayerMatchData
	RedTeamID    uint
	RedTeamData  []PlayerMatchData
	Duration     time.Duration
}

type PlayerMatchData struct {
	gorm.Model
	PlayerID           uint
	MatchID            uint
	Kills              uint
	Deaths             uint
	Assists            uint
	DamageDealt        uint
	GoldEarned         uint
	CreepScore         uint
	LaneGoldDifference uint
	SoloKills          uint
}
