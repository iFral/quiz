package main

import (
	"encoding/csv"
	"flag"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	//Declare flags
	var qFile string
	flag.StringVar(&qFile, "f", "problems.csv", "The quiz file to be used.")

	//Read csv
	log.Infoln("@@@  Reading file  @@@\n" +
		qFile)
	file, err := os.Open(qFile)
	if err != nil {
		log.Errorln("!!!  Error opening quiz file  !!!")
		log.Debugln(err)
	}
	reader := csv.NewReader(file)
	questions, _ := reader.ReadAll()

	log.Info(questions)
}
