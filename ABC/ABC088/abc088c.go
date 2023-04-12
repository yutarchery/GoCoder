package ABC088

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ABC088C() {
	var scanner = bufio.NewScanner(os.Stdin)

	c := make([][]int, 3)
	for i := range c {
		scanner.Scan()
		ci := strings.Split(scanner.Text(), " ")
		c[i] = make([]int, len(ci))
		for j, cij := range ci {
			c[i][j], _ = strconv.Atoi(cij)
		}
	}

	flag := true
	for i := 0; i < 2; i++ {
		for j := i + 1; j < 3; j++ {
			if !(c[i][0]-c[j][0] == c[i][1]-c[j][1] && c[i][1]-c[j][1] == c[i][2]-c[j][2]) {
				flag = false
			}
			if !(c[0][i]-c[0][j] == c[1][i]-c[1][j] && c[1][i]-c[1][j] == c[2][i]-c[2][j]) {
				flag = false
			}
		}
	}

	if flag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
