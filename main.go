package main

import (
	"encoding/json"
	"fmt"
	"os"

	"io/ioutil"
	"log"
	"net/http"

	"github.com/tkanos/gonfig"
)

type Pairs struct {
	A_key string `json:"A_key"`
	B_key string `json:"B_key"`
}

type Apis struct {
	Endpoint string  `json:"Endpoint"`
	Mapping  []Pairs `json:"Mapping"`
}

type Configuration struct {
	Name  string `json:"Name"`
	Mocks []Apis `json:"Mocks"`
}

type ToDo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	// config := GetConfig()
	// fmt.Print(config.Mocks[0].Mapping[0].A_key)
	http.HandleFunc("/mock", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening to port: %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func handler(response http.ResponseWriter, request *http.Request) {
	config := GetConfig()
	data := CallEndpoint(config.Mocks[0].Endpoint)
	Convert(data, config.Mocks[0].Mapping)
	response.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(response, PrettyPrint(data))
}

func GetConfig(params ...string) Configuration {
	configuration := Configuration{}
	fileName := "./config.json"
	gonfig.GetConf(fileName, &configuration)
	return configuration
}

func CallEndpoint(endpoint string) ToDo {
	response, err := http.Get(endpoint)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result ToDo
	if err := json.Unmarshal(responseData, &result); err != nil {
		// Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	return result
}

func Convert(response ToDo, pairs []Pairs) string {
	// for _, pair := range pairs {
	// 	fmt.Println(pair.A_key)
	// 	fmt.Println(pair.B_key)
	// }

	for key, value := range pairs {
		fmt.Println(key, value)
	}

	// newData, err := json.Marshal(response)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println(response)
	return ""
}

func PrettyPrint(i interface{}) string {
	s, _ := json.Marshal(i)
	return string(s)
}
