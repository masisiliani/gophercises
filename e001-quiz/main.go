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

	file, err := openFile()
	checkError(err)

	quizList, err := readQuizList(file)
	checkError(err)

	requestQuizToUser(quizList)

}

func openFile() (*os.File, error) {
	return os.OpenFile(csvFilename, os.O_RDWR|os.O_CREATE, 0755)
}

func readQuizList(file *os.File) ([][]string, error) {
	reader := csv.NewReader(file)
	return reader.ReadAll()
}

func requestQuizToUser(quizList [][]string) {
	var response string
	for _, quiz := range quizList {
		fmt.Printf("%s = ", quiz[0])
		fmt.Scanln(&response)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
