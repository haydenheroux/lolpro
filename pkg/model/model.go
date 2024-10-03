package model

import (
	"time"

	"gorm.io/gorm"
)

// Team represents a group of players competing in League of Legends esports.
type Team struct {
	gorm.Model
	Name    string
	Region  Region
	Players []Player
}

// Player represents an individual player competing in League of Legends esports.
type Player struct {
	gorm.Model
	Name      string
	Residency Region
	Role      Role
	TeamID    uint
}

// Match represents a match played between competing teams.
type Match struct {
	gorm.Model
	BlueTeamID    uint
	BlueTeam      Team
	BlueTeamData  []PlayerMatchData
	RedTeamID     uint
	RedTeam       Team
	RedTeamData   []PlayerMatchData
	WinningTeamID uint
	WinningTeam   Team
	LosingTeamID  uint
	LosingTeam    Team
	Duration      time.Duration
}

// PlayerMatchData represents data relating to a player in a specific match.
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
