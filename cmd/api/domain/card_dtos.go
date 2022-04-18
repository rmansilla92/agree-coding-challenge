package domain

type CardDTO struct {
	ID                 *int64   `json:"id"`
	Name               *string  `json:"name"`
	FirstEdition       bool     `json:"first_edition"`
	SerieCode          *string  `json:"serie_code"`
	SubTypeID          *int64   `json:"subtype_id"`
	SubTypeDescription *string  `json:"subtype_description"`
	TypeID             *int64   `json:"type_id"`
	TypeDescription    *string  `json:"type_description"`
	ATK                int64    `json:"atk"`
	DEF                int64    `json:"def"`
	Stars              int64    `json:"stars"`
	Description        string   `json:"description"`
	Price              *float64 `json:"price"`
	ImageID            *int64   `json:"image_id"`
	ImageDescription   *string  `json:"image_description"`
	DateCreated        *string  `json:"date_created"`
}

func (cdto *CardDTO) CardsDTOToEntity() *NewCardEntity {
	return &NewCardEntity{
		Name:               cdto.Name,
		FirstEdition:       cdto.FirstEdition,
		SerieCode:          cdto.SerieCode,
		SubTypeID:          cdto.SubTypeID,
		ATK:                cdto.ATK,
		DEF:                cdto.DEF,
		Stars:              cdto.Stars,
		Description:        cdto.Description,
		Price:              cdto.Price,
		ImageID:            cdto.ImageID,
		DateCreated:        cdto.DateCreated,
	}
}