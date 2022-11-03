package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

type question struct {
	q string
	a string
}

func main() {
	//Declare flags
	qFile := flag.String("f", "problems.csv", "The quiz file to be used.")
	flag.Parse()

	//Read csv
	log.Infoln("@@@  Reading file", *qFile, "  @@@")
	file, err := os.Open(*qFile)
	if err != nil {
		log.Errorln("!!!  Error opening quiz file  !!!")
		log.Debugln(err)
		os.Exit(1)
	}
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		log.Errorln("!!!  Error reading csv file  !!!")
		log.Debugln(err)
		os.Exit(1)
	}
	questions := make([]question, len(lines))
	for i, line := range lines {
		questions[i] = question{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	//Ask questions and get answers
	log.Infoln("@@@  Asking", len(questions), "question(s)  @@@")
	var index int = 1
	var results = make(map[int]bool)
	for _, i := range questions {
		fmt.Print(index, ". ", i.q, " = ")
		var ans string
		fmt.Scanf("%s\n", &ans)
		results[index] = i.a == ans
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
