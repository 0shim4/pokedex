package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Pokedex struct {
	Pokemon []Pokemon `json:"pokemons"`
}

type Pokemon struct {
	No    string `json:"no"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

func main() {
	raw, err := ioutil.ReadFile("docs/index.html")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var fc Pokedex

	json.Unmarshal(raw, &fc)

	os.Mkdir("docs/data", 0777)

	for id, ft := range fc.Pokemon {
		dir := fmt.Sprintf("./docs/data/%v", id+1)
		os.Mkdir(dir, 0777)

		file, _ := json.MarshalIndent(ft, "", " ")
		filePath := dir + "/index.html"
		err = ioutil.WriteFile(filePath, file, 0644)
		if err == nil {
			fmt.Println(filePath)
		}
	}
}
