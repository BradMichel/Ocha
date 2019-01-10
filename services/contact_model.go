package services

import (
	"database/sql"
	"time"
)

type (
	Contact struct {
		ID                     int64          `json:"id" db:"id" gorm:"primary_key"`
		Nombre                 string         `json:"nombre" db:"nombre"`
		Telefono               sql.NullString `json:"telefono" db:"telefono"`
		Fax                    sql.NullString `json:"fax" db:"fax"`
		Correo                 sql.NullString `json:"correo" db:"correo"`
		Celular                sql.NullString `json:"celular" db:"celular"`
		Apellido               sql.NullString `json:"apellido" db:"apellido"`
		Municipio              sql.NullInt64  `json:"municipio" db:"municipio"`
		MunicipioResidencia    sql.NullInt64  `json:"municipioResidencia" db:"municipioResidencia"`
		DepartamentoResidencia sql.NullInt64  `json:"departamentoResidencia" db:"departamentoResidencia"`
		FechaActualizacion     *time.Time     `json:"fechaActualizacion" db:"fechaActualizacion"`
		Tematica               Thematic       `gorm:"ForeignKey:TematicaID.Int64;AssociationForeignKey:ID"`
		Idtematica             sql.NullInt64  `json:"idTematica" db:"idTematica" gorm:"index"`
		Organizacion           Organization   `gorm:"ForeignKey:OrganizacionID.Int64;AssociationForeignKey:ID"`
		Idorganizacion         sql.NullInt64  `json:"idOrganizacion" db:"idOrganization" gorm:"index"`
	}
	Contacts struct {
		List   []Contact `json:"list"`
		Page   int       `json:"page"`
		IDLast int       `json:"idLast"`
		Limit  int       `json:"limit"`
	}
)
