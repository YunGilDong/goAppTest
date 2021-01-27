package app

import "database/sql"

type mariadbHandler struct {
	db *sql.DB
}

func (m *mariadbHandler) Ping() {

}

func newMariadbHandler() DBHandler {

	return &mariadbHandler{}
}
