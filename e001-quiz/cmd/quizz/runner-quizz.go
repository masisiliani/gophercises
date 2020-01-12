package quizz

import "fmt"

func RunQuizToUser(quizList [][]string) int {
	var (
		response            string
		countCorrectAnswers int
	)
	for _, quiz := range quizList {
		fmt.Printf("%s = ", quiz[0])
		fmt.Scanln(&response)

		isCorrect := verifyResult(response, quiz[1])

		calculateCorrectAnswer(&countCorrectAnswers, isCorrect)
	}

	return countCorrectAnswers
}
