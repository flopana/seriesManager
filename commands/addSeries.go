package commands

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

func AddSeries(db *sql.DB, logger *log.Logger)  {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\nType in the Title of the Series")
	fmt.Print("-> ")
	scanner.Scan()
	title := scanner.Text()

	fmt.Println("Season:")
	fmt.Print("-> ")
	scanner.Scan()
	season := scanner.Text()

	fmt.Println("Episode:")
	fmt.Print("-> ")
	scanner.Scan()
	episode := scanner.Text()

	query := "INSERT INTO serien (title, season, episode, edited) VALUES('"+ title +"', "+ season +", "+episode+", '"+time.Now().UTC().Add(time.Hour * 1).Format("2006-01-02T15:04:05Z07:00")+"')"
	logger.Printf("Inserting new series with: %s", query)
	_, err := db.Exec(query)
	if err != nil {
		logger.Fatalln(err)
	}

	fmt.Println("Added Successfully")
}
