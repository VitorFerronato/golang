package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "daumata", 32}
	fmt.Println(c)

	cachorroEmJson, erro := json.Marshal(c)

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(bytes.NewBuffer(cachorroEmJson))
}
