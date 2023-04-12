package ABC100B

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ABC100B() {
	var scanner = bufio.NewScanner(os.Stdin)

	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	d, _ := strconv.Atoi(input[0])
	n, _ := strconv.Atoi(input[1])

	cnt, now := 0, 1
	for {
		if times(now) == d {
			cnt++
			if cnt == n {
				fmt.Println(now)
				return
			}
		}
		now++
	}

}

func times(num int) int {
	cnt := 0
	for num%100 == 0 {
		cnt++
		num /= 100
	}
	return cnt
}
