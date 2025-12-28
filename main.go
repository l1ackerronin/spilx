package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

func exitWithUsage() {
	fmt.Fprintf(os.Stderr, `
Usage:
  splitter <file> <lines_per_file>
  cat <file> | splitter <lines_per_file>

Examples:
  splitter hackerone.txt 10000
  cat DoD_Scope.txt | splitter 20000
`)
	os.Exit(1)
}

func main() {
	var (
		input       io.Reader
		baseName    string
		linesPerFile int
		err         error
	)

	args := os.Args[1:]

	// stdin mode
	if len(args) == 1 {
		linesPerFile, err = strconv.Atoi(args[0])
		if err != nil || linesPerFile <= 0 {
			exitWithUsage()
		}
		input = os.Stdin
		baseName = "stdin"
	} else if len(args) == 2 {
		filePath := args[0]
		linesPerFile, err = strconv.Atoi(args[1])
		if err != nil || linesPerFile <= 0 {
			exitWithUsage()
		}

		file, err := os.Open(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		input = file
		baseName = filepath.Base(filePath)
		baseName = baseName[:len(baseName)-len(filepath.Ext(baseName))]
	} else {
		exitWithUsage()
	}

	scanner := bufio.NewScanner(input)
	part := 1
	lineCount := 0

	var outFile *os.File
	var writer *bufio.Writer

	createNewFile := func() {
		if outFile != nil {
			writer.Flush()
			outFile.Close()
		}

		filename := fmt.Sprintf("%s_part_%d.txt", baseName, part)
		f, err := os.Create(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
			os.Exit(1)
		}

		outFile = f
		writer = bufio.NewWriter(f)
		fmt.Printf("[+] Created %s\n", filename)
		part++
		lineCount = 0
	}

	createNewFile()

	for scanner.Scan() {
		writer.WriteString(scanner.Text() + "\n")
		lineCount++

		if lineCount >= linesPerFile {
			createNewFile()
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Read error: %v\n", err)
	}

	if writer != nil {
		writer.Flush()
	}
	if outFile != nil {
		outFile.Close()
	}

	fmt.Println("Split completed successfully")
}

