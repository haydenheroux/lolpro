package database

import (
	"github.com/haydenheroux/lolpro/pkg/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Database is a client using the database.
type Database struct {
	db *gorm.DB
}

// Create creates a client to the database pointed to by the name.
func Create(name string) (*Database, error) {
	db, err := gorm.Open(sqlite.Open(name), &gorm.Config{})

	if err != nil {
		return &Database{}, err
	}

	// TODO Potentially migrate to AutoMigrate([]interface{})
	db.AutoMigrate(&model.Team{})
	db.AutoMigrate(&model.Player{})
	db.AutoMigrate(&model.Match{})
	db.AutoMigrate(&model.PlayerMatchData{})

	return &Database{db}, nil
}

// SaveTeam saves the team to the database.
func (d Database) SaveTeam(team *model.Team) error {
	return d.db.Save(team).Error
}

func (d Database) GetTeams() ([]*model.Team, error) {
	var teams []*model.Team

	if err := d.db.Find(&teams).Error; err != nil {
		return nil, err
	}

	return teams, nil
}

// GetTeam retrieves the team matching the ID from the database.
func (d Database) GetTeam(id uint) (*model.Team, error) {
	var team model.Team

	if err := d.db.Preload("Players").First(&team, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &team, nil
}

// SavePlayer saves the player to the database.
func (d Database) SavePlayer(player *model.Player) error {
	return d.db.Save(player).Error
}

// GetPlayers retrieves all players that participated in the match from the database.
func (d Database) GetPlayers(match *model.Match) ([]*model.Player, error) {
	blue, _ := d.GetTeam(match.BlueTeamID)
	red, _ := d.GetTeam(match.RedTeamID)

	players := make([]*model.Player, 0)
	for _, player := range blue.Players {
		players = append(players, &player)
	}
	for _, player := range red.Players {
		players = append(players, &player)
	}

	return players, nil
}

// SaveMatch saves the match to the database.
func (d Database) SaveMatch(match *model.Match) error {
	return d.db.Save(match).Error
}

// GetMatches retrieves all matches from database.
func (d Database) GetMatches() ([]*model.Match, error) {
	var matches []*model.Match

	if err := d.db.Preload("BlueTeam").Preload("RedTeam").Preload("WinningTeam").Preload("LosingTeam").Find(&matches).Error; err != nil {
		return nil, err
	}

	return matches, nil
}

// SaveMatchData saves the match data to the database.
func (d Database) SaveMatchData(matchData *model.PlayerMatchData) error {
	return d.db.Save(matchData).Error
}
