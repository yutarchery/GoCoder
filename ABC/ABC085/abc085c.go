package ABC085

import "fmt"

func ABC085C() {
	var n, y int
	fmt.Scanf("%d %d", &n, &y)

	for i := 0; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			k := n - (i + j)

			if i+5*j+10*k == y/1000 {
				fmt.Printf("%d %d %d\n", k, j, i)
				return
			}
		}
	}
	fmt.Println("-1 -1 -1")
}
