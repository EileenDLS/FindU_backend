//implement Router function
package handler


import (
   // "net/http"
   "github.com/gorilla/mux"  
)

func InitRouter() *mux.Router {
   // initiate a router
   router := mux.NewRouter()
   // deal with upload handler
   // router.Handle("/upload", http.HandlerFunc(uploadHandler)).Methods("POST") // use net/http lib
   router.HandleFunc("/upload", uploadHandler).Methods("POST") // use gorilla/mux lib
   return router
}
