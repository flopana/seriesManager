package commands

import (
	"database/sql"
	"encoding/csv"
	"log"
	"os"
)

func ExportToCSV(db *sql.DB, logger *log.Logger){
	query := "SELECT * from serien"
	logger.Printf("querying db with: %s", query)
	rows, err := db.Query(query)
	if err != nil {
		logger.Panic(err)
	}

	var data[][] string = convertRowsToString(rows)

	file, err := os.Create("export.csv")
	if err != nil {
		logger.Fatalln(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	logger.Printf("Writing Data to CSV File")
	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			logger.Fatalln(err)
		}
	}
	logger.Printf("Finished writing to CSV")
}

func convertRowsToString(rows *sql.Rows) [][]string {
	var data[][] string
	for rows.Next(){
		var id, title, season, episode, ended, edited string
		_ = rows.Scan(&id, &title, &season, &episode, &ended, &edited)
		data = append(data,[]string{id, title, season, episode, ended, edited})
	}
	return data
}
