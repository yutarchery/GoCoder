package ABC100

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ABC100A() {
	var scanner = bufio.NewScanner(os.Stdin)

	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	a, _ := strconv.Atoi(input[0])
	b, _ := strconv.Atoi(input[1])

	if a <= 8 && b <= 8 {
		fmt.Println("Yay!")
	} else {
		fmt.Println(":(")
	}
}
