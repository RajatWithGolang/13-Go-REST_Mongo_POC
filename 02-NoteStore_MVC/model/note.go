package model
import "time"

// Note struct type is used for storing json data
type Note struct{
   Title string `json:"title"`
   Description string `json:"description"`
   CreatedOn time.Time `json:"createdon"`
}

