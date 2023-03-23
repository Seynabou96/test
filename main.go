package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Trimatoi(s string) int {
	signe := 1
	num := 0

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num *= 10
			num += int(s[i] - 48)
		} else if s[i] == '-' && num == 0 {
			signe = -1
		}
	}
	return num * signe
}

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
func IsAWorld(s string) bool {
	front := []string{"(cap)", "(low)", "(up)", "(low,", "(cap,", "(up,", "(hex)", "(bin)"}
	// index:= Index(tab,s)
	for i := 0; i < len(front); i++ {
		if s == front[i] {
			return false
		}
	}

	return true
}

func Permutation(s []string) []string {

	for i := 1; i < len(s); i++ {
		a := []string{"(cap)", "(low)", "(up)", "(low,"}
		p := []string{",", ".", ":"}
		// if i == 0{}
		if Contains(p, s[i]) && Contains(a, s[i-1]) {
			tampon := s[i]
			s[i] = s[i-1]
			s[i-1] = tampon

			// fmt.Println(s, 0)
			return Permutation(s)
		}
		if Contains(p, s[i]) && s[i-2] == "(low," {
			tampon := s[i]
			s[i] = s[i-1]
			s[i-1] = tampon

			return Permutation(s)
		}
	}

	return s

}

func ApplyChangeForPermutations(s string) string {
	p := []string{",", ".", ":", ";", "!", "?"}

	for i := 0; i < len(p); i++ {
		s = strings.ReplaceAll(s, " "+p[i], p[i])
	}
	// s = strings.ReplaceAll(s, " .", ".")

	return s
}

func Restructuration(s []string) []string {
	f := strings.Join(s, " ")
	s = strings.Fields(f)
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

func IsNumeric(s string) bool {
	for _, number := range s {
		if (number >= 0 && number <= 47) || (number >= 58 && number <= 127) {
			return false
		}
	}
	return true
}

func Index(s []string, e string) int {
	for i, v := range s {
		if v == e {
			return i
		}

	}
	return -1
}
func FindAnAlphaNumericalWord(s []string, e string) bool {
	// trouver l'index du mot
	index := Index(s, e)
	for i := 0; i < index; i++ {
		//appeler la fonction alphanumérique
		if !IsAWorld(s[i]) {
			return true
		}
	}
	return false
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

func VerifyFormat(s string) bool {
	// number, _ := strconv.Atoi(a[i+1]) (string, int)
	number := Trimatoi(s)
	//string du nombre
	numToString := strconv.Itoa(number)
	b := numToString + ")"

	return b == s

}

func TryIt(a []string) []string {
	var result []string
	front := []string{"(cap)", "(low)", "(up)", "(low,", "(cap,", "(up,", "(hex)", "(bin)"}

	for i := 0; i < len(a); i++ {
		if Contains(front, a[i]) && FindAnAlphaNumericalWord(a, a[i]) {
			// fmt.Println("i'am")

		} else if a[i] == "(cap)" {
			// a[i-1] = Capitalise(a[i-1])
			// fmt.Println(a[i], i, i-1, 3)
			// fmt.Println(a[i], a[i-1], result)
			fmt.Println(result[len(result)-1])

			index := Index(result, result[len(result)-1])
			result[index] = Capitalise(result[index])

		} else if a[i] == "(low)" {
			index := Index(result, result[len(result)-1])
			// fmt.Println(result[len(result)-1])
			result[index] = strings.ToLower(result[index])
			// fmt.Println(a[i], a[i-1], result)

		} else if a[i] == "(up)" {
			index := Index(result, result[len(result)-1])
			result[index] = strings.ToUpper(result[index])
			// fmt.Println(result[len(result)-1])

			// fmt.Println(a[i], a[i-1], result)

		} else if a[i] == "(low," && i != len(a)-1 {
			// number, _ := strconv.Atoi(a[i+1]) (string, int)
			number := Trimatoi(a[i+1])
			//string du nombre
			numToString := strconv.Itoa(number)
			b := numToString + ")"
			// fmt.Println(a[i+1], number, numToString, b)
			if a[i+1] == b {

				index := Index(result, result[len(result)-1])
				// fmt.Println(result[index], index)
				// fmt.Println(result)

				//string du nombre
				// numToString := strconv.Itoa(number)
				if number > 0 {
					if number < len(result) {
						for i := index; i >= len(result)-number; i-- {
							result[i] = strings.ToLower(result[i])

						}
					} else if number >= len(result) {
						for i := index; i >= 0; i-- {
							result[i] = strings.ToLower(result[i])

						}
					}
				}
				num := 2023
				a[i+1] = string(rune(num))
				// a = append(a[:i+1], a[i+2:]...)
				// mois = append(mois[:indexASupprimer], mois[(indexASupprimer+1):]...)
			}
		} else if (a[i] == "(low," || a[i] == "(cap," || a[i] == "(up,") && i == len(a)-1 {
		} else if a[i] == "ߧ" {

		} else {
			result = append(result, a[i])
		}
	}
	return result
}

func main() {
	// s := []string{"I", "TRY", "(low,", "2)", "(low,", "3)3"}
	// s := []string{",", "(cap)", "ect", ".", "try", "(up)", ".", "(low)", ".", "(cap)", "ect", ",", "basta", ",", "(cap)"}

	s := []string{"(low,", "(cap)", "(low)", "(up)", "ECT", ".", "try", "(up)", "(low,", "2)", ".", "(cap)", ".", "(up)", "eCt", ",", "baSta", ",", "(up)", "(low,", "2)", "(up)", "(low,", "(low,"}
	// fmt.Println(s)

	s = Permutation(s)
	// fmt.Println(s)
	f := strings.Join(s, " ")
	// fmt.Println(f)
	f = ApplyChangeForPermutations(f)
	// fmt.Println(f)
	s = strings.Fields(f)
	fmt.Println(s)

	s = TryIt(s)
	fmt.Println(s)
	fmt.Println(len(s))

	// fmt.Println(i)
	// fmt.Println(f)

}
