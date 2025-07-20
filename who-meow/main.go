package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Data struct {
	Name   string
	Points int
}

func WhoMeowing(in *bufio.Reader) []string {
	var err error
	var phrase string
	var count, stringCount int

	fmt.Fscanln(in, &count)
	result := make([]string, 0, count)
	for range count {
		points := make(map[string]int)
		fmt.Fscanln(in, &stringCount)
		{
			dataSlice := make([]Data, 0, stringCount)
			var action string
			for range stringCount {
				phrase, err = in.ReadString('\n')
				if err != nil {
					fmt.Println(err.Error())
					continue
				}

				words := strings.Fields(phrase)
				action = words[len(words)-1]
				points[words[0][:len(words[0])-1]] += 0
				if words[2] == "am" {
					if words[3] == "not" {
						points[words[0][:len(words[0])-1]] -= 1
					} else {
						points[words[0][:len(words[0])-1]] += 2
					}
				} else {
					if words[3] == "not" {
						points[words[1]] -= 1
					} else {
						points[words[1]] += 1
					}
				}
			}
			for key, val := range points {
				dataSlice = append(dataSlice, Data{key, val})
			}
			sort.Slice(dataSlice, func(i, j int) bool {
				return dataSlice[i].Points > dataSlice[j].Points
			})

			maxPoints := dataSlice[0].Points
			lastIdx := len(dataSlice)
			for idx, val := range dataSlice {
				if val.Points != maxPoints {
					lastIdx = idx
					break
				}
			}

			sort.Slice(dataSlice[:lastIdx], func(i, j int) bool {
				return dataSlice[i].Name < dataSlice[j].Name
			})

			for _, data := range dataSlice[:lastIdx] {
				result = append(result, fmt.Sprintf("%s is %s.\n", data.Name, action[:len(action)-1]))
			}
		}
	}

	return result
}

func main() {
	var in *bufio.Reader
	file, err := os.Open("./tests_data/" + "Номер тестового набора")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	in = bufio.NewReader(file)
	res := WhoMeowing(in)
	var out *bufio.Writer
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for _, data := range res {
		fmt.Fprint(out, data)
	}
}
