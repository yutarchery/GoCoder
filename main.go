package main

import (
	"bufio"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	/** introduce scanner */
	scanner.Split(bufio.ScanWords)

	/** start program */

}

/** inputs */
func scanInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func scanInt64() int64 {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return int64(num)
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}
