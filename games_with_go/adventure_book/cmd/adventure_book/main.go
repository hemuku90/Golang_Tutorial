package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	filename := flag.String("file", "gopher.json", "The json file with story")
	flag.Parse()
	fmt.Println(*filename)
	file, err := os.Open(*filename)
	if err != nil {
		fmt.Println("Error while parsing the file", file)
		os.Exit(1)
	}
	d := json.NewDecoder(file)
	var story adventure_book.story
	if err := d.Decode(&story); err != nil {
		fmt.Println("Some error while decoding the file")
	}
	fmt.Printf("%+v\n", story)

}
