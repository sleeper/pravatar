package main

import (
	"flag"
	"github.com/WeGoTogether/pravatar/diskstore"
)

const (
	defaultPort = "3333"
	defaultHost = ""
	defaultDir  = "images"
)

var host = defaultHost
var port = defaultPort
var dir = defaultDir

func init() {
	flag.StringVar(&port, "port", defaultPort, "port on which to listen")
	flag.StringVar(&host, "host", defaultHost, "host on which to listen")
	flag.StringVar(&dir, "dir", defaultDir, "root dir for images")

}

func main() {
	flag.Parse()

	var store = diskstore.NewStore(dir)
	var server = NewPravatar(host, port, store)

	server.Listen()
}
