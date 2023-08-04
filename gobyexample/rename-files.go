package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	file, err := os.Open("list.txt")
	if err != nil {
		log.Fatalln("Error opening file:" + err.Error())
	}
	defer file.Close()

	idx := 1
	regex := regexp.MustCompile(`\[(.*?)\]`)

	mapping := make(map[string]int)
	foundMapping := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		match := regex.FindStringSubmatch(line)
		if len(match) == 2 {
			mapping[normalizeName(match[1])] = idx
			foundMapping[normalizeName(match[1])] = false
			idx++
		}
	}

	fileNameRegex := regexp.MustCompile(`(\d{1,2})-(.+)\.go`)
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		match := fileNameRegex.FindStringSubmatch(path)
		if len(match) == 0 {
			return nil
		}
		idx := mapping[match[2]]
		if idx == 0 {
			log.Fatalln("unknown file: " + path)
		}
		if fmt.Sprintf("%02d", idx) != match[1] {
			// rename
			log.Println("Rename", path)
			os.Rename(path, fmt.Sprintf("%02d-%s.go", idx, match[2]))
		}
		foundMapping[match[2]] = true

		return nil
	})

	for filename, found := range foundMapping {
		newFilename := fmt.Sprintf("%02d-%s.go", mapping[filename], filename)
		if !found {
			log.Println("Create file", newFilename)
			os.Create(newFilename)
		}
	}
}

func normalizeName(name string) string {
	s := strings.ToLower(name)
	s = strings.Replace(s, " / ", "-", -1)
	s = strings.Replace(s, " ", "-", -1)
	s = strings.Replace(s, "/", "-", -1)
	return s
}
