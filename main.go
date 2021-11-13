package main

import (
	"command-line/app"
	"fmt"
	"log"
	"os"
)

func main() {
	// START OMIT
	fmt.Println("starting point")
	application := app.Generator()

	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	// END OMIT
}
