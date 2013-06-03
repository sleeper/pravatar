package main

import (
	"fmt"
	"github.com/WeGoTogether/pravatar/diskstore"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

var router = mux.NewRouter()
var store = diskstore.NewStore("tests")

func GetAvatarHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprintf(writer, "Get Avatar with hash %s received\n", vars["hash"])
	var file, err = store.Get(vars["hash"])
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(writer, file)
}

func PostAvatarHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprintf(writer, "Post Avatar with hash %s received\n", vars["hash"])
}

func init() {
	router.HandleFunc("/avatar/{hash}", GetAvatarHandler).Methods("GET")
	router.HandleFunc("/avatar/{hash}", PostAvatarHandler).Methods("POST")
}

func main() {
	http.Handle("/", router)
	http.ListenAndServe(":3333", nil)
}
