package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Coord [2]int

func (c Coord) Sum() int {
	return c[0] + c[1]
}

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

func findTopLeftWay(coord Coord, matr [][]rune, symb rune) {
	for {
		for coord[0] != 0 {
			newY := coord[0] - 1
			if matr[newY][coord[1]] == '.' {
				matr[newY][coord[1]] = symb
				coord[0] -= 1
			} else {
				break
			}
		}
		for coord[1] != 0 {
			newX := coord[1] - 1
			if matr[coord[0]][newX] == '.' {
				matr[coord[0]][newX] = symb
				coord[1] -= 1
			} else {
				break
			}
		}

		if coord.Sum() == 0 {
			return
		}
	}
}

func findBottomRight(coord Coord, matr [][]rune, symb rune) {
	Y := len(matr) - 1
	X := len(matr[0]) - 1

	for {
		for coord[0] != Y {
			newY := coord[0] + 1
			if matr[newY][coord[1]] == '.' {
				matr[newY][coord[1]] = symb
				coord[0] += 1
			} else {
				break
			}
		}
		for coord[1] != X {
			newX := coord[1] + 1
			if matr[coord[0]][newX] == '.' {
				matr[coord[0]][newX] = symb
				coord[1] += 1
			} else {
				break
			}
		}

		if coord.Sum() == X+Y {
			return
		}
	}
}

func FindRoad(in *bufio.Reader) [][][]rune {
	var count int
	_, err := fmt.Fscan(in, &count)
	if err != nil {
		log.Fatal(err)
	}

	result := make([][][]rune, 0, count)
	for range count {
		var rowCount, columtCount int
		fmt.Fscan(in, &rowCount, &columtCount)

		matr := make([][]rune, rowCount)
		for i := range matr {
			matr[i] = make([]rune, columtCount)
		}

		runeA := rune('A')
		runeB := rune('B')
		coordA := Coord{}
		coordB := Coord{}

		var str string
		for i := range rowCount {
			fmt.Fscan(in, &str)
			for j, val := range str {
				matr[i][j] = val
				switch val {
				case runeA:
					coordA[0] = i
					coordA[1] = j
				case runeB:
					coordB[0] = i
					coordB[1] = j
				}
			}
		}

		if coordA.Sum() < coordB.Sum() {
			findTopLeftWay(coordA, matr, 'a')
			findBottomRight(coordB, matr, 'b')
		} else {
			findTopLeftWay(coordB, matr, 'b')
			findBottomRight(coordA, matr, 'a')
		}

		result = append(result, matr)
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
	res := FindRoad(in)
	for _, val := range res {
		prettyPrint(val)
	}
}
