package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)
func initDb(logger *log.Logger) *sql.DB {
	db, err := sql.Open("sqlite3", "serien.sqlite")
	if err != nil {
		logger.Fatalln(err)
	}
	query := "CREATE TABLE IF NOT EXISTS serien (" +
		"id INTEGER not null primary key," +
		"title TEXT not null," +
		"season INTEGER not null," +
		"episode INTEGER not null," +
		"ended INTEGER DEFAULT 0," +
		"edited TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL);" + //either 1 or 0
		"" +
		"CREATE TABLE IF NOT EXISTS folder(" +
		"id INTEGER not null primary key," +
		"path TEXT not null," +
		"seriesid INTEGER not null," +
		"FOREIGN KEY(seriesid) REFERENCES serien(id));" +
		"" +
		"CREATE TABLE IF NOT EXISTS files(" +
		"id INTEGER not null primary key," +
		"path TEXT not null," +
		"folderid INTEGER not null," +
		"FOREIGN KEY(folderid) REFERENCES folder(id));"
	_, err = db.Exec(query)
	if err != nil {
		logger.Panic(err)
	}
	_, err = db.Exec("PRAGMA journal_mode=WAL;")
	if err != nil {
		logger.Panic(err)
	}
	return db
}
