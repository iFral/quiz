package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	//Declare flags
	var qFile string
	flag.StringVar(&qFile, "f", "problems.csv", "The quiz file to be used.")

	//Read csv
	log.Infoln("@@@  Reading file", qFile, "  @@@")
	file, err := os.Open(qFile)
	if err != nil {
		log.Errorln("!!!  Error opening quiz file  !!!")
		log.Debugln(err)
	}
	reader := csv.NewReader(file)
	questions, _ := reader.ReadAll()

	//Ask questions and get answers
	log.Infoln("@@@  Asking", len(questions), "question(s)  @@@")
	var index int = 1
	var results = make(map[int]bool)
	for _, q := range questions {
		fmt.Print(index, ". ", q[0], " = ")
		var ans string
		fmt.Scanln(&ans)
		results[index] = q[1] == ans
		index++
	}
	log.Infoln("@@@  Quiz completed. Tabulating score.  @@@")
	var correct int = 0
	for _, r := range results {
		if r {
			correct++
		}
	}
	fmt.Println("You got", correct, "of", len(questions), "correct.")

}
