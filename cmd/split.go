package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func splitFileByBytes(inputFile, outputFile, suffix string, chunkSize int64) error {
	file, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	chunkNumber := 1
	buffer := make([]byte, chunkSize)

	_, err = os.Stat(outputFile)
	if err == nil {
		err = os.RemoveAll(outputFile)
		if err != nil {
			return err
		}
	}

	err = os.MkdirAll(outputFile, 0755)
	if err != nil {
		return err
	}

	for {
		bytesRead, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}
		if bytesRead == 0 {
			break
		}

		outChunkName := fmt.Sprintf("%s/%s_%s%03d.chunk", outputFile, inputFile, suffix, chunkNumber)
		outChunkFile, err := os.Create(outChunkName)
		if err != nil {
			return err
		}
		defer outChunkFile.Close()

		_, err = outChunkFile.Write(buffer[:bytesRead])
		if err != nil {
			return err
		}

		chunkNumber++
	}

	return nil
}

func Split(flags map[string]any, mandatoryArgs []string) error {
	input := mandatoryArgs[0]
	output := flags["output"].(string)
	suffix := flags["suffix"].(string)
	chunkSize := int64(flags["chunksize"].(int))

	if flags["output"] == "" {
		output = strings.Split(input, ".")[0]
	}
	return splitFileByBytes(input, output, suffix, chunkSize)
}
