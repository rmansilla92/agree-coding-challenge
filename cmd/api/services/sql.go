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
		GetCardFromDB(cardID string) (*domain.CardEntity, error)
		CreateCardIntoDB(cardEntity *domain.NewCardEntity) error
		UpdateCardIntoDB(cardEntity *domain.NewCardEntity) error
		DeleteCardFromDB(cardID string) error
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

func (sqls *sqlServices) GetCardFromDB(cardID string) (*domain.CardEntity, error) {
	var cardEntity domain.CardEntity
	err := sqls.db.Table("cards").Select("cards.*, " +
		"subtypes.id as subtype_id, subtypes.description as subtype_description, " +
		"images.id as image_id, images.description as image_description, " +
		"types.id  as type_id, types.description as type_description").
		Joins("left join subtypes on cards.subtype_id=subtypes.id").
		Joins("left join types on subtypes.type_id=types.id").
		Joins("left join images on cards.subtype_id=images.id").
		First(&cardEntity, "cards.id = ?", cardID).Error
	if err != nil {
		return nil, err
	}
	return &cardEntity, nil
}

func (sqls *sqlServices) CreateCardIntoDB(cardEntity *domain.NewCardEntity) error {
	err := sqls.db.Create(&cardEntity).Error
	if err != nil {
		return err
	}
	return nil
}

func (sqls *sqlServices) UpdateCardIntoDB(cardEntity *domain.NewCardEntity) error {
	err := sqls.db.Model(cardEntity).Where("id = ?", cardEntity.ID).UpdateColumns(cardEntity).Error
	if err != nil {
		return err
	}
	return nil
}

func (sqls *sqlServices) DeleteCardFromDB(cardID string) error {
	err := sqls.db.Exec("DELETE FROM cards WHERE id = ?", cardID).Error
	if err != nil {
		return err
	}
	return nil
}