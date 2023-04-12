package ABC114A

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ABC114A() {
	/** introduce scanner */
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	/** start program */
	x := scanInt(scanner)

	if x == 3 || x == 5 || x == 7 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
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
