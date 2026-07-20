package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type ClientInfo struct {
	PORT    string
	isOwned bool
	length  int
}

func GetPort(client ClientInfo) string {
	if client.isOwned {

		client.PORT = ":3000"
		return client.PORT
	}

	client.PORT = ":3000"
	return string(client.PORT)
}

type Route struct {
	route       string
	textMessage string
}

func GetIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	fmt.Fprintf(w, string("I am Ryan And This is My Server"))
}

func main() {

	clients_info := GetPort(ClientInfo{
		isOwned: true,
		PORT:    ":3000",
		length:  0,
	})

	userAppRoute := Route{
		route:       "/",
		textMessage: strings.ToUpper("I am Ryan And This is My Server"),
	}

	router := httprouter.New()
	router.GET(string(userAppRoute.route), GetIndex)

	log.Fatal(http.ListenAndServe(string(clients_info), router))

}
