package database

import (
	"github.com/haydenheroux/lolpro/pkg/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	Database *gorm.DB
}

func Create(dsn string) (*Database, error) {
	database, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		return &Database{}, err
	}

	// TODO Refactor to AutoMigrate([]interface{})
	database.AutoMigrate(&model.Team{})
	database.AutoMigrate(&model.Player{})
	database.AutoMigrate(&model.PlayerMatchData{})
	database.AutoMigrate(&model.Match{})

	return &Database{Database: database}, nil
}

func (d Database) SaveTeam(team *model.Team) error {
	return d.Database.Save(team).Error
}

func (d Database) GetTeams() ([]*model.Team, error) {
	var teams []*model.Team

	if err := d.Database.Find(&teams).Error; err != nil {
		return nil, err
	}

	return teams, nil
}

func (d Database) GetTeam(id uint) (*model.Team, error) {
	var team model.Team

	if err := d.Database.Preload("Players").First(&team, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &team, nil
}

func (d Database) SavePlayer(player *model.Player) error {
	return d.Database.Save(player).Error
}

func (d Database) SaveMatch(match *model.Match) error {
	return d.Database.Save(match).Error
}

func (d Database) GetMatches() ([]*model.Match, error) {
	var matches []*model.Match

	if err := d.Database.Find(&matches).Error; err != nil {
		return nil, err
	}

	return matches, nil
}

func (d Database) SaveMatchData(matchData *model.PlayerMatchData) error {
	return d.Database.Save(matchData).Error
}
