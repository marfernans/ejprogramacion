// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello World");
// }
//*****************************************

package main

import "fmt"

func main (){
	var numerador, denominador, residuo;

	numerador = 8
	denominador = 5
	residuo = numerador - denominador * (numerador / denominador)
	fmt.Println(residuo)
}