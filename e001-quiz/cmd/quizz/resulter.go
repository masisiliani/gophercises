package quizz

import "fmt"

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
