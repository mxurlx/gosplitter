package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func mergeFiles(inputDir, outputFile string) error {
	files, err := os.ReadDir(inputDir)
	if err != nil {
		return err
	}

	var sortedFiles []string
	for _, file := range files {
		if !file.IsDir() {
			sortedFiles = append(sortedFiles, fmt.Sprintf("%s/%s", inputDir, file.Name()))
		}
	}

	sort.Slice(sortedFiles, func(i, j int) bool {
		numI := extractChunkNumber(filepath.Base(sortedFiles[i]))
		numJ := extractChunkNumber(filepath.Base(sortedFiles[j]))
		return numI < numJ
	})

	outFileHandle, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFileHandle.Close()

	for _, inputFile := range sortedFiles {
		inpHandle, err := os.Open(inputFile)
		if err != nil {
			return err
		}
		defer inpHandle.Close()
		_, err = io.Copy(outFileHandle, inpHandle)
		if err != nil {
			return err
		}
	}

	return nil
}

func extractChunkNumber(filename string) int {
	re := regexp.MustCompile(`_(\d+)`)
	match := re.FindStringSubmatch(filename)
	if len(match) > 1 {
		num, _ := strconv.Atoi(match[1])
		return num
	}

	return 0
}

func Merge(flags map[string]any, mandatoryArgs []string) error {
	input := mandatoryArgs[0]
	output := flags["output"].(string)

	if flags["output"] == "" {
		inputDir, err := os.ReadDir(input)
		if err != nil {
			return err
		}
		for _, file := range inputDir {
			if !file.IsDir() {
				output = strings.Split(file.Name(), "_")[0]
				break
			}
		}
	}

	return mergeFiles(input, output)
}
