package main

// cmd/my-cli/cli.go
import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	err := cli.NewApp().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
