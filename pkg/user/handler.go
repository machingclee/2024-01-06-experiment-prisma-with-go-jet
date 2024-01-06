package user

import (
	"authentication/.gen/udemy/public/model"
	"authentication/.gen/udemy/public/table"
	"database/sql"

	. "github.com/go-jet/jet/v2/postgres"
)

type HandlerRepo struct {
	db *sql.DB
}

var Repo *HandlerRepo

func NewHandler(db *sql.DB) {
	Repo = &HandlerRepo{
		db: db,
	}
}

func (m *HandlerRepo) GetUsers() []model.User {
	db := m.db
	statement := SELECT(table.User.AllColumns).FROM(table.User)
	var users []model.User
	statement.Query(db, &users)
	return users
}
