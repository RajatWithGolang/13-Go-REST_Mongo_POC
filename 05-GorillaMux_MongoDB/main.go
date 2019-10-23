package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
	"net/http"
)

type DB struct {
	session    *mgo.Session
	collection *mgo.Collection
}

// Movie holds a movie data
type Movie struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name     string        `json:"name" bson:"name"`
	Year     string        `json:"year" bson:"year"`
	Director string        `json:"directors" bson:"directors"`
	Writer   string        `json:"writers" bson:"writers"`
}

func (db *DB) GetMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	var movie Movie
	err := db.collection.Find(bson.M{"_id": bson.ObjectIdHex(vars["id"])}).One(&movie)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(movie)
		w.Write(response)
	}

}
func (db *DB) PostMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	postBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(postBody, &movie)
	// Create a Hash ID to insert
	movie.ID = bson.NewObjectId()
	err := db.collection.Insert(movie)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(movie)
		w.Write(response)
	}
}

// UpdateMovie modifies the data of given resource
func (db *DB) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var movie Movie
	putBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(putBody, &movie)
	// Create an Hash ID to insert
	err := db.collection.Update(bson.M{"_id": bson.ObjectIdHex(vars["id"])}, bson.M{"$set": &movie})
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "text")
		w.Write([]byte("Updated succesfully!"))
	}
}
func (db *DB) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// Create an Hash ID to insert
	err := db.collection.Remove(bson.M{"_id": bson.ObjectIdHex(vars["id"])})
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "text")
		w.Write([]byte("Deleted succesfully!"))
	}
}

func main() {
	session, err := mgo.Dial("127.0.0.1")
	c := session.DB("appdb").C("movies")
	db := &DB{session: session, collection: c}
	if err != nil {
		panic(err)
	}
	defer session.Close()
	// Create a new router
	router := mux.NewRouter()
	// Attach an elegant path with handler
	router.HandleFunc("/v1/movies/{id:[a-zA-Z0-9]*}", db.GetMovie).Methods("GET")
	router.HandleFunc("/v1/movies", db.PostMovie).Methods("POST")
	router.HandleFunc("/v1/movies/{id:[a-zA-Z0-9]*}", db.UpdateMovie).Methods("PUT")
	router.HandleFunc("/v1/movies/{id:[a-zA-Z0-9]*}", db.DeleteMovie).Methods("DELETE")
	log.Fatalln(http.ListenAndServe(":8080", router))
}
