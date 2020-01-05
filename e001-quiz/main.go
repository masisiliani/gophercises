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

	runQuizToUser(quizList)

}

func openFile() (*os.File, error) {
	return os.OpenFile(csvFilename, os.O_RDWR|os.O_CREATE, 0755)
}

func readQuizList(file *os.File) ([][]string, error) {
	reader := csv.NewReader(file)
	return reader.ReadAll()
}

func runQuizToUser(quizList [][]string) {
	var (
		response            string
		countCorrectAnswers int
	)
	for _, quiz := range quizList {
		fmt.Printf("%s = ", quiz[0])
		fmt.Scanln(&response)

		calculateCorrectAnswer(
			&countCorrectAnswers,
			verifyResult(response, quiz[1]))
	}

}

func verifyResult(response string, expectedResponse string) bool {
	if response == expectedResponse {
		return true
	}
	return false
}

func calculateCorrectAnswer(totalCorrect *int, isCorrect bool) {
	if isCorrect {
		*totalCorrect = *totalCorrect + 1
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
