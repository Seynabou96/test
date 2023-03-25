package main

import (
	"fmt"
	"os"
	"strings"
)

// ------------ Créer une fonction pour lire un l'input -----------//

func ReadTheFile(input string) string {

	data, err := os.ReadFile(input)
	if err != nil {
		errFile := "ERROR: open " + string(input) + ": no such file or directory"
		fmt.Println(errFile)
		os.Exit(1)
	}
	return string(data)
}

//------------Transforme un string en nombre---------//

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

// ----------------Transformer haque élément en nombre------------//

func ArrayData(arr []string) []int {
	var intArray []int
	for i := 0; i < len(arr); i++ {
		intArray = append(intArray, Trimatoi(arr[i]))
	}

	return intArray
}

//------------Calcul Moyenne---------------//

func Average(intArray []int) int {

	var average, somme int

	//calculer la somme
	for i := 0; i < len(intArray); i++ {
		somme = somme + intArray[i]
	}
	//calculer average
	average = somme / len(intArray)

	return average
}

// ---------Renvoie la median si elle est paire------------//
func Abort(tabb []int) int {
	var t int
	index := len(tabb) / 2
	for i := 0; i < len(tabb); i++ {
		for j := 0; j < len(tabb)-1; j++ {
			if tabb[j] > tabb[j+1] {
				t = tabb[j+1]
				tabb[j+1] = tabb[j]
				tabb[j] = t
			}
		}
	}
	return tabb[index]
}

// ---------Renvoie la median si elle est impaire------------//
func Abort2(tabb []int) int {
	var t int
	index := len(tabb) / 2
	for i := 0; i < len(tabb); i++ {
		for j := 0; j < len(tabb)-1; j++ {
			if tabb[j] > tabb[j+1] {
				t = tabb[j+1]
				tabb[j+1] = tabb[j]
				tabb[j] = t
			}
		}
	}
	return (tabb[index-1] + tabb[index]) / 2
}

// --------------Median-----------------//
func Median(intArray []int) int {
	var median int
	// si impair
	if len(intArray)%2 != 0 {
		median1 := Abort(intArray)
		median = median1
	} else {
		// si pair
		median1 := Abort2(intArray)
		median = median1

	}

	return median
}

//-----------Fonction qui s'occupe de calculer et d'afficher----------//

func MathsSkills(input string) {
	var average, median int
	dataString := ReadTheFile(input)
	// fmt.Println(dataString)
	/////---Mettre les données dans un tableau------/////
	dataArray := strings.Fields(dataString)
	// fmt.Println(dataArray)
	intArray := ArrayData(dataArray)
	//calculs
	average = Average(intArray)
	median = Median(intArray)
	//Affichage
	fmt.Println("Average:", average)
	fmt.Println("Median:", median)

}

func main() {
	arg := os.Args

	if len(arg) == 2 {
		MathsSkills(arg[1])
	} else {
		fmt.Println("Sorry, something wrong with your command")
	}
}
