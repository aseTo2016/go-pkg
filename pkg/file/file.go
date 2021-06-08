package file

import (
	"bufio"
	"io"
	"os"
)

//ReadFileLine read from file by line
func ReadFileLine(fileName string, do func(oneLine string) (isQuit bool)) error {
	fd, err := os.Open(fileName)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(fd)

	for {
		oneLine, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		isQuit := do(string(oneLine))
		if isQuit {
			return nil
		}
	}
}
