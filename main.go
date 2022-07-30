package main

import (
	"fmt"
	"github/gibberish/utils"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// go router.OpenPort2()
	// go router.OpenPort3()
	// go router.OpenPort4()

	r := mux.NewRouter()
	r.HandleFunc("/one", HandlerOne)

	port := os.Getenv("PORT")

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if err != nil {
		log.Println(err)
		return
	}
}

func HandlerOne(w http.ResponseWriter, r *http.Request) {
	utils.WriteResponse(w, nil, "Hello one", http.StatusOK)
}
