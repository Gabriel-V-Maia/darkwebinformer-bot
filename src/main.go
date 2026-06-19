package main

import (
	"fmt"
	"os"

	"darkwebinformer-bot/src/modules/envreader"
)

func main() {
	envreader.LoadEnv(".env")

	fmt.Println(os.Getenv("TOKEN"))
}
