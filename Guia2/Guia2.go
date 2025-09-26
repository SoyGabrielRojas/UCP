package guia2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Ejercicio1() {
	fmt.Print("Ingrese la categoría del empleado (1-Repositor, 2-Cajero, 3-Supervisor): ")
	categoria := leerInt()
	var sueldoBruto float64

	switch categoria {
	case 1:
		sueldoBruto = 32335
	case 2:
		sueldoBruto = 38630.89
	case 3:
		sueldoBruto = 45560.20
	default:
		fmt.Println("Categoría inválida.")
		return
	}

	descuentoJubilacion := sueldoBruto * 0.11
	descuentoObraSocial := sueldoBruto * 0.03
	sueldoNeto := sueldoBruto - descuentoJubilacion - descuentoObraSocial

	fmt.Printf("Sueldo bruto: $%.2f\n", sueldoBruto)
	fmt.Printf("Descuento Jubilación (11%%): $%.2f\n", descuentoJubilacion)
	fmt.Printf("Descuento Obra Social (3%%): $%.2f\n", descuentoObraSocial)
	fmt.Printf("Sueldo Neto: $%.2f\n", sueldoNeto)

	if categoria == 1 {
		bono := sueldoBruto * 0.08
		fmt.Printf("Bono en compras (8%%): $%.2f\n", bono)
	}
}

func Ejercicio2() {
	fmt.Print("\nIngrese la cantidad de paquetes comprados: ")
	cantidad := leerInt()

	if cantidad < 5 {
		fmt.Println("No se permiten compras inferiores a 5 productos.")
	} else if cantidad <= 15 {
		fmt.Println("El costo de envío es de $200.")
	} else {
		fmt.Println("El envío es gratuito.")
	}
}

func Ejercicio3() {
	var precios, kgs [3]float64
	for i := 0; i < 3; i++ {
		fmt.Printf("\nIngrese el precio por Kg del producto %d: ", i+1)
		precios[i] = leerFloat()
		fmt.Printf("Ingrese la cantidad en Kg del producto %d: ", i+1)
		kgs[i] = leerFloat()
	}

	var totales [3]float64
	for i := 0; i < 3; i++ {
		totales[i] = precios[i] * kgs[i]
		fmt.Printf("\nMonto del producto %d: $%.2f\n", i+1, totales[i])
	}

	totalCompra := totales[0] + totales[1] + totales[2]
	fmt.Printf("Total de la compra: $%.2f\n", totalCompra)

	if totalCompra > 100 {
		descuento := totalCompra * 0.10
		nuevoTotal := totalCompra - descuento
		fmt.Printf("Se aplica un descuento del 10%% ($%.2f).\n", descuento)
		fmt.Printf("Nuevo total a pagar: $%.2f\n", nuevoTotal)
	}
}

func Ejercicio4() {
	fmt.Print("\nIngrese el DNI del empleado: ")
	dni := leerLinea()
	fmt.Print("Ingrese la categoría (0-Maestranza, 1-Administración, 2-Gerencia): ")
	cat := leerInt()

	var nombreCat string
	var bruto, descJub, descOS, descClub float64

	switch cat {
	case 0:
		nombreCat = "Maestranza"
		bruto = 23600
		descJub = 0.11
		descOS = 0.03
		descClub = 0.0
	case 1:
		nombreCat = "Administración"
		bruto = 35800
		descJub = 0.11
		descOS = 0.05
		descClub = 0.0
	case 2:
		nombreCat = "Gerencia"
		bruto = 60420
		descJub = 0.11
		descOS = 0.05
		descClub = 0.04
	default:
		fmt.Println("Categoría inválida.")
		return
	}

	descuentoTotal := bruto * (descJub + descOS + descClub)
	sueldoNeto := bruto - descuentoTotal

	fmt.Printf("\nDNI: %s\n", dni)
	fmt.Printf("Categoría: %s\n", nombreCat)
	fmt.Printf("Sueldo Bruto: $%.2f\n", bruto)
	fmt.Printf("Descuento total: $%.2f\n", descuentoTotal)
	fmt.Printf("Sueldo Neto: $%.2f\n", sueldoNeto)
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

func leerFloat() float64 {
	for {
		text := leerLinea()
		val, err := strconv.ParseFloat(text, 64)
		if err == nil {
			return val
		}
		fmt.Print("Ingrese un número válido: ")
	}
}

func main() {
	for {
		fmt.Println("\nSeleccione una opción:")
		fmt.Println("1 - Sueldos Despensa")
		fmt.Println("2 - Control de ventas en mercería")
		fmt.Println("3 - Compra de 3 productos")
		fmt.Println("4 - Sueldo de empleado por DNI y categoría")
		fmt.Println("0 - Salir")
		fmt.Print("Opción: ")

		opcion := leerInt()

		switch opcion {
		case 1:
			Ejercicio1()
		case 2:
			Ejercicio2()
		case 3:
			Ejercicio3()
		case 4:
			Ejercicio4()
		case 0:
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción inválida.")
		}
	}
}
