package ABC085A

import (
	"bufio"
	"fmt"
	"os"
)

func ABC085A() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	s := scanner.Text()

	fmt.Println("2018" + s[4:])
}
