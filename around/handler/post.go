package handler


import (
   "encoding/json"
   "fmt"
   "net/http"
   "around/model"
)


// deal with upload http request
func uploadHandler(w http.ResponseWriter, r *http.Request) {
   // 1. Parse from body of request to get a json object. (request => struct)
   fmt.Println("Received one upload request") // print to terminal
   decoder := json.NewDecoder(r.Body)
   var p model.Post
   if err := decoder.Decode(&p); err != nil {
       panic(err)
   }


   // 2. call service to handle request
   // 3. construct response
   fmt.Fprintf(w, "Post received: %s\n", p.Message) // print to frontend
}
