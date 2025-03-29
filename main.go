package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var movieCollection *mongo.Collection

type Movie struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Isbn     string             `json:"isbn" bson:"isbn"`
	Title    string             `json:"title" bson:"title"`
	Director *Director          `json:"director" bson:"director"`
}

type Director struct {
	Firstname string `json:"firstname" bson:"firstname"`
	Lastname  string `json:"lastname" bson:"lastname"`
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cursor, err := movieCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, "Failed to fetch movies", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var movies []Movie
	if err = cursor.All(context.TODO(), &movies); err != nil {
		http.Error(w, "Error parsing movies", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	movie.ID = primitive.NewObjectID()
	_, err := movieCollection.InsertOne(context.TODO(), movie)
	if err != nil {
		http.Error(w, "Failed to create movie", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(movie)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}
	res, err := movieCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil || res.DeletedCount == 0 {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Movie deleted"})
}

func main() {
	ctx := context.TODO()
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	movieCollection = client.Database("moviedb").Collection("movies")

	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Server running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
