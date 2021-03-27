package commands

import (
	"database/sql"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
)

func ListSeries(db *sql.DB, logger *log.Logger)  {
	logger.Println("Listing all series")
	var data[][] string

	query := "SELECT * from serien order by ended desc, edited"
	rows, err := db.Query(query)
	if err != nil {
		logger.Panic(err)
	}
	defer rows.Close()
	for rows.Next(){
		var id, title, season, episode, ended, edited string
		err = rows.Scan(&id, &title, &season, &episode, &ended, &edited)
		data = append(data,[]string{id, title, season, episode, ended, edited})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Title", "Season", "Episode", "Ended", "Edited"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output

}
