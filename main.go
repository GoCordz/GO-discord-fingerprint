package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func fingerprint() {
	Client := &http.Client{}
	req, err := http.NewRequest("GET", "https://discord.com/api/v9/experiments", nil)
	if err != nil {
		panic(err)
	}
	for x, o := range map[string]string{
		"accept":          "*/*",
		"accept-language": "en-US,en;q=0.9",
	} {
		req.Header.Set(x, o)
	}
	resp, err := Client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	fingerprint := data["fingerprint"].(string)
	log.Println(" \033[36m|\033[39m ", fingerprint)
	f, err := os.OpenFile("fingerprints.txt", os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, ers := f.WriteString(fingerprint + "\n")
	if ers != nil {
		log.Fatal(ers)
	}

}

func main() {
	fmt.Println("\033[36m\nGO \033[39m-\033[36m DCFG")
	fmt.Println("_________________________________\033[39m")
	for true {
		fingerprint()
	}

}
