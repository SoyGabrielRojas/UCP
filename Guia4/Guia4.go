package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Ejercico1() {
	var n int
	fmt.Print("Ingrese la cantidad de notas: ")
	fmt.Scan(&n)

	notas := make([]float64, n)
	var suma float64
	var max float64

	for i := 0; i < n; i++ {
		fmt.Printf("Ingrese la nota %d: ", i+1)
		fmt.Scan(&notas[i])
		if i == 0 || notas[i] > max {
			max = notas[i]
		}
		suma += notas[i]
	}

	promedio := suma / float64(n)
	fmt.Printf("La nota más alta es: %.2f\n", max)
	fmt.Printf("El promedio de notas es: %.2f\n", promedio)
}

func Ejercicio2() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\nIngrese la cantidad de notas: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	n, _ := strconv.Atoi(input)

	notas := make([]float64, n)

	for i := range n {
		fmt.Printf("Ingrese la nota %d: ", i+1)
		notaStr, _ := reader.ReadString('\n')
		notaStr = strings.TrimSpace(notaStr)
		nota, _ := strconv.ParseFloat(notaStr, 64)
		notas[i] = nota
	}

	aprobados := 0
	desaprobados := 0

	for _, nota := range notas {
		if nota >= 6 {
			aprobados++
		} else {
			desaprobados++
		}
	}

	fmt.Printf("Aprobados: %d\n", aprobados)
	fmt.Printf("Desaprobados: %d\n", desaprobados)
}

func Ejercicio3() {
	const empleados = 10
	sueldos := make([]float64, empleados)
	reader := bufio.NewReader(os.Stdin)

	for i := range empleados {
		fmt.Printf("Ingrese el sueldo del empleado %d: ", i+1)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		sueldo, _ := strconv.ParseFloat(input, 64)
		sueldos[i] = sueldo
	}

	fmt.Println("\nSueldos de empleados:")
	for i, sueldo := range sueldos {
		fmt.Printf("Empleado %d: $%.2f\n", i+1, sueldo)
	}

	maxSueldo := sueldos[0]
	for _, sueldo := range sueldos {
		if sueldo > maxSueldo {
			maxSueldo = sueldo
		}
	}
	fmt.Printf("Mayor sueldo: $%.2f\n", maxSueldo)
}

func Ejercicio4() {
	const personas = 20
	edades := make([]int, personas)
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < personas; i++ {
		fmt.Printf("Ingrese la edad de la persona %d: ", i+1)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		edad, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Edad inválida, intente de nuevo.")
			i--
			continue
		}
		edades[i] = edad
	}

	var suma int
	menor := edades[0]
	for _, edad := range edades {
		suma += edad
		if edad < menor {
			menor = edad
		}
	}
	promedio := float64(suma) / float64(personas)
	fmt.Printf("\nPromedio de edad: %.2f\n", promedio)
	fmt.Printf("Edad más joven: %d\n", menor)
}

func Ejercicio5() {
	const vendedoresCount = 15
	vendedores := make([]string, vendedoresCount)
	ventas := make([]float64, vendedoresCount)
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < vendedoresCount; i++ {
		fmt.Printf("Ingrese el nombre del vendedor %d: ", i+1)
		nombre, _ := reader.ReadString('\n')
		nombre = strings.TrimSpace(nombre)
		fmt.Printf("Ingrese la venta en dólares de %s: ", nombre)
		ventaStr, _ := reader.ReadString('\n')
		ventaStr = strings.TrimSpace(ventaStr)
		venta, _ := strconv.ParseFloat(ventaStr, 64)
		vendedores[i] = nombre
		ventas[i] = venta
	}

	indiceMax, indiceMin := 0, 0
	for i := 1; i < vendedoresCount; i++ {
		if ventas[i] > ventas[indiceMax] {
			indiceMax = i
		}
		if ventas[i] < ventas[indiceMin] {
			indiceMin = i
		}
	}

	mayorPesos := ventas[indiceMax] * 140
	menorPesos := ventas[indiceMin] * 140

	fmt.Printf("\nMayor venta: Vendedor %s, $%.2f USD ($%.2f ARS)\n", vendedores[indiceMax], ventas[indiceMax], mayorPesos)
	fmt.Printf("Menor venta: Vendedor %s, $%.2f USD ($%.2f ARS)\n", vendedores[indiceMin], ventas[indiceMin], menorPesos)
}

func Ejercicio6() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nIngrese la cantidad de productos: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	n, _ := strconv.Atoi(input)

	cantidades := make([]int, n)
	costos := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Ingrese la cantidad del producto %d: ", i+1)
		cantStr, _ := reader.ReadString('\n')
		cantStr = strings.TrimSpace(cantStr)
		cantidad, _ := strconv.Atoi(cantStr)
		cantidades[i] = cantidad

		fmt.Printf("Ingrese el costo del producto %d: ", i+1)
		costoStr, _ := reader.ReadString('\n')
		costoStr = strings.TrimSpace(costoStr)
		costo, _ := strconv.ParseFloat(costoStr, 64)
		costos[i] = costo
	}

	var precioTotal float64
	for i := 0; i < n; i++ {
		subtotal := float64(cantidades[i]) * costos[i]
		precioTotal += subtotal
		if subtotal > 1000 {
			fmt.Printf("Producto %d supera $1000: subtotal = $%.2f\n", i+1, subtotal)
		}
	}
	fmt.Printf("Precio total de todos los productos: $%.2f\n", precioTotal)
}

func Ejercicio7() {
	const camionesCount = 3
	type Camion struct {
		Patente string
		Chofer  string
		Carga   string
		Hora    string
	}
	camiones := make([]Camion, camionesCount)
	contadorTe := 0
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < camionesCount; i++ {
		fmt.Printf("\nIngrese la patente del camión %d: ", i+1)
		patente, _ := reader.ReadString('\n')
		patente = strings.TrimSpace(patente)

		fmt.Print("Ingrese el nombre y apellido del chofer: ")
		chofer, _ := reader.ReadString('\n')
		chofer = strings.TrimSpace(chofer)

		fmt.Print("Ingrese el tipo de carga (madera/yerba/té): ")
		carga, _ := reader.ReadString('\n')
		carga = strings.TrimSpace(strings.ToLower(carga))

		fmt.Print("Ingrese la hora de egreso: ")
		hora, _ := reader.ReadString('\n')
		hora = strings.TrimSpace(hora)

		camiones[i] = Camion{
			Patente: patente,
			Chofer:  chofer,
			Carga:   carga,
			Hora:    hora,
		}

		if carga == "té" || carga == "te" {
			contadorTe++
		}

		if hora < "00:00" || hora > "23:59" {
			fmt.Println("Hora inválida, debe estar entre 00:00 y 23:59.")
			i--
			continue
		}
	}

	fmt.Println("\n--- Datos de los camiones ---")
	for _, camion := range camiones {
		fmt.Printf("Patente: %s, Chofer: %s, Carga: %s, Hora egreso: %s\n",
			camion.Patente, camion.Chofer, camion.Carga, camion.Hora)
	}
	fmt.Printf("\nCantidad de camiones que transportan té: %d\n", contadorTe)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go <numero-ejercicio>")
		return
	}

	opcion := os.Args[1]

	switch opcion {
	case "1":
		Ejercico1()
	case "2":
		Ejercicio2()
	case "3":
		Ejercicio3()
	case "4":
		Ejercicio4()
	case "5":
		Ejercicio5()
	case "6":
		Ejercicio6()
	case "7":
		Ejercicio7()
	default:
		fmt.Println("Opción no válida. Usa 1-7.")
	}
}
