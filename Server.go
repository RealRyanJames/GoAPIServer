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

type DATA struct {
	dataName  string
	isRunning bool
	len       int
}

type MongoURL struct {
	dbName       string
	isLogginedIn bool
	collName     string
	data         []string
}

func (d *DATA) getData(data DATA, name string) string {

	if data.isRunning {
		data.len = 0
		data.len += 1
	}

	data.dataName = string(name)

	return string(data.dataName)
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

	data := DATA{}
	data.getData(DATA{
		dataName:  userAppRoute.route,
		len:       0,
		isRunning: true,
	}, "App")

	if data.isRunning {
		router := httprouter.New()
		router.GET(string(userAppRoute.route), GetIndex)

		log.Fatal(http.ListenAndServe(string(clients_info), router))
	}

	dataURIClient := MongoURL{
		dbName:       "SimpleAPI",
		isLogginedIn: false,
		collName:     "API",
		data:         []string{""},
	}

	if dataURIClient.isLogginedIn == false {
		dataURIClient.isLogginedIn = true

	}
}
