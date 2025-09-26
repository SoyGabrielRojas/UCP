package main

import (
	"fmt"
)

func ejercicio1() {
	var edad1, edad2 int
	fmt.Print("Ingrese la edad del hermano 1: ")
	fmt.Scan(&edad1)
	fmt.Print("Ingrese la edad del hermano 2: ")
	fmt.Scan(&edad2)
	if edad1 > edad2 {
		fmt.Printf("El hermano 1 es mayor por %d años.\n", edad1-edad2)
	} else if edad2 > edad1 {
		fmt.Printf("El hermano 2 es mayor por %d años.\n", edad2-edad1)
	} else {
		fmt.Println("Ambos hermanos tienen la misma edad.")
	}
}

func ejercicio2() {
	var n1, n2, n3 int
	fmt.Print("Ingrese las tres notas del alumno: ")
	fmt.Scan(&n1, &n2, &n3)
	prom := float64(n1+n2+n3) / 3.0
	aprobado := n1 > 4 && n2 > 4 && n3 > 4 && prom >= 7
	fmt.Println("¿Está aprobado?", aprobado)
}

func ejercicio3() {
	var l1, l2, l3 int
	fmt.Print("Ingrese los tres lados del triángulo: ")
	fmt.Scan(&l1, &l2, &l3)
	if l1 == l2 && l2 == l3 {
		fmt.Println("Tipo de triángulo: Equilátero")
	} else if l1 == l2 || l2 == l3 || l1 == l3 {
		fmt.Println("Tipo de triángulo: Isósceles")
	} else {
		fmt.Println("Tipo de triángulo: Escaleno")
	}
}

func ejercicio4() {
	var anio int
	fmt.Print("Ingrese un año: ")
	fmt.Scan(&anio)
	bisiesto := (anio%4 == 0 && anio%100 != 0) || (anio%400 == 0)
	fmt.Println("¿Es bisiesto?", bisiesto)
}

func ejercicio5() {
	var a, b int
	fmt.Print("Ingrese dos números para ordenar: ")
	fmt.Scan(&a, &b)
	if a < b {
		fmt.Printf("Ordenados: %d, %d\n", a, b)
	} else {
		fmt.Printf("Ordenados: %d, %d\n", b, a)
	}
}

func main() {
	for {
		fmt.Println("\nSeleccione una opción:")
		fmt.Println("1 - Comparar edades de hermanos")
		fmt.Println("2 - Verificar si un alumno está aprobado")
		fmt.Println("3 - Tipo de triángulo")
		fmt.Println("4 - Verificar si un año es bisiesto")
		fmt.Println("5 - Ordenar dos números")
		fmt.Println("0 - Salir")
		fmt.Print("Opción: ")

		var opcion int
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			ejercicio1()
		case 2:
			ejercicio2()
		case 3:
			ejercicio3()
		case 4:
			ejercicio4()
		case 5:
			ejercicio5()
		case 0:
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción inválida.")
		}
	}
}
