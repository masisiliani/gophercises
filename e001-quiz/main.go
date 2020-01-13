package main

import (
	"flag"

	errorPkg "github.com/masisiliani/gophercises/e001-quiz/cmd/error"
	filePkg "github.com/masisiliani/gophercises/e001-quiz/cmd/file"
	"github.com/masisiliani/gophercises/e001-quiz/cmd/quizz"
)

var csvFilename string

const defaultFileName = "problems.csv"

func init() {
	flag.StringVar(&csvFilename, "filename", defaultFileName, "nome do arquivo csv com os desafios, no formato 'desafio, resposta'")
	flag.Parse()
}

func main() {

	objFile, err := filePkg.OpenFile(csvFilename)
	defer objFile.Close()
	errorPkg.CheckError(err)

	quizList, err := filePkg.ReadQuizList(objFile)
	errorPkg.CheckError(err)

	quizz.RunQuizToUser(quizList)

}
