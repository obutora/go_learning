package learn

import (
	"encoding/json"
	"log"
	"os"
)

type Child struct {
	StructName string `json:"structName"`
}

type Data struct {
	Name  string `json:"name"`
	Int   int    `json:"int"`
	Child Child  `json:"child"`
}

func ReadJson() {
	f, err := os.Open("testJson.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var data Data
	err = json.NewDecoder(f).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("% +v", data)
}
