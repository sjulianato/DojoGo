package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) float64 {

	//hint 1: Usar ciclo que itere 10 veces
	//hint 2: Usar una variable que se calcule en cada interación, 			empezar en el valor medio
	//hint3: delta 1e-6
	z := float64(x/2)
	counter := 0

	for delta:=z; math.Abs(delta) > 1e-6; {
		zi:= z
		z = z - (z*z-x)/(2*z)
		delta = z - zi
		counter++
		
	}
	fmt.Println(counter)
	return z;

}

func main() {
  fmt.Println(Sqrt(2));
  fmt.Println(math.Sqrt(2));	
}
