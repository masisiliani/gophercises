package main

import "flag"

var fileName string

const defaultFileName = "problems.csv"

func init() {
	flag.StringVar(&fileName, "filename", defaultFileName, "nome do arquivo com os desafios")
	flag.Parse()
}

func main() {
}
