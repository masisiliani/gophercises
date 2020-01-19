package main

import (
	"flag"

	"github.com/masisiliani/gophercises/e001-quiz/cmd/quizz"
)

var csvFilename string

const defaultFileName = "problems.csv"

func init() {
	flag.StringVar(&csvFilename, "filename", defaultFileName, "nome do arquivo csv com os desafios, no formato 'desafio, resposta'")
	flag.Parse()
}

func main() {

	quizz.RunnerQuizz(csvFilename)

}
