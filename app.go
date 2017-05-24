package main

import (
	"flag"
	"log"
	"os"

	"github.com/google/gops/agent"
	"github.com/julienschmidt/httprouter"
	"github.com/tokopedia/botlog/src/config"
	"github.com/tokopedia/botlog/src/handler"
	"github.com/tokopedia/botlog/src/utils"
	grace "gopkg.in/tokopedia/grace.v1"
	logging "gopkg.in/tokopedia/logging.v1"
)

func init() {
	flag.Parse()
	logging.LogInit()
}

func main() {
	// init gops
	if err := agent.Listen(nil); err != nil {
		log.Fatal(err)
	}

	config.Init()

	router := httprouter.New()

	//for tracking add product bot/user issue
	router.POST("/track/v1/add_product", handler.AddProductTracker)
	router.OPTIONS("/track/v1/add_product", handler.OptionsHandler)

	// run http server
	l := log.New(os.Stdout, "[botlog] ", 0)
	grace.Serve(":8910", utils.RequestLogger{Handle: router, Logger: l})
}
