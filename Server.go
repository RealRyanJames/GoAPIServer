package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
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

type URIClient struct {
	uri string
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

func LowerURI() string {
	return "mongodb://localhost:27017"
}

func main() {

	clientsURI := URIClient{
		uri: strings.ToLower(LowerURI()),
	}

	client, _ := mongo.Connect(options.Client().ApplyURI(clientsURI.uri))

	c, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer func() {
		if err := client.Disconnect(c); err != nil {
			panic(err)
		}
	}()

	defer cancel()
	_ = client.Ping(c, readpref.Primary())

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
	data.isRunning = true
	data.getData(DATA{
		dataName:  userAppRoute.route,
		len:       0,
		isRunning: true,
	}, "App")

	dataURIClient := MongoURL{
		dbName:       "SimpleAPI",
		isLogginedIn: false,
		collName:     "API",
		data:         []string{""},
	}

	collection := client.Database(dataURIClient.dbName).Collection(dataURIClient.collName)
	res, _ := collection.InsertOne(c, bson.M{"name": fmt.Sprint(string("I am Ryan And This is My Server"))})

	id := res.InsertedID
	fmt.Printf("%s", id)

	if data.isRunning {
		router := httprouter.New()
		router.GET(string(userAppRoute.route), GetIndex)

		log.Fatal(http.ListenAndServe(string(clients_info), router))
	}
}
