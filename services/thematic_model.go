package services

type (
	Thematic struct {
		ID             int64          `json:"id" db:"id" gorm:"primary_key"`
		Nombre         string         `json:"nombre" db:"nombre"`
		Typeofthematic Typeofthematic `gorm:"ForeignKey:Idtype;AssociationForeignKey:ID"`
		Idtype         int64          `json:"idType" db:"idtype"`
	}
	Thematics struct {
		List []Thematic
	}

	Typeofthematic struct {
		ID     int64  `json:"id" db:"id" gorm:"primary_key"`
		Nombre string `json:"nombre" db:"nombre"`
	}
	Typeofthematics struct {
		List []Typeofthematic
	}
)
