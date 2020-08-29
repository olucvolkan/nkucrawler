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

	htmlsDocument := makeRequest("http://erdincuzun.cv.nku.edu.tr/")

	instutationInfos := getinsitutationInfos(htmlsDocument)
	insertTeacherInfoDb(instutationInfos)
	fmt.Println(instutationInfos)
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
