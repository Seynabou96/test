package main

import (
	"fmt"
)

func Apostrophe(s string) string {
	so := []rune(s)

	var l, last []rune

	num := 2023
	year := rune(num)

	count := 0
	for i := 0; i < len(so); i++ {
		if so[i] == 39 {
			fmt.Println(count)
			count++
		}
		if count%2 != 0 && so[i] == ' ' && so[i-1] == 39 {
			l = append(l, year)

		} else if count%2 == 0 && so[i] == 39 && i != 0 {
			// l = append(l, year)
			// so[i-1] = year
			l[i-1] = year
			l = append(l, so[i])

			fmt.Println("yes")

		} else {
			l = append(l, so[i])

		}
	}

	for _, v := range l {
		if v != year {
			last = append(last, v)
		}
	}
	return string(last)
}

// 	} else {
// 		count = 0
// 	}

// 	if count == 1 && so[i] >= 'a' && so[i] <= 'z' {
// 		so[i] -= 32

// 	} else if count > 1 && so[i] >= 'A' && so[i] <= 'Z' {
// 		so[i] += 32

// 	}
// }

func main() {
	s := "i am exactly how they describe me:  ' again ' ' "

	// s = strings.ReplaceAll(s, " ' ", " '")
	fmt.Println(s)
	// tab := strings.Fields(s)
	// fmt.Println(tab)

	fmt.Println(Apostrophe(s))

}
