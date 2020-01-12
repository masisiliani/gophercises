package file

import (
	"encoding/csv"
	"os"
)

func OpenFile(csvFilename string) (*os.File, error) {
	return os.OpenFile(csvFilename, os.O_RDWR|os.O_CREATE, 0755)
}

func ReadQuizList(file *os.File) ([][]string, error) {
	reader := csv.NewReader(file)
	return reader.ReadAll()
}
