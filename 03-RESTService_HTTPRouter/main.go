package main

import (

	"net/http"
	
	"github.com/julienschmidt/httprouter"
	
	"github.com/RSR2019/GO-POC/03-RESTService/controllers"
)

func main(){
    r := httprouter.New()
    uc := controllers.NewUserController()
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)	
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}