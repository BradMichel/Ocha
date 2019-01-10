package services

import (
	"database/sql"
	"time"
)

type (
	Organization struct {
		ID                 int64          `json:"id" db:"id"`
		Nombre             sql.NullString `json:"nombre" db:"nombre"`
		Sigla              sql.NullString `json:"sigla" db:"sigla"`
		Sede               sql.NullString `json:"sede" db:"sede"`
		Direccion          sql.NullString `json:"direccion" db:"direccion"`
		Telefono           sql.NullString `json:"telefono" db:"telefono"`
		Representante      sql.NullString `json:"representante" db:"representante"`
		FechaActualizacion *time.Time     `json:"fechaActualizacion" db:"fechaActualizacion"`
	}

	Organizations struct {
		List   []Organization `json:"list"`
		Page   int            `json:"page"`
		IDLast int            `json:"idLast"`
		Limit  int            `json:"limit"`
	}
)
