package ABC088

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ABC088B() {
	var scanner = bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	a_input := strings.Split(scanner.Text(), " ")
	a := make([]int, len(a_input))
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(a_input[i])
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})

	alice, bob := 0, 0
	for i, ai := range a {
		if i%2 == 0 {
			alice += ai
		} else {
			bob += ai
		}
	}

	fmt.Printf("%v\n", (alice - bob))
}
