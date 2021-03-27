package commands

import (
	"bufio"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func AddFolder(db *sql.DB, logger *log.Logger)  {
	var path string
	var seriesid string
	fmt.Println("Enter the Folder Path you'd like to add")
	fmt.Print("-> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	path = scanner.Text()
	fmt.Println("Enter the associated series id")
	fmt.Print("-> ")
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	seriesid = scanner.Text()
	if !seriesExists(seriesid, db, logger){
		fmt.Printf("Series does not exist!\n")
		return
	}
	query := "INSERT INTO folder (path, seriesid) VALUES ('"+path+"', "+seriesid+");"
	logger.Printf("queyring db with: %s",query)
	_, err := db.Exec(query)
	if err != nil {
		logger.Fatalln(err)
	}

	scanFolderForFiles(seriesid ,db, logger)
}

func scanFolderForFiles(id string, db *sql.DB, logger *log.Logger) {
	var files[] string
	query := "SELECT id, path from folder where seriesid ="+id
	logger.Printf("querying db with: %s", query)
	rows, err := db.Query(query)
	if err != nil {
		logger.Fatalln(err)
	}
	for rows.Next(){
		var path string
		var folderid int
		err = rows.Scan(&folderid, &path)
		items, _ := ioutil.ReadDir(path)
		for _, item := range items {
			if !item.IsDir(){
				logger.Printf("found file %s", item.Name())
				files = append(files, path+"\\"+item.Name())
			}
		}
		insertFiles(db, logger, files, folderid)
	}
}

func insertFiles(db *sql.DB, logger *log.Logger, files []string, folderid int) {
	query := "INSERT INTO files (path, folderid) VALUES "
	for _, file := range files{
		query += "('"+file+"', "+strconv.Itoa(folderid)+"),"
	}
	query = strings.TrimSuffix(query, ",")
	logger.Printf("queyring db with: %s", query)
	_, err := db.Exec(query)
	if err != nil {
		logger.Fatalln(err)
	}
}

func seriesExists(id string, db *sql.DB, logger *log.Logger) bool {
	query := "SELECT * FROM serien where id = "+id
	logger.Printf("querying db with: %s", query)
	rows, err := db.Query(query)
	if err != nil {
		logger.Panic(err)
	}
	defer rows.Close()
	if rows.Next(){
		var id, title, season, episode, ended, edited string
		err = rows.Scan(&id, &title, &season, &episode, &ended, &edited)
		var input string
		fmt.Printf("Do you really want to add this Folder to Title: %s  (y/n)\n", title)
		fmt.Print("-> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input = scanner.Text()
		input = strings.ToLower(input)
		if input == "y" {
			return true
		}else{
			return false
		}
	}else {
		fmt.Println("NOTHING FOUND")
		logger.Printf("Couldn't find a series with id = %s", id)
		return false
	}
}
