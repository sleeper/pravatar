package main

import (
	"flag"
	"fmt"
	"github.com/WeGoTogether/pravatar/diskstore"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

const (
	defaultPort = "3333"
	defaultHost = ""
	defaultDir  = "images"
)

var router = mux.NewRouter()
var store *diskstore.DiskStore
var host = defaultHost
var port = defaultPort
var dir = defaultDir

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
	flag.StringVar(&port, "port", defaultPort, "port on which to listen")
	flag.StringVar(&host, "host", defaultHost, "host on which to listen")
	flag.StringVar(&dir, "dir", defaultDir, "root dir for images")

	router.HandleFunc("/avatar/{hash}", GetAvatarHandler).Methods("GET")
	router.HandleFunc("/avatar/{hash}", PostAvatarHandler).Methods("POST")
}

func main() {
	flag.Parse()
	var hostAndPort = host + ":" + port

	log.Printf("Listening on %s", hostAndPort)
	log.Printf("Images root dir is %s", dir)

	store = diskstore.NewStore(dir)

	http.Handle("/", router)
	http.ListenAndServe(hostAndPort, nil)
}
