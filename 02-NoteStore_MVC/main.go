package main

import (
	"net/http"

	"github.com/RajatWithGolang/13-GO-REST_Mongo_POC/02-NoteStore_MVC/controller"
	"github.com/gorilla/mux"
)

func main() {

	rtr := mux.NewRouter().StrictSlash(false)
	uc := controller.NewUserController()
	rtr.HandleFunc("/api/notes", uc.GetNoteHandler).Methods("GET")
	rtr.HandleFunc("/api/notes", uc.PostNoteHandler).Methods("POST")
	rtr.HandleFunc("/api/notes/{id}", uc.GetNoteHandler).Methods("PUT")
	rtr.HandleFunc("/api/notes/{id}", uc.GetNoteHandler).Methods("DELETE")

	http.ListenAndServe("8080", rtr)

}
