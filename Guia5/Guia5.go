package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ejercicio1() {
	var matriz [4][4]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			matriz[i][j] = 0
		}
	}
	fmt.Println("Ejercicio 1:")
	for i := 0; i < 4; i++ {
		fmt.Println(matriz[i])
	}
}

func ejercicio2() {
	var matriz [4][4]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i+j == 3 {
				matriz[i][j] = 1
			} else {
				matriz[i][j] = 0
			}
		}
	}
	fmt.Println("Ejercicio 2:")
	for i := 0; i < 4; i++ {
		fmt.Println(matriz[i])
	}
}

func ejercicio3() {
	var notas [40][5]float64
	for i := 0; i < 40; i++ {
		for j := 0; j < 5; j++ {
			notas[i][j] = 7
		}
	}
	fmt.Println("Ejercicio 3:")
	for i := 0; i < 40; i++ {
		suma := 0.0
		for j := 0; j < 5; j++ {
			suma += notas[i][j]
		}
		fmt.Printf("Alumno %d: Promedio = %.2f\n", i+1, suma/5)
	}
}

func ejercicio4() {
	var CANT [50][15]int
	var TOTAL [15]int
	for i := 0; i < 50; i++ {
		for j := 0; j < 15; j++ {
			CANT[i][j] = 1
			TOTAL[j] += CANT[i][j]
		}
	}
	maxVenta := TOTAL[0]
	vendedor := 1
	for j := 1; j < 15; j++ {
		if TOTAL[j] > maxVenta {
			maxVenta = TOTAL[j]
			vendedor = j + 1
		}
	}
	fmt.Println("Ejercicio 4:")
	fmt.Printf("El vendedor con mayor venta es el número %d con %d ventas\n", vendedor, maxVenta)
}

func ejercicio5() {
	var tabla = [4][4]float64{
		{8, 8, 7, 0},
		{7, 9, 10, 0},
		{10, 9, 5, 0},
		{4, 9, 8.5, 0},
	}
	for i := 0; i < 4; i++ {
		tabla[i][3] = (tabla[i][0] + tabla[i][1] + tabla[i][2]) / 3
	}
	fmt.Println("Ejercicio 5:")
	fmt.Println("         Nota 1  Nota 2  Nota 3  Prom.")
	for i := 0; i < 4; i++ {
		fmt.Printf("Alumno %d: %5.2f   %5.2f   %5.2f   %5.2f\n", i+1, tabla[i][0], tabla[i][1], tabla[i][2], tabla[i][3])
	}
}

func leerLinea() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func leerInt() int {
	for {
		text := leerLinea()
		val, err := strconv.Atoi(text)
		if err == nil {
			return val
		}
		fmt.Print("Ingrese un número válido: ")
	}
}

func main() {
	for {
		fmt.Println("\nSeleccione una opción:")
		fmt.Println("1 - Matriz 4x4 en cero")
		fmt.Println("2 - Matriz con diagonal secundaria en uno")
		fmt.Println("3 - Promedio de 40 alumnos")
		fmt.Println("4 - Vendedor con mayor venta")
		fmt.Println("5 - Tabla de notas y promedios")
		fmt.Println("0 - Salir")
		fmt.Print("Opción: ")

		opcion := leerInt()

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
