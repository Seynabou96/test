package main

import (
	"fmt"
	"strings"
)

func Contains(a []string, element string) bool {
	// f[]string
	for i := 0; i < len(a); i++ {
		if a[i] == element {
			return true
		}
	}
	return false
}

func IsAlpha(s string) bool {
	for _, letter := range s {
		if (letter >= 0 && letter <= 31) || (letter >= 33 && letter <= 47) || (letter >= 58 && letter <= 64) || (letter >= 91 && letter <= 96) || (letter >= 123 && letter <= 127) {
			return false
		}
	}
	return true
}

func Permutation(s []string) []string {

	for i := 1; i < len(s); i++ {
		a := []string{"(cap)", "(low)", "(up)"}
		p := []string{",", ".", ":"}
		// if i == 0{}
		if Contains(a, s[i]) && Contains(p, s[i-1]) {
			tampon := s[i]
			s[i] = s[i-1]
			s[i-1] = tampon

			// fmt.Println(s, 0)
			return Permutation(s)
		}
	}

	return s

}

//	func Index(s string, toFind string) int {
//		i := []rune(s)
//		j := []rune(toFind)
//		a := len(i)
//		b := len(j)
//		for i := 0; i < a-b; i++ {
//			if toFind == s[i:i+b] {
//				return i
//			}
//		}
//		return -1
//	}
func Index(s []string, e string) int {
	for i, v := range s {
		if v == e {
			return i
		}

	}
	return -1
}

func TryIt(a []string) []string {
	var result []string
	for i := 0; i < len(a); i++ {
		if (a[i] == "(cap)" || a[i] == "(low)" || a[i] == "(up)") && i == 0 {
			fmt.Println(a[i])
			a[i] = ""
		} else if a[i] == "(cap)" {
			// a[i-1] = Capitalise(a[i-1])
			index := Index(result, result[len(result)-1])
			result[index] = Capitalise(result[index])

		} else if a[i] == "(low)" {
			index := Index(result, result[len(result)-1])
			result[index] = strings.ToLower(result[index])

		} else if a[i] == "(up)" {
			index := Index(result, result[len(result)-1])
			result[index] = strings.ToUpper(result[index])

		} else {
			result = append(result, a[i])
		}
	}
	return result
}

func Capitalise(s string) string {

	so := []rune(s)
	count := 0
	for i := 0; i < len(s); i++ {
		if (so[i] >= 'a' && so[i] <= 'z') || (so[i] >= 'A' && so[i] <= 'Z') || (so[i] >= '0' && so[i] <= '9') {
			count++
		} else {
			count = 0
		}

		if count == 1 && so[i] >= 'a' && so[i] <= 'z' {
			so[i] -= 32

		} else if count > 1 && so[i] >= 'A' && so[i] <= 'Z' {
			so[i] += 32

		}
	}

	return string(so)
}

func main() {
	// s := []string{",", "(cap)", "ect", ".", "try", "(cap)", ".", "(low)", ".", "(up)", "ect", ",", "basta", ",", "(cap)"}

	s := []string{"(cap)", "ect", ".", "TRY", "(cap)", ".", "(up)", ".", "(low)", "ect", ",", "basta", ",", "(cap)"}
	fmt.Println(s)

	s = Permutation(s)
	fmt.Println(s)

	s = TryIt(s)
	fmt.Println(s)
	// fmt.Println(i)
	// fmt.Println(f)

}
