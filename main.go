package main

import (
	"log"
    // "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/repo/{username}",GitHandle).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080",router))	
}



func GitHandle(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	username := vars["username"]
	statement := fmt.Sprintf("https://api.github.com/users/%s/repos",username)
	fmt.Println(statement)
    response, err := http.Get(statement)
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
		// fmt.Println(string(data))
		var jsonObjs interface{}
		json.Unmarshal([]byte(data),&jsonObjs)
		objSlice, ok := jsonObjs.([]interface{})
		if !ok {
			fmt.Println("cannot convert the JSON objects")
		}
   
		// fmt.Println("Number of JSON objects : ", len(objSlice))
		num_of_repo := fmt.Sprintf("Number of JSON objects : %d", len(objSlice))
		// response, _ := json.Marshal(objSlice)
		// w.Write(len(response))
		response, _ := json.Marshal(num_of_repo)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(response)
   
		// iterate the JSON slice
		// for _, obj := range objSlice {
   
		// 	// convert each obj to map[string]interface{} from interface{}
		// 	// because, interface{} alone does not support indexing
   
		// 	objMap, ok := obj.(map[string]interface{})
   
		// 	if !ok {
		// 		fmt.Println("cannot convert interface{} to type map[string]interface{}")
		// 	}
		// 	fmt.Println(objMap)
   
		// 	// now we can access the data with key index
		// }

	}

	
    // jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
    // jsonValue, _ := json.Marshal(jsonData)
    // response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
    // if err != nil {
    //     fmt.Printf("The HTTP request failed with error %s\n", err)
    // } else {
    //     data, _ := ioutil.ReadAll(response.Body)
    //     fmt.Println(string(data))
    // }
    // fmt.Println("Terminating the application...")

}