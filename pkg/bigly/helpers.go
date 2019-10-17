package bigly

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

// IsValid checks a filepath for validity
func IsValid(fp string) bool {
	// Check if file already exists
	if _, err := os.Stat(fp); err == nil {
		return true
	}

	// Attempt to create it
	var d []byte
	if err := ioutil.WriteFile(fp, d, 0644); err == nil {
		os.Remove(fp) // And delete it
		return true
	}

	return false
}

// ReadTrainer ...
func ReadTrainer(fp string) ([][]string, error) {
	var trainer [][]string

	file, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		trainer = append(trainer, strings.Split(scanner.Text(), " "))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return trainer, nil
}
