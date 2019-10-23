package controller

import(
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/gorilla/mux"
	"github.com/RSR2019/GO-POC/02-NoteBook_MVC/model"
)
// store the note collection
var noteBook = make(map[string]model.Note)

// to generate key for collection
var id int = 0;

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc UserController) PostNoteHandler(res http.ResponseWriter,req *http.Request){
	var note model.Note
	// decode the json from request body and place it into the struct
	err := json.NewDecoder(req.Body).Decode(&note)
	if err != nil{
		panic(err)
	}
	
	note.CreatedOn = time.Now()
	id++
	k := strconv.Itoa(id)
	noteBook[k] = note

	// now marshal struct into json
	jsonData,err := json.Marshal(note)
    if err != nil{
		panic(err)
	}
	res.Header().Set("Content-Type","application/json")
	res.WriteHeader(http.StatusCreated)
	res.Write(jsonData)
}

func (uc UserController) GetNoteHandler(res http.ResponseWriter,req *http.Request){
	var notes []model.Note
	for _,v := range noteBook{
      notes = append(notes,v)
	}
	res.Header().Set("Content-Type","application/json")
	jsonData,err := json.Marshal(notes)
    if err != nil{
		panic(err)
	}
	res.WriteHeader(http.StatusOK)
	res.Write(jsonData)
}

func (uc UserController) PutNoteHandler(res http.ResponseWriter,req *http.Request){
	var err error
	vars := mux.Vars(req)
	k := vars["id"]
	var noteToUpdate model.Note
	err = json.NewDecoder(req.Body).Decode(&noteToUpdate)
	if err != nil{
		panic(err)
	}
	if note,ok := noteBook[k];ok{
		  noteToUpdate.CreatedOn = note.CreatedOn
		  delete(noteBook,k)
		  noteBook[k] = noteToUpdate
	}else{
		log.Printf("couldn't find key %s to delete", k)
	}
	res.WriteHeader(http.StatusNoContent)
	
}
func (uc UserController) DeleteNoteHandler(res http.ResponseWriter,req *http.Request){
 vars := mux.Vars(req)
 k := vars["id"]
 if _,ok := noteBook[k];ok{
	 delete(noteBook,k)
 }else{
	 log.Printf("Could Not find key of Note %s to delete",k)
 }
res.WriteHeader(http.StatusNoContent)
}