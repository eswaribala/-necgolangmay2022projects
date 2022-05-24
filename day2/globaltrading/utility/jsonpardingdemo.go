package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	var user, _ = ioutil.ReadFile("data.json")

	var result []map[string]interface{}

	json.Unmarshal([]byte(user), &result)

	for _, value := range result {

		for innerkey, innervalue := range value {
			fmt.Printf("%s=>%s\n", innerkey, innervalue)
		}
	}

}
