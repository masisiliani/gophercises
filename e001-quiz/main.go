package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

var csvFilename string

const defaultFileName = "problems.csv"

func init() {
	flag.StringVar(&csvFilename, "filename", defaultFileName, "nome do arquivo csv com os desafios, no formato 'desafio, resposta'")
	flag.Parse()
}

func main() {

	openFile()
}

func openFile() {
	file, err := os.OpenFile(csvFilename, os.O_RDWR|os.O_CREATE, 0755)
	checkError(err)

	reader := csv.NewReader(file)
	quizzes, err := reader.ReadAll()
	checkError(err)

	var response string
	for _, quiz := range quizzes {
		fmt.Printf("%s = ", quiz[0])
		fmt.Scanln(&response)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
