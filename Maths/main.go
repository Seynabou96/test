package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
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

// ----------------Transformer chaque élément en nombre------------//

func ArrayData(arr []string) []float64 {
	float := 1.2
	var intArray []float64
	for i := 0; i < len(arr); i++ {
		float, _ = strconv.ParseFloat(arr[i], 64)
		intArray = append(intArray, float)
	}

	return intArray
}

//------------Calcul Moyenne---------------//

func Average(intArray []float64) float64 {

	var average, somme float64

	//calculer la somme
	for i := 0; i < len(intArray); i++ {
		somme = somme + intArray[i]
	}
	//calculer average
	average = somme / float64(len(intArray))

	return average
}

// ---------Renvoie la median si elle est paire------------//
func Abort(tabb []float64) float64 {
	var t float64
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
func Abort2(tabb []float64) float64 {
	var t float64
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
func Median(intArray []float64) float64 {
	var median float64
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

//-------------Variance------------//

func Variance(intArray []float64) float64 {
	var variance, x, xi float64
	for i := 0; i < len(intArray); i++ {
		nb := intArray[i] - Average(intArray)
		xi = math.Pow(nb, 2)
		x += xi
	}

	variance = x / float64(len(intArray))

	return variance
}

//-------------Ecart type----------------//

func StandardDeviation(variance float64) float64 {

	ecartType := math.Sqrt(variance)
	return ecartType
}

//-----------Fonction qui s'occupe de calculer et d'afficher----------//

func MathsSkills(input string) {
	var average, median, variance, ecartType float64
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
	vari0 := math.Round(variance)
	// variance1 := strconv.ParseInt(variance)
	vari := int(vari0)
	ecartType = StandardDeviation(variance)
	//Affichage
	fmt.Println("Average:", math.Round(average))
	fmt.Println("Median:", math.Round(median))
	fmt.Println("Variance:", vari)
	fmt.Println("Standard Deviation:", math.Round(ecartType))
}

func main() {
	arg := os.Args

	if len(arg) == 2 {
		MathsSkills(arg[1])
	} else {
		fmt.Println("Sorry, something wrong with your command")
	}
}
