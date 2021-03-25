package challenge_04

import (
	"fmt"
	"strings"
)

func Triangle(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(strings.Repeat("*", i))
	}
}
