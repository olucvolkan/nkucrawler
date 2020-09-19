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
	htmlDocuments := makeRequest("http://cmf.nku.edu.tr/PersonelListesi/0/s/9790/801")
	urlList := urlList(htmlDocuments)
	for i, s := range urlList {
		htmlDocuments := makeRequest(s)
		instutationInfos := getinsitutationInfos(htmlDocuments)
		insertTeacherInfoDb(instutationInfos)

		primaryKey := i + 1
		academicJobs := academicJobs(htmlDocuments, primaryKey)
		fmt.Println(academicJobs)

		insertAdministrativeDuties := administrativeDuties(htmlDocuments, primaryKey)
		fmt.Println(insertAdministrativeDuties)

		givenLessons := givenLessons(htmlDocuments, primaryKey)
		fmt.Println(givenLessons)

		lectures := lectures(htmlDocuments, primaryKey)
		fmt.Println(lectures)

		research := research(htmlDocuments, primaryKey)
		fmt.Println(research)

		projects := projects(htmlDocuments, primaryKey)
		fmt.Println(projects)

		teachInfo := teachInfo(htmlDocuments, primaryKey)
		fmt.Println(teachInfo)
	}
}

func urlList(document *goquery.Document) []string {
	urlList := []string{}

	document.Find("#icerik > div:nth-child(1) > div > b").Each(func(i int, s *goquery.Selection) {
		siteName := s.Find("div > div.col-md-9.col-xs-8 > b > b > h6:nth-child(3) > a").Text()
		urlList = append(urlList, siteName)
	})

	return urlList
}

func dbConn() (db *sql.DB) {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "nkucrawler_prod"
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
	//Burda ki sorunu cozemedim tekrar bakacagim

	document.Find("#dersler > div.breadcrumb").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Html())
		year := s.Find("td:nth-child(1)").Text()
		lessonName := s.Find("td:nth-child(2)").Text()
		period := s.Find("td:nth-child(3)").Text()

		insertGivenLessons(year, lessonName, period, teacherID)
	})
	return true
}

func insertGivenLessons(year string, lessonName string, period string, teacherID int) {
	db := dbConn()

	insertQuery, err := db.Prepare("INSERT INTO given_lessons(period, lessonName,time,teacherId) VALUES(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	insertQuery.Exec(year, lessonName, period, teacherID)

	fmt.Println("ADDED given lessons")

	defer db.Close()
}
func lectures(document *goquery.Document, teacherID int) bool {
	document.Find("#yayinlar > div:nth-child(2) > div.panel-body.table-responsive > table > tbody > tr").Each(func(i int, s *goquery.Selection) {
		lessonName := s.Find("td:nth-child(2)").RemoveAttr("<br>").Text()
		fmt.Println(lessonName)
		insertLectures(lessonName, teacherID)
	})
	return true
}

func insertLectures(lessonName string, teacherID int) {
	db := dbConn()

	insertQuery, err := db.Prepare("INSERT INTO lectures(lessonName,teacherId) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}

	insertQuery.Exec(lessonName, teacherID)

	fmt.Println("ADDED given lectures")

	defer db.Close()
}

func research(document *goquery.Document, teacherID int) bool {

	document.Find("#arastirma > div.panel.panel-danger > div.panel-body.table-responsive > h6").Each(func(i int, s *goquery.Selection) {
		researchSubject := s.Text()
		insertResearch(researchSubject, teacherID)
	})
	return true
}

func insertResearch(researchSubject string, teacherID int) {
	db := dbConn()

	insertQuery, err := db.Prepare("INSERT INTO research(researchSubject,teacherID) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}

	insertQuery.Exec(researchSubject, teacherID)

	fmt.Println("ADDED given lectures")

	defer db.Close()
}

func projects(document *goquery.Document, teacherID int) bool {
	document.Find("#proje > div.panel.panel-info > div.panel-body.table-responsive > table > tbody > tr").Each(func(i int, s *goquery.Selection) {
		project := s.Find("td:nth-child(2)").Text()
		fmt.Println(project)
		insertProject(project, teacherID)
	})
	return true
}

func insertProject(project string, teacherID int) {
	db := dbConn()

	insertQuery, err := db.Prepare("INSERT INTO project(project,teacherID) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}

	insertQuery.Exec(project, teacherID)

	fmt.Println("ADDED project")

	defer db.Close()
}
func teachInfo(document *goquery.Document, teacherID int) bool {

	document.Find("#ogrenimbilgileri > div").Each(func(i int, s *goquery.Selection) {
		teachType := s.Find("div.panel-heading").Text()
		s.Find("div.panel-body.table-responsive > table > tbody > tr").Each(func(i int, s *goquery.Selection) {
			facultyType := s.Find("td:nth-child(1)").Text()
			facultyName := s.Find("td:nth-child(3)").Text()
			fmt.Println(teachType)
			fmt.Println(facultyType)
			fmt.Println(facultyName)

			teachInfoInsert(facultyType, facultyName, teachType, teacherID)
		})
	})
	return true
}

func teachInfoInsert(facultyType string, facultyName string, teachType string, teacherID int) {
	db := dbConn()

	insertQuery, err := db.Prepare("INSERT INTO teach_info(facultyType,facultyName,teachType,teacherID) VALUES(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	insertQuery.Exec(facultyType, facultyName, teachType, teacherID)

	fmt.Println("ADDED teach info")

	defer db.Close()
}
