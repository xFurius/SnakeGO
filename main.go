package main

import (
	"fmt"
	"log"
	"os"
	"snake/main/game"
)

func Init() {
	logFile, err := os.Create("./log.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	log.SetOutput(logFile)
}

func main() {
	game.Ihasfi()
}
