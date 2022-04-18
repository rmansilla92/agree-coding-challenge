package domain

type TypeEntity struct {
	ID          *int64  `json:"id" gorm:"column:id;primaryKey"`
	Description *string `json:"description" gorm:"column:description"`
}

func (TypeEntity) TableName() string {
	return "types"
}

type SubtypeEntity struct {
	ID          *int64  `json:"id" gorm:"column:id;primaryKey"`
	Description *string `json:"description" gorm:"column:description"`
}

func (SubtypeEntity) TableName() string {
	return "subtypes"
}

type ImageEntity struct {
	ID          *int64  `json:"id" gorm:"column:id;primaryKey"`
	Description *string `json:"description" gorm:"column:description"`
	URL         *string `json:"url" gorm:"column:url"`
}

func (ImageEntity) TableName() string {
	return "images"
}

type CardEntity struct {
	ID                 *int64   `json:"id" gorm:"column:id;primaryKey"`
	Name               *string  `json:"name" gorm:"column:name"`
	FirstEdition       bool     `json:"first_edition" gorm:"column:first_edition"`
	SerieCode          *string  `json:"serie_code" gorm:"column:serie_code"`
	SubTypeID          *int64   `json:"subtype_id" gorm:"column:subtype_id"`
	SubtypeDescription *string  `json:"subtype_description"`
	TypeID             *int64   `json:"type_id"`
	TypeDescription    *string  `json:"type_description"`
	ATK                int64    `json:"atk" gorm:"column:atk"`
	DEF                int64    `json:"def" gorm:"column:def"`
	Stars              int64    `json:"stars" gorm:"column:stars"`
	Description        string   `json:"description" gorm:"column:description"`
	Price              *float64 `json:"price" gorm:"column:price"`
	ImageID            *int64   `json:"image_id" gorm:"column:image_id"`
	ImageDescription   *string  `json:"image_description"`
	DateCreated        *string  `json:"date_created" gorm:"column:date_created"`
}

func (CardEntity) TableName() string {
	return "cards"
}

type NewCardEntity struct {
	ID                 *int64   `json:"id" gorm:"column:id;primaryKey"`
	Name               *string  `json:"name" gorm:"column:name"`
	FirstEdition       bool     `json:"first_edition" gorm:"column:first_edition"`
	SerieCode          *string  `json:"serie_code" gorm:"column:serie_code"`
	SubTypeID          *int64   `json:"subtype_id" gorm:"column:subtype_id"`
	ATK                int64    `json:"atk" gorm:"column:atk"`
	DEF                int64    `json:"def" gorm:"column:def"`
	Stars              int64    `json:"stars" gorm:"column:stars"`
	Description        string   `json:"description" gorm:"column:description"`
	Price              *float64 `json:"price" gorm:"column:price"`
	ImageID            *int64   `json:"image_id" gorm:"column:image_id"`
	DateCreated        *string  `json:"date_created" gorm:"column:date_created"`
}

func (NewCardEntity) TableName() string {
	return "cards"
}

func (ce *CardEntity) CardsEntityToDTO() *CardDTO {
	return &CardDTO{
		ID:                 ce.ID,
		Name:               ce.Name,
		FirstEdition:       ce.FirstEdition,
		SerieCode:          ce.SerieCode,
		SubTypeID:          ce.SubTypeID,
		SubTypeDescription: ce.SubtypeDescription,
		TypeID:             ce.TypeID,
		TypeDescription:    ce.TypeDescription,
		ATK:                ce.ATK,
		DEF:                ce.DEF,
		Stars:              ce.Stars,
		Description:        ce.Description,
		Price:              ce.Price,
		ImageID:            ce.ImageID,
		ImageDescription:   ce.ImageDescription,
		DateCreated:        ce.DateCreated,
	}
}
