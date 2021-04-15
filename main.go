package main

import (
	"bufio"
	"fmt"
	"os"
	"seriesManager/commands"
	"seriesManager/config"
)

func main() {
	logger := GetLogger()
	var conf = config.ParseConfig(logger)
	db := initDb(logger)
	printLogo()
	printOptions()


	for {
		var input string
		fmt.Print("-> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input = scanner.Text()
		switch input {
		case "1":
			commands.ListSeries(db, logger)
		case "2":
			commands.EditSeries(db, logger)
		case "3":
			commands.AddSeries(db, logger)
		case "4":
			commands.DeleteSeries(db, logger)
		case "5":
			commands.Extras(db, logger, conf)
		case "6":
			logger.Printf("Exiting")
			err := db.Close()
			if err != nil {
				logger.Fatalln(err)
			}
			os.Exit(0)

		}
		printOptions()
	}
}

func printLogo() {
	fmt.Println("______ _       _       _____           _             ___  ___                                  \n|  ___| |     ( )     /  ___|         (_)            |  \\/  |                                  \n| |_  | | ___ |/ ___  \\ `--.  ___ _ __ _  ___ _ __   | .  . | __ _ _ __   __ _  __ _  ___ _ __ \n|  _| | |/ _ \\  / __|  `--. \\/ _ \\ '__| |/ _ \\ '_ \\  | |\\/| |/ _` | '_ \\ / _` |/ _` |/ _ \\ '__|\n| |   | | (_) | \\__ \\ /\\__/ /  __/ |  | |  __/ | | | | |  | | (_| | | | | (_| | (_| |  __/ |   \n\\_|   |_|\\___/  |___/ \\____/ \\___|_|  |_|\\___|_| |_| \\_|  |_/\\__,_|_| |_|\\__,_|\\__, |\\___|_|   \n                                                                                __/ |          \n                                                                               |___/          ")
}

func printOptions() {
	fmt.Println("\n1. List all Series sorted by last edited\n" +
		"2. Edit a series\n" +
		"3. Add a new one\n" +
		"4. Delete\n" +
		"5. Extras\n"+
		"6. Quit")
}
