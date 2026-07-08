package logs

import (
	"bufio"
	"fmt"
	"os"
)

func View() error {

	file, err := os.Open(LogFile)
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return scanner.Err()
}
