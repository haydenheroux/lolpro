package database

import (
	"github.com/haydenheroux/lolpro/pkg/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func Create(dsn string) (*Database, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		return &Database{}, err
	}

	// TODO Refactor to AutoMigrate([]interface{})
	db.AutoMigrate(&model.Team{})
	db.AutoMigrate(&model.Player{})
	db.AutoMigrate(&model.PlayerMatchData{})
	db.AutoMigrate(&model.Match{})

	return &Database{db}, nil
}

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

func (d Database) GetTeam(id uint) (*model.Team, error) {
	var team model.Team

	if err := d.db.Preload("Players").First(&team, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &team, nil
}

func (d Database) SavePlayer(player *model.Player) error {
	return d.db.Save(player).Error
}

func (d Database) SaveMatch(match *model.Match) error {
	return d.db.Save(match).Error
}

func (d Database) GetMatches() ([]*model.Match, error) {
	var matches []*model.Match

	if err := d.db.Find(&matches).Error; err != nil {
		return nil, err
	}

	return matches, nil
}

func (d Database) SaveMatchData(matchData *model.PlayerMatchData) error {
	return d.db.Save(matchData).Error
}
