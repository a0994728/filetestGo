package entity

import (
	"database/sql"

	"github.com/shopspring/decimal"
)

type World struct {
	CountryCode       sql.NullString      `db:"country_code"`
	CountryName       sql.NullString      `db:"country_name"`
	Continent         sql.NullString      `db:"continent"`
	Region            sql.NullString      `db:"region"`
	SurfaceArea       decimal.NullDecimal `db:"surface_area"`
	IndepYear         decimal.NullDecimal `db:"indep_year"`
	CountryPopulation decimal.NullDecimal `db:"country_population"`
	LifeExpectancy    decimal.NullDecimal `db:"life_expectancy"`
	Gnp               decimal.NullDecimal `db:"GNP"`
	Gnpold            decimal.NullDecimal `db:"GNPOld"`
	LocalName         sql.NullString      `db:"local_name"`
	GovernmentForm    sql.NullString      `db:"government_form"`
	HeadOfState       sql.NullString      `db:"head_of_state"`
	Capital           decimal.NullDecimal `db:"capital"`
	Code2             sql.NullString      `DB:"code2"`
}

type WorldList []World
