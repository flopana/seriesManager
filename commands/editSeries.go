package commands

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

func EditSeries(db *sql.DB, logger *log.Logger)  {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\nWhich ID do you want to edit?")
	fmt.Print("-> ")
	scanner.Scan()
	id := scanner.Text()
	query := "SELECT title, season, episode, ended FROM serien where id = "+id
	logger.Println("querying db with: " + query)
	rows, err := db.Query(query)
	if err != nil {
		logger.Panic(err)
	}
	defer rows.Close()
	if rows.Next(){
		var title, season, episode, ended string
		err = rows.Scan(&title, &season, &episode, &ended)
		logger.Println("Found ID: "+id+" Title: "+title)
		fmt.Println("\nWhich column do you want to edit?\n" +
			"1) Title: "+title+
			"\n2) Season: "+season+
			"\n3) Episode: "+episode+
			"\n4) Ended: "+ended)
		fmt.Print("-> ")
		scanner.Scan()
		input := scanner.Text()
		switch input {
		case "1":
			fmt.Println("Enter a new title")
			fmt.Print("-> ")
			scanner.Scan()
			input := scanner.Text()
			update(input, "title", id, db, logger)
		case "2":
			fmt.Println("Enter a new season")
			fmt.Print("-> ")
			scanner.Scan()
			input := scanner.Text()
			update(input, "season", id, db, logger)
		case "3":
			fmt.Println("Enter a new episode")
			fmt.Print("-> ")
			scanner.Scan()
			input := scanner.Text()
			update(input, "episode", id, db, logger)
		case "4":
			fmt.Println("Enter a new ended state 0 or 1")
			fmt.Print("-> ")
			scanner.Scan()
			input := scanner.Text()
			if input != "0" && input != "1"{
				fmt.Println("INVALID END STATE")
				return
			}
			update(input, "ended", id, db, logger)
		}
	}else {
		ListSeries(db, logger)
		fmt.Println("NOTHING FOUND PLEASE CHOOSE A ID FROM ABOVE")
		EditSeries(db, logger)
	}
}

func update(str string, column string, id string, db *sql.DB, logger *log.Logger){
	query := "UPDATE serien" +
		" SET "+column+" = '"+str+"'," +
		" edited = '"+time.Now().UTC().Add(time.Hour * 1).Format("2006-01-02T15:04:05Z07:00")+"'" +
		" WHERE id = "+id
	logger.Printf("querying db with: %s", query)
	_, err := db.Exec(query)
	if err != nil {
		logger.Panic(err)
	}
}
