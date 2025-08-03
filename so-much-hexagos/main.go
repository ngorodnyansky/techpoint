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
		for j := range columtCount {
			fmt.Fprint(out, string(matr[i][j]))
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

func UpdateCoords(matr [][]rune, lineCount, slashCount, iStart, jStart, jEnd int) (int, int, int) {
	if jEnd+lineCount*2+(2*slashCount) > len(matr[0])-2 {
		if jStart%(lineCount*2+(2*slashCount)) == 1 {
			jStart = lineCount + slashCount + 1
			jEnd = jStart + 2*slashCount + lineCount
			iStart = iStart + slashCount
			if jEnd > len(matr[0])-2 {
				jStart = 1
				jEnd = jStart + lineCount + (2 * slashCount)
				iStart = iStart + slashCount
				return iStart, jStart, jEnd
			}
		} else {
			jStart = 1
			jEnd = jStart + lineCount + (2 * slashCount)
			iStart = iStart + slashCount
		}
	} else {
		jStart += lineCount*2 + (2 * slashCount)
		jEnd += lineCount*2 + (2 * slashCount)
	}

	return iStart, jStart, jEnd
}

func PrintHexagons(matr [][]rune, lineCount, slashCount, hexagonCount int) {
	count := 0
	iStart := 1
	jStart := 1
	jEnd := jStart + 2*slashCount + lineCount
	subMatr := make([][]rune, slashCount*2+1)
	for i := range subMatr {
		subMatr[i] = matr[iStart+i][jStart:jEnd]
	}
	PrintHexagon(subMatr, lineCount, slashCount)
	count++
	for count < hexagonCount {
		iStart, jStart, jEnd = UpdateCoords(matr, lineCount, slashCount, iStart, jStart, jEnd)
		for i := range subMatr {
			subMatr[i] = matr[iStart+i][jStart:jEnd]
		}
		PrintHexagon(subMatr, lineCount, slashCount)
		count++
	}

}

func GetHexagons(in *bufio.Reader) [][]rune {
	var maplinesCount, mapRowsCount, lineCount, slashCount, hexagonCount int
	fmt.Fscan(in, &mapRowsCount, &maplinesCount, &lineCount, &slashCount, &hexagonCount)

	matr := make([][]rune, maplinesCount+2)
	for i := range matr {
		matr[i] = make([]rune, mapRowsCount+2)
	}

	for i := 0; i < len(matr); i++ {
		for j := 0; j < len(matr[0]); j++ {
			if i == 0 || i == len(matr)-1 {
				matr[i][j] = '-'
				continue
			} else if j == 0 || j == len(matr[0])-1 {
				matr[i][j] = '|'
				continue
			}

			matr[i][j] = ' '
		}
		matr[0][0] = '+'
		matr[len(matr)-1][0] = '+'
		matr[0][len(matr[0])-1] = '+'
		matr[len(matr)-1][len(matr[0])-1] = '+'
	}
	if hexagonCount == 0 {
		return matr
	}
	PrintHexagons(matr, lineCount, slashCount, hexagonCount)

	return matr
}

func main() {
	var in *bufio.Reader
	file, err := os.Open("./tests_data/" + "Номер тестового набора")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	in = bufio.NewReader(file)
	res := GetHexagons(in)

	prettyPrint(res)

}
