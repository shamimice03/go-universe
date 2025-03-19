package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(inputFile, OutputFile string) *FileManager {
	return &FileManager{
		InputFilePath:  inputFile,
		OutputFilePath: OutputFile,
	}
}

func (fm FileManager) ReadFileContents() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		file.Close()
		return nil, errors.New("file not found")
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("failed to read file contents")
	}

	file.Close()
	return lines, nil
}

func (fm FileManager) WriteJSON(data any) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		file.Close()
		return errors.New("failed to write file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("failed to encode data")
	}

	file.Close()
	return nil
}
