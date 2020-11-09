package main

import (

	"context"
	"os"
	"encoding/json"
	"log"
	"net/http"
	"github.com\sidmod\project\helper"
	"github.com\sidmod\project\models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

var collection = helper.ConnectDB()

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var articles []models.Article

	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var article models.Article
		err := cur.Decode(&article) 
		if err != nil {
			log.Fatal(err)
		}

		
		articles = append(articles, article)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(articles) 
}




func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var article models.Article
	
	_ = json.NewDecoder(r.Body).Decode(&article)

	result, err := collection.InsertOne(context.TODO(), article)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/articles", getBooks).Methods("GET")
	r.HandleFunc("/articles", createBook).Methods("POST")
	
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), r))


}