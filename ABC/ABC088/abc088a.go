package ABC088

import "fmt"

func ABC088A() {
	var n, a int
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &a)

	if n%500 <= a {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
