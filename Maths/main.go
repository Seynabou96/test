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

// ----------------Transformer chaque élément en nombre------------//

func ArrayData(arr []string) []int {
	var intArray []int
	for i := 0; i < len(arr); i++ {
		intArray = append(intArray, Trimatoi(arr[i]))
	}

	return intArray
}

//------------Obtenir le nombre entier le plus proche---------//
func Rounded(num int){
	
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

// ---------------Calcul puissance-----------//
func IterativePower(nb int, power int) int {
	iteration := 1
	if nb == 0 && power == 0 {
		return 1
	}
	if nb == 0 || power < 0 {
		return 0
	} else {
		for i := 1; i <= power; i++ {
			iteration *= nb
		}
		return iteration
	}
}

//-------------Variance------------//

func Variance(intArray []int) int {
	var variance, x int
	for i := 0; i < len(intArray); i++ {
		nb := intArray[i] - Average(intArray)
		x = x + IterativePower(nb, 2)
	}

	variance = x / 2

	return variance
}

// -----------Racine carré d'un nombre----------------//
func Sqrt(nb int) int {
	carre := 0
	if nb < 0 {
		return 0
	}
	if nb == 1 {
		return 1
	}
	for i := 1; i < nb; i++ {
		if i*i == nb {
			carre = i
		}
	}
	return carre
}

//-------------Ecart type----------------//

func StandardDeviation(variance int) int {

	ecartType := Sqrt(variance)

	return ecartType
}

//-----------Fonction qui s'occupe de calculer et d'afficher----------//

func MathsSkills(input string) {
	var average, median, variance, ecartType int
	dataString := ReadTheFile(input)
	// fmt.Println(dataString)
	/////---Mettre les données dans un tableau------/////
	dataArray := strings.Fields(dataString)
	// fmt.Println(dataArray)
	intArray := ArrayData(dataArray)
	//calculs
	average = Average(intArray)
	median = Median(intArray)
	variance = Variance(intArray)
	ecartType = StandardDeviation(variance)
	//Affichage
	fmt.Println("Average:", average)
	fmt.Println("Median:", median)
	fmt.Println("Variance:", variance)
	fmt.Println("Standard Deviation:", ecartType)

}

func main() {
	arg := os.Args

	if len(arg) == 2 {
		MathsSkills(arg[1])
	} else {
		fmt.Println("Sorry, something wrong with your command")
	}
}
