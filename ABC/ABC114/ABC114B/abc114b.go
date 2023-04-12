package ABC114B

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ABC114B() {
	/** introduce scanner */
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	/** start program */
	s := scanString(scanner)
	ans := 100000000
	for i := 0; i+2 < len(s); i++ {
		num, _ := strconv.Atoi(s[i : i+3])

		if abs(num-753) < ans {
			ans = abs(num - 753)
		}
	}

	fmt.Println(ans)
}

func abs(n int) int {
	if n <= 0 {
		return -1 * n
	} else {
		return n
	}
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
