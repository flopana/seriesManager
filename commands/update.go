package commands

import (
	"log"
	"net/http"

	goUpdate "github.com/inconshreveable/go-update"
)

func DoUpdate(logger *log.Logger){
	url := "s"
	resp, err := http.Get(url)
	if err != nil {
		logger.Panic(err)
	}
	defer resp.Body.Close()
	err = goUpdate.Apply(resp.Body, goUpdate.Options{})
	if err != nil {
		logger.Panic(err)
	}
}
