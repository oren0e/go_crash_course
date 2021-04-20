package main

import (
	"fmt"
	"math"

	"github.com/oren0e/go_crash_course/03_packages/str_util"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(str_util.Reverse("olleH"))
}
