package quizz

import "fmt"

import "strings"

func RunQuizToUser(quizList [][]string) {
	var (
		response                            string
		countCorrectAnswered, countAnswered int32
	)

	for _, quiz := range quizList {
		fmt.Printf("%s = ", quiz[0])
		fmt.Scanln(&response)

		if strings.Compare(response, "*") == 0 {
			break
		}

		isCorrect := verifyResult(response, quiz[1])

		calculateCorrectAnswer(&countCorrectAnswered, isCorrect)
		countAnswered++
	}

	countWrongAnswers := calculateResultQuiz(countAnswered, countCorrectAnswered)
	printResultQuiz(countAnswered, countCorrectAnswered, countWrongAnswers, int32(len(quizList)))

}
