package gocp

import (
	"bufio"
	"os"
)

func Scan() (value string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	value = scanner.Text()

	return
}
