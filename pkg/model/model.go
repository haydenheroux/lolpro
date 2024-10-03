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
	BlueTeam     Team
	BlueTeamData []PlayerMatchData
	RedTeamID    uint
	RedTeam      Team
	RedTeamData  []PlayerMatchData
	Duration     time.Duration
}

type PlayerMatchData struct {
	gorm.Model
	PlayerID           uint
	Player             Player
	MatchID            uint
	Match              Match
	Kills              uint
	Deaths             uint
	Assists            uint
	DamageDealt        uint
	GoldEarned         uint
	CreepScore         uint
	LaneGoldDifference uint
	SoloKills          uint
}
