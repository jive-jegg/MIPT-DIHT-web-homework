package main 

import (
	"encoding/json"
	"net/http"
	"strconv"
	"io/ioutil"

	"github.com/gorilla/mux"
)


type TodoItem struct {
	Url  string `json:"url"`
}


var reductions = make(map[string]TodoItem)


func ReduceUrl(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil {
        panic(err)
    }
    var obj TodoItem
    err = json.Unmarshal(data, &obj)
	i := len(reductions)
	i++
	reductions[strconv.Itoa(i)] = obj
	mapVar := map[string]string{"key": strconv.Itoa(i)}
	answ, err := json.Marshal(mapVar)
	if err != nil {
		panic(err)
	}
	w.Write(answ)
}


func GetAnswer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i := vars["key"]
	w.Header().Set("Location", reductions[i].Url)
	w.WriteHeader(http.StatusMovedPermanently)
}


func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", ReduceUrl).Methods("POST")
	router.HandleFunc("/{key}", GetAnswer)

	http.ListenAndServe(":8082", router)
}