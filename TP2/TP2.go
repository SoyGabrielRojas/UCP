package main

import (
	"fmt"
)

const filas, columnas = 5, 10

var matrizEj1 [filas][columnas]int
var matrizEj1Completada bool

func ejercicio1() {
	for i := 0; i < filas; i++ {
		for j := 0; j < columnas; j++ {
			fmt.Printf("Ingrese el valor para [%d][%d]: ", i, j)
			fmt.Scan(&matrizEj1[i][j])
		}
	}

	max, min := matrizEj1[0][0], matrizEj1[0][0]
	maxPos, minPos := [2]int{0, 0}, [2]int{0, 0}
	for i := 0; i < filas; i++ {
		for j := 0; j < columnas; j++ {
			if matrizEj1[i][j] > max {
				max = matrizEj1[i][j]
				maxPos = [2]int{i, j}
			}
			if matrizEj1[i][j] < min {
				min = matrizEj1[i][j]
				minPos = [2]int{i, j}
			}
		}
	}
	fmt.Printf("Máximo: %d en posición [%d][%d]\n", max, maxPos[0], maxPos[1])
	fmt.Printf("Mínimo: %d en posición [%d][%d]\n", min, minPos[0], minPos[1])

	matrizEj1Completada = true
}

func ejercicio2() {
	if !matrizEj1Completada {
		fmt.Println("Primero debe completar el ejercicio 1 para poder ejecutar el ejercicio 2.")
		return
	}

	var traspuesta [columnas][filas]int
	for i := 0; i < filas; i++ {
		for j := 0; j < columnas; j++ {
			traspuesta[j][i] = matrizEj1[i][j]
		}
	}

	fmt.Println("Matriz traspuesta:")
	for i := 0; i < columnas; i++ {
		for j := 0; j < filas; j++ {
			fmt.Printf("%d ", traspuesta[i][j])
		}
		fmt.Println()
	}
}

func ejercicio3() {
	var p int
	fmt.Print("Ingrese el orden de la matriz cuadrada (misma cantidad de filas y columnas): ")
	fmt.Scan(&p)
	var matriz = make([][]int, p)
	for i := range matriz {
		matriz[i] = make([]int, p)
	}

	for i := 0; i < p; i++ {
		for j := 0; j < p; j++ {
			fmt.Printf("Ingrese el valor para [%d][%d]: ", i, j)
			fmt.Scan(&matriz[i][j])
		}
	}

	fmt.Print("(Elementos donde fila = columna) Diagonal principal: ")
	for i := 0; i < p; i++ {
		fmt.Printf("%d ", matriz[i][i])
	}
	fmt.Println()

	fmt.Print("(Elementos donde fila + columna = orden - 1) Diagonal secundaria: ")
	for i := 0; i < p; i++ {
		fmt.Printf("%d ", matriz[i][p-1-i])
	}
	fmt.Println()

	fmt.Print("(Elementos arriba de la diagonal principal) Triangular superior: ")
	for i := 0; i < p; i++ {
		for j := i + 1; j < p; j++ {
			fmt.Printf("%d ", matriz[i][j])
		}
	}
	fmt.Println()

	fmt.Print("(Elementos debajo de la diagonal principal) Triangular inferior: ")
	for i := 1; i < p; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("%d ", matriz[i][j])
		}
	}
	fmt.Println()
}

func ejercicio4() {
	const empleados = 3
	var codigos [empleados]int
	var horasNormales [empleados]int
	var horasExtras [empleados]int

	for i := 0; i < empleados; i++ {
		fmt.Printf("Ingrese el código del empleado %d (1-100): ", i+1)
		fmt.Scan(&codigos[i])
		fmt.Printf("Ingrese horas normales trabajadas: ")
		fmt.Scan(&horasNormales[i])
		fmt.Printf("Ingrese horas extras trabajadas: ")
		fmt.Scan(&horasExtras[i])
	}

	fmt.Println("\n Informe de empleados")
	for i := 0; i < empleados; i++ {
		pagoNormal := float64(horasNormales[i]) * 2350
		pagoExtra := float64(horasExtras[i]) * 3500
		total := pagoNormal + pagoExtra
		fmt.Printf("Código: %d | Horas normales: %d | $ normales: %.2f | Horas extras: %d | $ extras: %.2f | Total: %.2f\n",
			codigos[i], horasNormales[i], pagoNormal, horasExtras[i], pagoExtra, total)
	}
}

func main() {
	for {
		fmt.Println("\nSeleccione una opción:")
		fmt.Println("1 - Matriz 5x10: máximo, mínimo y posiciones")
		fmt.Println("2 - Matriz traspuesta (requiere completar ejercicio 1)")
		fmt.Println("3 - Diagonales y triangulares de matriz cuadrada")
		fmt.Println("4 - Informe de empleados La Huelga S.A.")
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
		case 0:
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción inválida.")
		}
	}
}
