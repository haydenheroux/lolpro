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

func (d Database) SavePlayer(player *model.Player) error {
	return d.Database.Save(player).Error
}

func (d Database) SaveMatch(match *model.Match) error {
	return d.Database.Save(match).Error
}
