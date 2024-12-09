package main

import "net/http"

func handler_readiness(w http.ResponseWriter, r *http.Request){
	respondWithJSON(w, 200, struct{}{})
}

func handler_err(w http.ResponseWriter, r *http.Request){
	respondWithError(w, 400, "something went wrong")
}