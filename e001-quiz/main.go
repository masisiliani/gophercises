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
	defer file.Close()
	checkError(err)

	quizList, err := readQuizList(file)
	checkError(err)

	countCorrectAnswers := runQuizToUser(quizList)

	countCorrectAnswers, countWrongAnswers, totalQuestions := calculateResultQuiz(countCorrectAnswers, len(quizList))
	printResultQuiz(countCorrectAnswers, countWrongAnswers, totalQuestions)

}

func openFile() (*os.File, error) {
	return os.OpenFile(csvFilename, os.O_RDWR|os.O_CREATE, 0755)
}

func readQuizList(file *os.File) ([][]string, error) {
	reader := csv.NewReader(file)
	return reader.ReadAll()
}

func runQuizToUser(quizList [][]string) int {
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

	return countCorrectAnswers
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

func calculateResultQuiz(countCorrectAnswers int, totalQuestions int) (int, int, int) {
	countWrongAnswers := totalQuestions - countCorrectAnswers
	return countCorrectAnswers, countWrongAnswers, totalQuestions
}

func printResultQuiz(countCorrectAnswers int, countWrongAnswers int, totalQuestions int) {
	fmt.Printf("corretas: %v | erradas: %v | total: %v", countCorrectAnswers, countWrongAnswers, totalQuestions)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
