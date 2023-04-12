package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	/** introduce scanner */
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	/** start program */

}

/** inputs */
func scanInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func scanInt64(scanner *bufio.Scanner) int64 {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return int64(num)
}
