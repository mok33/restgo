package main

import (
	"fmt"
	"os"
    "encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)
type attr struct{
	key string `json: "key"`
	val string `json: "val"`
}

type Attributes struct{
	Props []attr
}
type planet_info struct{
	Name string `json: "name"`
	Distance string `json: "distance from Sun"`
	Mass string `json: "mass"`
	Temperature string `json: "surface temperature (mean)"`
	Volume string `json: volume`
}

var Solar_system map[string]interface{}


func main() {
	file, e := ioutil.ReadFile("./solar-system.json")
	if e != nil {
	    fmt.Printf("File error: %v\n", e)
	    os.Exit(1)
	}
	    
	json.Unmarshal(file, &Solar_system)

	router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/planet/{name}", getPlanet)

    log.Fatal(http.ListenAndServe(":8080", router))
   
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcoume kamline")
}

func getPlanet(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    planet_name := params["name"]

   	planets := Solar_system["sections"]
    var pl planet_info
    var tags [5]string

    for key, val := range planets.(map[string]interface{}){
    	if key == planet_name{
    		i := 0
    		for _,e := range val.([]interface{}){
    			for _, v := range e.(map[string]interface{}){
    				tags[i] = v.(string)
    			}
    			i++
    		}

    		pl = planet_info{
    			planet_name,
    			tags[0],
    			tags[1],
    			tags[2],
    			tags[3],
    		}
    	}
    }

      json.NewEncoder(w).Encode(pl)
}
