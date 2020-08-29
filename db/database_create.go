package main

import (
	"database/sql"
	"log"
)

func dbConn() (db *sql.DB) {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "nkucrawler"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	return db

}

func main() {
	createTeachersTable()
}

func createTeachersTable() {
	db := dbConn()
	query := `
	CREATE TABLE teachers (
		id int unsigned NOT NULL AUTO_INCREMENT,
		name char(100) CHARACTER SET utf8 COLLATE utf8_turkish_ci DEFAULT NULL,
		phone_number char(20) CHARACTER SET utf8 COLLATE utf8_turkish_ci DEFAULT NULL,
		mail char(100) COLLATE utf8_turkish_ci DEFAULT NULL,
		site char(100) COLLATE utf8_turkish_ci DEFAULT NULL,
		faculty char(100) COLLATE utf8_turkish_ci DEFAULT NULL,
		branch char(100) COLLATE utf8_turkish_ci DEFAULT NULL,
		PRIMARY KEY (id)
    );`

	// Executes the SQL query in our database. Check err to ensure there was no error.

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}
