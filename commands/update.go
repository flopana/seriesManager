package commands

import (
	"fmt"
	"log"
	"net/http"

	goUpdate "github.com/inconshreveable/go-update"
)

func DoUpdate(logger *log.Logger){
	url := "https://github.com/flopana/seriesManager/releases/latest/download/seriesManager.exe"
	resp, err := http.Get(url)
	if err != nil {
		logger.Panic(err)
	}
	defer resp.Body.Close()
	err = goUpdate.Apply(resp.Body, goUpdate.Options{})
	if err != nil {
		logger.Panic(err)
	}
	fmt.Println("PLEASE RESTART")
}
