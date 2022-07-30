package router

import (
	"github/gibberish/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func OpenPort2() {
	r := mux.NewRouter()
	r.HandleFunc("/two", HandlerTwo)

	err := http.ListenAndServe(":8081", r)
	if err != nil {
		log.Println(err)
		return
	}
}

func OpenPort3() {
	r := mux.NewRouter()
	r.HandleFunc("/three", HandlerThree)

	err := http.ListenAndServe(":8082", r)
	if err != nil {
		log.Println(err)
		return
	}
}

func OpenPort4() {
	r := mux.NewRouter()
	r.HandleFunc("/four", HandlerFour)

	err := http.ListenAndServe(":8083", r)
	if err != nil {
		log.Println(err)
		return
	}
}

func HandlerTwo(w http.ResponseWriter, r *http.Request) {
	utils.WriteResponse(w, nil, "Hello two", http.StatusOK)
}

func HandlerThree(w http.ResponseWriter, r *http.Request) {
	utils.WriteResponse(w, nil, "Hello three", http.StatusOK)
}

func HandlerFour(w http.ResponseWriter, r *http.Request) {
	utils.WriteResponse(w, nil, "Hello four", http.StatusOK)
}
