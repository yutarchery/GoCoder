package ABC100

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ABC100C() {
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	ans := 0
	for i := 0; i < n; i++ {
		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())
		ans += cnt2(a)
	}

	fmt.Println(ans)
}

func cnt2(num int) int {
	ans := 0
	for num%2 == 0 {
		ans++
		num /= 2
	}
	return ans
}
