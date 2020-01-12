package main

import (
	"flag"
)

var csvFilename string

const defaultFileName = "problems.csv"

func init() {
	flag.StringVar(&csvFilename, "filename", defaultFileName, "nome do arquivo csv com os desafios, no formato 'desafio, resposta'")
	flag.Parse()
}

func main() {

	// objFile, err := file.OpenFile()
	// defer objFile.Close()
	// checkError(err)

	// quizList, err := readQuizList(objFile)
	// checkError(err)

	// countCorrectAnswers := runQuizToUser(quizList)

	// countCorrectAnswers, countWrongAnswers, totalQuestions := calculateResultQuiz(countCorrectAnswers, len(quizList))
	// printResultQuiz(countCorrectAnswers, countWrongAnswers, totalQuestions)

}
