package handler

import (
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tokopedia/botlog/src/config"
)

func AddProductTracker(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	l := config.Logger["product"]

	origin := r.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	byteData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		l.Println("fail to log track add product ->", err)
	}
	content := string(byteData)

	l.Println(content)

	w.WriteHeader(http.StatusOK)
	return
}

func OptionsHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	return
}
