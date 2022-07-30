package main

import (
	"github/gibberish/router"
	"github/gibberish/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	go router.OpenPort2()
	go router.OpenPort3()
	go router.OpenPort4()

	r := mux.NewRouter()
	r.HandleFunc("/one", HandlerOne)

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println(err)
		return
	}
}

func HandlerOne(w http.ResponseWriter, r *http.Request) {
	utils.WriteResponse(w, nil, "Hello one", http.StatusOK)
}
