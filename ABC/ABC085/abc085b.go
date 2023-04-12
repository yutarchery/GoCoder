package ABC085

import (
	"fmt"
	"sort"
)

func ABC085B() {
	var n int
	fmt.Scanf("%d", &n)
	d := make([]int, n)
	for i := range d {
		fmt.Scanf("%d", &d[i])
	}
	sort.Slice(d, func(i, j int) bool {
		return d[i] < d[j]
	})

	ans := 1
	for i := 1; i < len(d); i++ {
		if d[i-1] < d[i] {
			ans++
		}
	}
	fmt.Println(ans)
}
