package commands

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"serienManager/config"
)

func Extras(db *sql.DB, logger *log.Logger, conf config.Config){
	var input string
	fmt.Println("\nExtras:\n" +
		"1. Export Database to CSV\n" +
		"2. Add Folder\n" +
		"3. Watch next Episode\n" +
		"4. Rescan Folders")
	fmt.Print("-> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	switch input {
	case "1":
		ExportToCSV(db, logger)
	case "2":
		AddFolder(db, logger)
	case "3":
		WatchNextEpisode(db, logger, conf)
	case "4":
		fmt.Println("NOT IMPLEMENTED YET")

	}
}
