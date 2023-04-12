package ABC088

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Location struct {
	I int
	J int
}

func ABC088D() {
	diff := []Location{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	h_w := strings.Split(scanner.Text(), " ")
	h, _ := strconv.Atoi(h_w[0])
	w, _ := strconv.Atoi(h_w[1])

	s := make([]string, h)
	dist := make([][]int, h)
	ans := 0
	for i := range s {
		scanner.Scan()
		s[i] = scanner.Text()
		for j := range s[i] {
			if s[i][j] == '.' {
				ans++
			}
		}
		dist[i] = make([]int, w)
	}

	que := make([]Location, 0)
	que = append(que, Location{0, 0})
	dist[0][0] = 1

	for len(que) > 0 {
		q := que[0]
		que = que[1:]

		for k := range diff {
			next_i, next_j := q.I+diff[k].I, q.J+diff[k].J

			if 0 <= next_i && next_i < h && 0 <= next_j && next_j < w {
				if s[next_i][next_j] == '.' && dist[next_i][next_j] == 0 {
					dist[next_i][next_j] = dist[q.I][q.J] + 1
					que = append(que, Location{next_i, next_j})
				}
			}
		}
	}

	if dist[h-1][w-1] == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans - dist[h-1][w-1])
	}
}
