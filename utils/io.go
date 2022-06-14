package utils

import (
	"bufio"
	"os"
)

func WriteLinesList(file string, lines []string) error {
	f, err := os.Create(file)
	defer f.Close()
	if err != nil {
		return err
	}
	w := bufio.NewWriter(f)
	for _, l := range lines {
		_, err := w.WriteString(l + "\n")
		if err != nil {
			return err
		}
	}
	err = w.Flush()
	if err != nil {
		return err
	}
	return nil
}

func ReadLinesList(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return []string{}, err
	}
	s := bufio.NewScanner(f)

	s.Split(bufio.ScanLines)

	var lines []string = []string{}
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines, err
}
