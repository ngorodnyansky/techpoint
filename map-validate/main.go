package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}

type MapElem struct {
	X, Y int
}

func (m MapElem) IsNeighbour(elem MapElem) bool {
	diffX := Abs(m.X - elem.X)
	diffY := Abs(m.Y - elem.Y)
	if (diffX == 1 && diffY == 1) || (diffX == 0 && diffY == 2) {
		return true
	}

	return false
}

type ElementsSet map[MapElem]bool

func (s ElementsSet) Add(e MapElem) {
	s[e] = true
}

func (s ElementsSet) AddSet(newSet ElementsSet) {
	for key, val := range newSet {
		s[key] = val
	}
}

func (s ElementsSet) Contains(e MapElem) bool {
	_, ok := s[e]
	return ok
}

func deleteElem(arr []MapElem, delIdx int) []MapElem {
	arr[delIdx] = arr[len(arr)-1]
	return arr[:len(arr)-1]
}

func FindNeighbours(startPoint MapElem, points []MapElem) ElementsSet {
	result := ElementsSet{}
	result.Add(startPoint)

	neighbours := make([]MapElem, 0)
	delIdxs := make([]int, 0)
	for i := 0; i < len(points); i++ {
		if startPoint.IsNeighbour(points[i]) {
			neighbours = append(neighbours, points[i])
			delIdxs = append(delIdxs, i)
		}
	}

	for i := len(delIdxs) - 1; i>=0; i--{
		points = deleteElem(points, delIdxs[i])
	}

	
	for _, point := range neighbours {
		result.Add(point)
	}

	for _, point := range neighbours {
		result.AddSet(FindNeighbours(point, points))
	}

	return result
}

func MapIsValid(mapData map[rune][]MapElem) string {
	for _, val := range mapData {
		if len(val) == 1 {
			continue
		}
		startPoint := val[0]
		mapPoints := FindNeighbours(startPoint, val[1:])
		if len(mapPoints) == len(val) {
			continue
		}
		return "NO"
	}

	return "YES"
}

func MapsIsValid(in *bufio.Reader) []string {
	var count int
	_, err := fmt.Fscan(in, &count)
	if err != nil {
		log.Fatal(err)
	}

	result := make([]string, 0, count)
	for range count {
		var rowCount, columtCount int
		fmt.Fscan(in, &rowCount, &columtCount)

		mapData := make(map[rune][]MapElem)
		var str string
		for i := range rowCount {
			fmt.Fscan(in, &str)
			for j, val := range str {
				if val == '.' {
					continue
				}
				mapData[val] = append(mapData[val], MapElem{X: i, Y: j})
			}
		}
		result = append(result, MapIsValid(mapData))
	}

	return result
}

func main() {
	var in *bufio.Reader
	file, err := os.Open("./tests_data/" + "4")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	in = bufio.NewReader(file)
	res := MapsIsValid(in)
	var out *bufio.Writer
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, val := range res {
		fmt.Fprint(out, val)
		fmt.Fprint(out, "\n")
	}
}
