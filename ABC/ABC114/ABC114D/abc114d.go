package ABC114D

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)
var cnt = make([]int, 100)

func ABC114D() {
	/** introduce scanner */
	scanner.Split(bufio.ScanWords)

	/** start program */
	n := scanInt()

	for i := 2; i <= n; i++ {
		prime(i)
	}
	fmt.Println(solve())
}

func solve() int {
	ans := 0

	for i := 1; i < 100; i++ {
		if cnt[i] >= 74 {
			ans++
		}
		for j := 1; j < 100; j++ {
			if i == j {
				continue
			}
			if cnt[i] >= 4 && cnt[j] >= 14 {
				ans++
			}
			if cnt[i] >= 2 && cnt[j] >= 24 {
				ans++
			}

			for k := j + 1; k < 100; k++ {
				if i == k || j == k {
					continue
				}

				if cnt[i] >= 2 && cnt[j] >= 4 && cnt[k] >= 4 {
					ans++
				}
			}
		}
	}

	return ans
}

func prime(num int) {
	rest := num
	for i := 2; i*i <= num; i++ {
		for rest%i == 0 {
			cnt[i]++
			rest /= i
		}
	}
	if rest > 1 {
		cnt[rest]++
	}
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
