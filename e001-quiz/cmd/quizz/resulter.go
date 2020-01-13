package quizz

import "fmt"

func verifyResult(response string, expectedResponse string) bool {
	if response == expectedResponse {
		return true
	}
	return false
}

func calculateResultQuiz(countCorrectAnswers int, totalQuestions int) (int, int) {
	countWrongAnswers := totalQuestions - countCorrectAnswers
	return countWrongAnswers, totalQuestions
}

func calculateCorrectAnswer(totalCorrect *int, isCorrect bool) {
	if isCorrect {
		*totalCorrect = *totalCorrect + 1
	}
}

func printResultQuiz(countCorrectAnswers int, countWrongAnswers int, totalQuestions int) {
	fmt.Printf("corretas: %v | erradas: %v | total: %v", countCorrectAnswers, countWrongAnswers, totalQuestions)
}
