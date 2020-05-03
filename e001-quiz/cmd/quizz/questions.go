package quizz

import "strings"

type Question struct {
	problem string
	answer  string
}

func convertToQuestionMap(quizzSlice [][]string) []Question {

	questions := make([]Question, len(quizzSlice))

	for index, line := range quizzSlice {
		questions[index] = Question{
			problem: line[0],
			answer:  strings.TrimSpace(line[1]),
		}
	}

	return questions
}
