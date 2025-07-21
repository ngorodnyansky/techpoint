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
	fmt.Fprint(out, "\n")
}

func PaintOver(from [][]rune, to [][]rune) {
	for i := range to {
		for j := range to[i] {
			if from[i][j] != '.' {
				to[i][j] = from[i][j]
			}
		}
	}
}

func FlatMatrix(views [][][]rune) [][]rune {
	if len(views) == 1 {
		return views[0]
	}

	lastView := len(views) - 1
	for i := len(views) - 2; i >= 0; i-- {
		PaintOver(views[i], views[lastView])
	}

	return views[lastView]
}

func GetView(in *bufio.Reader) [][][]rune {
	var count int
	_, err := fmt.Fscan(in, &count)
	if err != nil {
		log.Fatal(err)
	}

	result := make([][][]rune, 0, count)
	for range count {
		var viewCount, rowCount, columtCount int
		fmt.Fscan(in, &viewCount, &rowCount, &columtCount)

		viewMatricies := make([][][]rune, 0, viewCount)
		for range viewCount {
			matr := make([][]rune, rowCount)
			for i := range matr {
				matr[i] = make([]rune, columtCount)
			}

			var str string
			for i := range rowCount {
				fmt.Fscan(in, &str)
				for j, val := range str {
					matr[i][j] = val
				}
			}

			viewMatricies = append(viewMatricies, matr)
		}
		result = append(result, FlatMatrix(viewMatricies))
	}

	return result
}

func main() {
	var in *bufio.Reader
	file, err := os.Open("./tests_data/" + "номер набора")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	in = bufio.NewReader(file)
	res := GetView(in)
	for _, val := range res {
		prettyPrint(val)
	}
}
