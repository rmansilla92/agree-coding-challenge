package services

import (
	"rmansilla92/agree-coding-challenge/cmd/api/config"
	"rmansilla92/agree-coding-challenge/cmd/api/domain"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type (
	SQLServices interface {
		GetCardsFromDB() ([]domain.CardEntity, error)
	}
	sqlServices struct {
		db *gorm.DB
	}
)

func NewSQLServices(db *gorm.DB) SQLServices {
	return &sqlServices{db}
}

func OpenDB(optionalConnectionString string) (*gorm.DB, error) {
	var db *gorm.DB
	var errOpen error
	if !strings.EqualFold(optionalConnectionString, "default") {
		db, errOpen = gorm.Open(mysql.Open(optionalConnectionString), &gorm.Config{})
	} else {
		db, errOpen = gorm.Open(mysql.Open(config.ConnectionString), &gorm.Config{}) 
	}
	if errOpen != nil {
		db.Logger.LogMode(logger.Error)
		return nil, errOpen
	}
	db.Logger.LogMode(logger.Info)
	return db, nil
}

func (sqls *sqlServices) GetCardsFromDB() ([]domain.CardEntity, error) {
	var cards []domain.CardEntity
	err := sqls.db.Table("cards").Select("cards.*, " +
		"subtypes.id as subtype_id, subtypes.description as subtype_description, " +
		"images.id as image_id, images.description as image_description, " +
		"types.id  as type_id, types.description as type_description").
		Joins("left join subtypes on cards.subtype_id=subtypes.id").
		Joins("left join types on subtypes.type_id=types.id").
		Joins("left join images on cards.subtype_id=images.id").
		Find(&cards).Error
	if err != nil {
		return nil, err
	}
	return cards, nil
}
