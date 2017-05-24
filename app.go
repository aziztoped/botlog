package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/gops/agent"
	"github.com/julienschmidt/httprouter"

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

	// Endpoint for tracking add product bot/user issue
	router.POST("/log/v1/add_product", addProductTracker)
	router.OPTIONS("/log/v1/add_product", optionsHandler)

	// run http server
	grace.Serve(":8910", router)
}

func addProductTracker(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	origin := r.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	byteData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("fail to log track add product ->", err)
	}
	content := string(byteData)

	// write log
	log.Println(content)

	w.WriteHeader(http.StatusOK)
	return
}

func optionsHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	return
}
