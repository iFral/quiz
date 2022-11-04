package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type question struct {
	q string
	a string
}

func main() {
	//Declare flags
	qFile := flag.String("f", "problems.csv", "The quiz file to be used.")
	qTime := flag.Int("t", 30, "The quiz time limit. (Default is 30s)")
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

	//Start the quiz and the timer
	ready := false
	for {
		var inCh byte
		fmt.Printf("The quiz will start when you are ready. You will have %d seconds to complete it.\nAre you ready? (y/N) ", *qTime)
		fmt.Scanf("%c\n", &inCh)
		ready = (inCh == 'y')
		if ready {
			break
		} else {
			fmt.Println("Start again when you're ready.")
			os.Exit(0)
		}
	}

	//Ask questions and get answers
	log.Infoln("@@@  Asking", len(questions), "question(s)  @@@")
	var index int = 1
	var results = make(map[int]bool)

	//Looping until complete or time is up
	timer := time.After(time.Duration(*qTime) * time.Second)
quizloop:
	for _, i := range questions {
		fmt.Print(index, ". ", i.q, " = ")
		ansCh := make(chan bool)
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			ansCh <- ans == i.a
		}()

		select {
		case <-timer:
			fmt.Println("\nTime's up!")
			break quizloop
		case answered := <-ansCh:
			results[index] = answered
		}
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
