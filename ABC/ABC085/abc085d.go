package ABC085

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Attack struct {
	Value      int64
	Repeatable bool
}

func ABC085D() {
	var scanner = bufio.NewScanner(os.Stdin)

	var n int
	var h int64
	fmt.Scanf("%d %d", &n, &h)

	attacks := make([]Attack, 0)
	for i := 0; i < n; i++ {
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		a, _ := strconv.Atoi(input[0])
		b, _ := strconv.Atoi(input[1])

		attacks = append(attacks, Attack{int64(a), true})
		attacks = append(attacks, Attack{int64(b), false})
	}
	sort.Slice(attacks, func(i, j int) bool {
		return attacks[i].Value > attacks[j].Value
	})

	ans := int64(0)
	for _, attack := range attacks {
		if attack.Repeatable {
			ans += (h-1)/attack.Value + 1
			break
		} else {
			ans++
			h -= attack.Value
			if h <= 0 {
				break
			}
		}
	}

	fmt.Println(ans)
}
