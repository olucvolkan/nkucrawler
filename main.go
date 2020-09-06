package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	for i, s := range urlList() {

		htmlsDocument := makeRequest(s)
		//instutationInfos := getinsitutationInfos(htmlsDocument)
		//	insertTeacherInfoDb(instutationInfos)

		primaryKey := i + 1
		//academicJobs := academicJobs(htmlsDocument, primaryKey)
		//fmt.Println(academicJobs)
		insertAdministrativeDuties := administrativeDuties(htmlsDocument, primaryKey)
		fmt.Println(insertAdministrativeDuties)

	}
}

func urlList() []string {

	urlList := []string{
		"http://erdincuzun.cv.nku.edu.tr/",
	}

	return urlList
}

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

func makeRequest(url string) *goquery.Document {
	resp, _ := http.Get(url)
	// Convert HTML into goquery document
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	return doc
}

func removeCharacters(input string, characters string) string {
	filter := func(r rune) rune {
		if strings.IndexRune(characters, r) < 0 {
			return r
		}
		return -1
	}

	return strings.Map(filter, input)

}

// GetLatestBlogTitles gets the latest blog title headings from the url
// given and returns them as a list.
func getinsitutationInfos(document *goquery.Document) []string {

	infos := []string{}
	document.Find("#anasayfa > div.panel.panel-info > div.panel-body.table-responsive > table > tbody > tr >  td:nth-child(3)").Each(func(i int, s *goquery.Selection) {
		infos = append(infos, s.Text())
	})

	teacherInfoArray := []string{}

	teacherName := document.Find("#anasayfa > div.panel.panel-success > div.panel-body > p:nth-child(1)").Text()
	phoneNumber := document.Find("#anasayfa > div.panel.panel-success > div.panel-body > p:nth-child(2)").Text()
	formattedPhoneNumber := removeCharacters(phoneNumber, "T:")
	email := document.Find("#anasayfa > div.panel.panel-success > div.panel-body > p:nth-child(3) > a").Text()
	website := document.Find("#anasayfa > div.panel.panel-success > div.panel-body > p:nth-child(4) > a").Text()
	teacherInfoArray = append(teacherInfoArray, teacherName, formattedPhoneNumber, email, website)

	allData := append(teacherInfoArray, infos...)
	return allData

}

func insertTeacherInfoDb(teacherInfoData []string) {
	db := dbConn()

	teacherName := teacherInfoData[0]
	phoneNumber := teacherInfoData[1]
	mailAddress := teacherInfoData[2]
	siteAddress := teacherInfoData[3]
	faculty := teacherInfoData[4]
	branchScience := teacherInfoData[5]

	insertQuery, err := db.Prepare("INSERT INTO teachers(name, phone_number,mail,site,faculty,branch) VALUES(?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	insertQuery.Exec(teacherName, phoneNumber, mailAddress, siteAddress, faculty, branchScience)

	fmt.Println("ADDED: Name: " + teacherName + " | Mail: " + mailAddress)

	defer db.Close()

}

func academicJobs(document *goquery.Document, teacherID int) []string {
	titles := []string{}
	document.Find("#akademikgorevler > div > div.panel-body.table-responsive > table > tbody > tr ").Each(func(i int, s *goquery.Selection) {
		title := s.Find("td:nth-child(1)").Text()
		school := s.Find("td:nth-child(2)").Text()
		year := s.Find("td:nth-child(3)").Text()

		insertAcademicJobs(title, school, year, teacherID)
	})
	return titles
}
func insertAcademicJobs(title string, school string, year string, teacherID int) {
	db := dbConn()

	insertQuery, err := db.Prepare("INSERT INTO academic_jobs(title, school,year,teacherId) VALUES(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	insertQuery.Exec(title, school, year, teacherID)

	fmt.Println("ADDED academis jobs")

	defer db.Close()
	//#idarigorevler > div > div.panel-body.table-responsive > table > tbody > tr:nth-child(1)
}

func administrativeDuties(document *goquery.Document, teacherID int) bool {

	document.Find("#idarigorevler > div > div.panel-body.table-responsive > table > tbody > tr").Each(func(i int, s *goquery.Selection) {
		title := s.Find("td:nth-child(1)").Text()
		school := s.Find("td:nth-child(2)").Text()
		year := s.Find("td:nth-child(3)").Text()

		insertAdministrativeDuties(title, school, year, teacherID)
	})
	return true
}

func insertAdministrativeDuties(title string, school string, year string, teacherID int) {
	db := dbConn()

	insertQuery, err := db.Prepare("INSERT INTO administrative_duties(title, school,year,teacherId) VALUES(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	insertQuery.Exec(title, school, year, teacherID)

	fmt.Println("ADDED adminstrative Duties")

	defer db.Close()

}

func givenLessons(document *goquery.Document, teacherID int) bool {

	document.Find("#dersler > div:nth-child(2) > div.panel-body.table-responsive > table > tbody > tr").Each(func(i int, s *goquery.Selection) {
		year := s.Find("td:nth-child(1)").Text()
		lessonName := s.Find("td:nth-child(2)").Text()
		clock := s.Find("td:nth-child(3)").Text()

		insertGivenLessons(year, lessonName, clock, teacherID)
	})
	return true
}

func insertGivenLessons(year string, lessonName string, clock string, teacherID int) {
	db := dbConn()

	insertQuery, err := db.Prepare("INSERT INTO given_lessons(title, school,year,teacherId) VALUES(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	insertQuery.Exec(year, lessonName, clock, teacherID)

	fmt.Println("ADDED given lessons")

	defer db.Close()
}
