package main

import (
	"fmt"
	"math/big"
	"reflect"
	"Gooooo/playingaround"
)

func main() {

	// variable testing
	variablesAndPlayingAround()
	
	// loop testing 
	playingaround.ForTest()
}

func variablesAndPlayingAround() {
	fmt.Println("Hello World")

	name := "Rohit"

	var name2 string = "Kya"

	var myAge int = 1231231
	var sabji = "Ajnkjanskdjnak"

	println(sabji)


	var myMarks, myMathsMarks int = 100, 70

	fmt.Print(name)
	fmt.Println("WHy dear ehy")
	fmt.Println("Kya haal batcheetetetet")
	fmt.Print(name2)

	printMyAge(myAge)

	println(myMarks, myMathsMarks)

	var mybool bool
	println(mybool)


	var float32variable float32 = 123456792
	fmt.Printf("%.10f\n", float32variable)

	var float64variable float64 = 123456789123456789123456789
	fmt.Printf("%.10f\n", float64variable)


	var bigFloat = new(big.Float).SetPrec(200)
	bigFloat.SetString("123456789123456789123456789")
	fmt.Println("", bigFloat)

	var floatString = "123456789123456789123456789"
	fmt.Println(floatString)


	fmt.Println(reflect.TypeOf(floatString))
	fmt.Println(reflect.TypeOf(bigFloat))
	fmt.Println(reflect.TypeOf(123))
	fmt.Println(reflect.TypeOf(12123123.0))
	fmt.Println(reflect.TypeOf(floatString))
}

func printMyAge(myAge int) {
	fmt.Println(myAge)
}