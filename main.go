package main

import (
	"fmt"
	"os"

	"io/ioutil"
	"log"
	"net/http"

	"github.com/tkanos/gonfig"
)

type Pairs struct {
	A_key string
	B_key string
}

type Apis struct {
	Endpoint string
	Mapping  []Pairs
}

type Configuration struct {
	Name  string
	Mocks []Apis
}

func main() {
	config := GetConfig()
	response := CallEndpoint(config.Mocks[0].Endpoint)
	fmt.Print(Convert(response, config.Mocks[0].Mapping))
}

func GetConfig(params ...string) Configuration {
	configuration := Configuration{}
	fileName := "./config.json"
	gonfig.GetConf(fileName, &configuration)
	return configuration
}

func CallEndpoint(endpoint string) string {
	response, err := http.Get(endpoint)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(responseData)
}

func Convert(response string, pairs []Pairs) string {
	for _, pair := range pairs {
		fmt.Println(pair.A_key)
		fmt.Println(pair.B_key)
	}
	return ""
}
