package quizz

import (
	"fmt"
	"strings"

	filePkg "github.com/masisiliani/gophercises/e001-quiz/cmd/file"
)

func RunnerQuizz(csvFilename string) {
	objFile, err := filePkg.OpenFile(csvFilename)
	defer objFile.Close()
	if err != nil {
		panic(err)
	}

	quizzSlice, err := filePkg.ReadQuizSlice(objFile)
	if err != nil {
		panic(err)
	}

	questions := convertToQuestionMap(quizzSlice)

	runQuizz(questions)
}

func runQuizz(questions []Question) {
	var (
		response                            string
		countCorrectAnswered, countAnswered int32
	)

	for index, question := range questions {
		fmt.Printf("Problem #%d: %s = \n", index+1, question.problem)
		fmt.Scanf("%s\n", &response)

		if strings.Compare(response, "*") == 0 {
			break
		}

		isCorrect := verifyResult(response, question.answer)

		calculateCorrectAnswer(&countCorrectAnswered, isCorrect)
		countAnswered++
	}

	countWrongAnswers := calculateResultQuiz(countAnswered, countCorrectAnswered)
	printResultQuiz(countAnswered, countCorrectAnswered, countWrongAnswers, int32(len(questions)))

}
