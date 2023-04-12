package ABC114C

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ABC114C() {
	/** introduce scanner */
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	/** start program */
	n := scanInt(scanner)
	fmt.Println(solve(0, n, false, false, false))
}

func solve(now, max int, three, five, seven bool) int {
	if now > max {
		return 0
	}

	ans := 0
	if three && five && seven {
		ans++
	}

	return ans +
		solve(now*10+3, max, true, five, seven) +
		solve(now*10+5, max, three, true, seven) +
		solve(now*10+7, max, three, five, true)
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

func scanString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
