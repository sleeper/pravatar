package main

import (
	"fmt"
	"github.com/WeGoTogether/pravatar/diskstore"
	"github.com/gorilla/mux"
	"io"
  "io/ioutil"
	"log"
	"net/http"
)

type Pravatar struct {
	Host   string
	Port   string
	Dir    string
	store  diskstore.Storer
	Router *mux.Router
}

func NewPravatar(host string, port string, store diskstore.Storer) *Pravatar {
	var p = &Pravatar{Host: host, Port: port, store: store}
	var router = mux.NewRouter()

	p.Router = router

	p.initializeRouter()

	return p
}

func (p *Pravatar) initializeRouter() {
	p.Router.HandleFunc("/avatar/{hash}", p.getAvatarHandler()).Methods("GET")
	// router.HandleFunc("/avatar/{hash}", GetAvatarHandler).Methods("GET")
	p.Router.HandleFunc("/avatar/{hash}", p.postAvatarHandler()).Methods("POST")
}

func (pravatar *Pravatar) getAvatarHandler() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
//		fmt.Fprintf(writer, "Get Avatar with hash %s received\n", vars["hash"])
		var file, err = pravatar.store.Get(vars["hash"])
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(writer, file)
	}
}

func (pravatar *Pravatar) postAvatarHandler() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
    var body []byte
    var err error

		vars := mux.Vars(request)

    if request.Body == nil {
      log.Printf("Received POST request for hash %s with no body", vars["hash"])
      return
    }

    body, err = ioutil.ReadAll(request.Body)

    if err != nil {
			log.Fatal(err)
    }

//    body := request.Body.String()

		fmt.Fprintf(writer, "Post Avatar with hash %s received\n", vars["hash"])
		fmt.Fprintf(writer, "                 body %s received\n", body)
    err = pravatar.store.Put(vars["hash"], body)

    if err != nil {
      log.Printf("Unable to store image: %s", err)
    }
	}
}

func (p *Pravatar) Listen() {
	var hostAndPort = p.Host + ":" + p.Port

	log.Printf("Listening on %s", hostAndPort)
	//log.Printf("Images root dir is %s", p.store.Dir)

	http.Handle("/", p.Router)
	http.ListenAndServe(hostAndPort, nil)
}
