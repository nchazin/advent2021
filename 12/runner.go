package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func cave_in_path(cave string, path []string) bool {
	for _, c := range path {
		if c == cave {
			return true
		}
	}
	return false
}

var paths [][]string

func find_paths(cavemap map[string][]string, cave string, path []string) {
	//fmt.Println("Called with", cave)
	if cave == "start" && len(path) > 0 {
		//	fmt.Println("Skipping start from second pass")
		return
	}
	if cave == "end" {
		path = append(path, "end")
		paths = append(paths, path)
		return
	}
	if cave[0] >= 'a' && cave[0] <= 'z' && cave_in_path(cave, path) {
		//fmt.Printf("Skipping %s for %s\n", cave, path)
		return
	}
	//fmt.Println("Adding cave to path:", cave)
	path = append(path, cave)
	for _, nextcave := range cavemap[cave] {
		find_paths(cavemap, nextcave, path)
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cavemap := map[string][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		caves := strings.Split(line, "-")
		if cavelist, ok := cavemap[caves[0]]; ok {
			cavemap[caves[0]] = append(cavelist, caves[1])
		} else {
			cavemap[caves[0]] = []string{caves[1]}
		}
		if cavelist, ok := cavemap[caves[1]]; ok {
			cavemap[caves[1]] = append(cavelist, caves[0])
		} else {
			cavemap[caves[1]] = []string{caves[0]}
		}
	}
	find_paths(cavemap, "start", []string{})
	//fmt.Println(paths)
	fmt.Println(len(paths))

	/*
		for k, v := range cavemap {
			fmt.Printf("%s -> %s\n", k, v)

		}
	*/

}
