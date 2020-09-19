package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
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
	//createTeachersTable()
	//createEducationTable()
	//createAcademicJobsTable()
	//createAdministrativeDuties()
	//createGivenLessons()
	//createLectures()
	//createResearch()
	//createProject()
	createTeachInfo()
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

func createEducationTable() {
	db := dbConn()
	query := `CREATE TABLE nkucrawler.educaton  (
		teacher_id int NOT NULL,
		type char(255) NULL,
		university char(255) NULL,
		faculty char(255) NULL,
		department char(255) NULL,
		year char(255) NULL,
		thesis char(255) NULL,
		CONSTRAINT teacher_id FOREIGN KEY (teacher_id) REFERENCES nkucrawler.teachers (id)
	  );`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

func createAcademicJobsTable() {
	db := dbConn()

	query := `CREATE TABLE academic_jobs (
	  id int NOT NULL AUTO_INCREMENT,
	  teacherId int DEFAULT NULL,
	  title char(255) COLLATE utf8_turkish_ci DEFAULT NULL,
	  school varchar(255) COLLATE utf8_turkish_ci DEFAULT NULL,
	  year varchar(255) COLLATE utf8_turkish_ci DEFAULT NULL,
	  PRIMARY KEY (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_turkish_ci;`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

func createAdministrativeDuties() {
	db := dbConn()

	query := `CREATE TABLE administrative_duties (
	  id int NOT NULL AUTO_INCREMENT,
	  teacherId int DEFAULT NULL,
	  title char(255) COLLATE utf8_turkish_ci DEFAULT NULL,
	  school varchar(255) COLLATE utf8_turkish_ci DEFAULT NULL,
	  year varchar(255) COLLATE utf8_turkish_ci DEFAULT NULL,
	  PRIMARY KEY (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_turkish_ci;`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}
func createGivenLessons() {
	db := dbConn()

	query := `CREATE TABLE given_lessons (
	  id int NOT NULL AUTO_INCREMENT,
	  teacherId int DEFAULT NULL,
	  period char(255) COLLATE utf8_turkish_ci DEFAULT NULL,
	  lessonName varchar(255) COLLATE utf8_turkish_ci DEFAULT NULL,
	  time int DEFAULT NULL,
	  PRIMARY KEY (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_turkish_ci;`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

func createLectures() {
	db := dbConn()

	query := `CREATE TABLE lectures (
	  id int NOT NULL AUTO_INCREMENT,
	  teacherId int DEFAULT NULL,
	  lessonName varchar(255) COLLATE utf8_turkish_ci DEFAULT NULL,
	  PRIMARY KEY (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_turkish_ci;`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

func createResearch() {
	db := dbConn()

	query := `CREATE TABLE research (
	  id int NOT NULL AUTO_INCREMENT,
	  teacherId int DEFAULT NULL,
	  researchSubject varchar(255) COLLATE utf8_turkish_ci DEFAULT NULL,
	  PRIMARY KEY (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_turkish_ci;`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

func createProject() {
	db := dbConn()

	query := `CREATE TABLE project (
	  id int NOT NULL AUTO_INCREMENT,
	  teacherId int DEFAULT NULL,
	  project varchar(255) COLLATE utf8_turkish_ci DEFAULT NULL,
	  PRIMARY KEY (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_turkish_ci;`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

func createTeachInfo() {
	db := dbConn()

	query := `CREATE TABLE teach_info (
	  id int NOT NULL AUTO_INCREMENT,
	  teacherId int DEFAULT NULL,
	  teachType varchar(255) COLLATE utf8_turkish_ci DEFAULT NULL,
	  facultyType varchar(255) COLLATE utf8_turkish_ci DEFAULT NULL,
	  facultyName varchar(255) COLLATE utf8_turkish_ci DEFAULT NULL,
	  PRIMARY KEY (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_turkish_ci;`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}
