package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func prettyPrint(matr [][]rune) {
	var out *bufio.Writer
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	rowCount := len(matr)
	columtCount := len(matr[0])

	for i := range rowCount {
		lastSymb := 0
		for j := columtCount - 1; j >= 0; j-- {
			if matr[i][j] != ' ' {
				lastSymb = j + 1
				break
			}
		}
		for j := range columtCount {
			if j < lastSymb {
				fmt.Fprint(out, string(matr[i][j]))
			}
		}
		fmt.Fprint(out, "\n")
	}
}

func PrintHexagon(matr [][]rune, lineCount, slashCount int) {
	for i := 0; i < len(matr); i++ {
		for j := 0; j < len(matr[0]); j++ {
			if (i == 0 || i == len(matr)-1) && (j >= slashCount && j < lineCount+slashCount) {
				matr[i][j] = '_'
			} else if i >= 1 && i <= slashCount && j == slashCount-i {
				matr[i][j] = '/'
			} else if i >= 1 && i <= slashCount && j == slashCount+lineCount-1+i {
				matr[i][j] = '\\'
			} else if i > slashCount && i <= 2*slashCount && j == (i-slashCount)-1 {
				matr[i][j] = '\\'
			} else if i > slashCount && i <= 2*slashCount && j == len(matr[0])-(i-slashCount) {
				matr[i][j] = '/'
			}
		}
	}
}

func GetHexagos(in *bufio.Reader) [][][]rune {
	var count int
	_, err := fmt.Fscan(in, &count)
	if err != nil {
		log.Fatal(err)
	}

	viewMatricies := make([][][]rune, 0, count)
	for range count {
		var lineCount, slashCount int
		fmt.Fscan(in, &lineCount, &slashCount)

		matr := make([][]rune, (slashCount*2)+1)
		for i := range matr {
			matr[i] = make([]rune, lineCount+(2*slashCount))
		}

		for i := 0; i < len(matr); i++ {
			for j := 0; j < len(matr[0]); j++ {
				matr[i][j] = ' '
			}
		}

		PrintHexagon(matr, lineCount, slashCount)
		viewMatricies = append(viewMatricies, matr)
	}

	return viewMatricies
}

func main() {
	var in *bufio.Reader
	file, err := os.Open("./tests_data/" + "номер набора данных")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	in = bufio.NewReader(file)
	res := GetHexagos(in)
	for _, val := range res {
		prettyPrint(val)
	}
}
