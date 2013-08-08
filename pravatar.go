package main

import (
	"fmt"
	"github.com/WeGoTogether/pravatar/diskstore"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type Pravatar struct {
	Host   string
	Port   string
	Dir    string
	store  diskstore.Storer
	router *mux.Router
}

func NewPravatar(host string, port string, store *diskstore.DiskStore) *Pravatar {
	var p = &Pravatar{Host: host, Port: port, store: store}
	var router = mux.NewRouter()

	p.router = router

	p.initializeRouter()

	return p
}

func (p *Pravatar) initializeRouter() {
	p.router.HandleFunc("/avatar/{hash}", p.getAvatarHandler()).Methods("GET")
	// router.HandleFunc("/avatar/{hash}", GetAvatarHandler).Methods("GET")
	p.router.HandleFunc("/avatar/{hash}", p.postAvatarHandler()).Methods("POST")
}

func (pravatar *Pravatar) getAvatarHandler() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		fmt.Fprintf(writer, "Get Avatar with hash %s received\n", vars["hash"])
		var file, err = pravatar.store.Get(vars["hash"])
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(writer, file)
	}
}

func (pravatar *Pravatar) postAvatarHandler() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		fmt.Fprintf(writer, "Post Avatar with hash %s received\n", vars["hash"])
	}
}

func (p *Pravatar) Listen() {
	var hostAndPort = p.Host + ":" + p.Port

	log.Printf("Listening on %s", hostAndPort)
	log.Printf("Images root dir is %s", p.store.Dir)

	http.Handle("/", p.router)
	http.ListenAndServe(hostAndPort, nil)
}
