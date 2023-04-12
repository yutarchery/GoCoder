package ABC100D

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func ABC100D() {
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	n := scanInt(scanner)
	m := scanInt(scanner)

	arr := make([][]int64, 8)
	for i := 0; i < n; i++ {
		x := make([]int64, 3)
		for j := range x {
			x[j] = scanInt64(scanner)
		}

		for k := 0; k < 8; k++ {
			cnt := 1
			sum := int64(0)
			for j := 0; j < 3; j++ {
				if (k/cnt)%2 == 0 {
					sum += x[j]
				} else {
					sum -= x[j]
				}
				cnt *= 2
			}
			arr[k] = append(arr[k], sum)
		}
	}

	ans := int64(0)
	for k := 0; k < 8; k++ {
		sort.Slice(arr[k], func(i, j int) bool {
			return arr[k][i] > arr[k][j]
		})

		sum := int64(0)
		for j := 0; j < m; j++ {
			sum += arr[k][j]
		}

		if ans < sum {
			ans = sum
		}
	}

	fmt.Println(ans)
}

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
