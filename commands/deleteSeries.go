package commands

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
)

func DeleteSeries(db *sql.DB, logger *log.Logger){
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\nWhich ID do you want to DELETE?")
	fmt.Print("-> ")
	scanner.Scan()
	id := scanner.Text()

	query := "SELECT title, season, episode, ended FROM serien where id = "+id
	logger.Printf("Querying the db to find out if ID: %s exists", id)
	logger.Printf("Query: %s", query)
	rows, err := db.Query(query)
	if err != nil {
		logger.Fatalln(err)
	}
	defer rows.Close()
	if rows.Next(){
		var title, season, episode, ended string
		err = rows.Scan(&title, &season, &episode, &ended)
		fmt.Println("Do you really want to DELETE ID: "+id+" Title: "+ title + " y/n")
		fmt.Print("-> ")
		scanner.Scan()
		input := scanner.Text()
		input = strings.ToLower(input)
		if input == "y"{
			query := "DELETE FROM serien where id = "+id
			_, err = db.Exec(query)
			if err != nil {
				logger.Fatalf("%s",err)
			}
			logger.Println("Deleted series ID: "+id+" Title: "+ title + " y/n")
			logger.Printf("Query: %s", query)
		}else {
			fmt.Println("Nothing deleted")
		}

	}
}
