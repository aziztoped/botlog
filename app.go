package main

import (
	"flag"
	"log"

	"github.com/google/gops/agent"
	"github.com/julienschmidt/httprouter"
	"github.com/tokopedia/botlog/handler"
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

	router := httprouter.New()

	//for tracking add product bot/user issue
	router.POST("/bottrack/v1/add_product", handler.AddProductTracker)
	router.OPTIONS("/bottrack/v1/add_product", handler.OptionsHandler)

	// run http server
	grace.Serve(":8910", router)
}
