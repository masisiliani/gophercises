package quizz

import "fmt"

func verifyResult(response string, expectedResponse string) bool {
	if response == expectedResponse {
		return true
	}
	return false
}

func calculateResultQuiz(countAnswered, countCorrectAnswers int32) int32 {
	countWrongAnswers := countAnswered - countCorrectAnswers
	return countWrongAnswers
}

func calculateCorrectAnswer(totalCorrect *int32, isCorrect bool) {
	if isCorrect {
		*totalCorrect = *totalCorrect + 1
	}
}

func printResultQuiz(countResponse, countCorrectAnswers, countWrongAnswers, totalQuestions int32) {
	fmt.Printf("respondidas: %v | corretas: %v | erradas: %v | total: %v", countResponse, countCorrectAnswers, countWrongAnswers, totalQuestions)
}
