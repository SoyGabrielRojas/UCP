package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Ejercicio1() {
	for i := 1; i <= 35; i++ {
		fmt.Println(i)
	}
}

func Ejercicio2() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nIngrese el límite numérico: ")
	limiteStr, _ := reader.ReadString('\n')
	limite, _ := strconv.Atoi(strings.TrimSpace(limiteStr))
	for i := 1; i <= limite; i++ {
		fmt.Println(i)
	}
}

func Ejercicio3() {
	for i := 200; i <= 250; i += 2 {
		fmt.Println(i)
	}
}

func Ejercicio4() {
	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}
}

func Ejercicio5() {
	reader := bufio.NewReader(os.Stdin)
	for cliente := 0; cliente < 5; cliente++ {
		fmt.Print("\nIngrese el DNI del cliente: ")
		dni, _ := reader.ReadString('\n')
		dni = strings.TrimSpace(dni)
		fmt.Print("Ingrese el tipo de servicio (1-30M, 2-50M, 3-100M): ")
		tipoStr, _ := reader.ReadString('\n')
		tipo, _ := strconv.Atoi(strings.TrimSpace(tipoStr))
		var monto float64
		switch tipo {
		case 1:
			monto = 750
		case 2:
			monto = 1100
		case 3:
			monto = 1500 * 0.95
		default:
			fmt.Println("Tipo de servicio inválido.")
			continue
		}
		fmt.Printf("DNI: %s | Servicio: %d | Monto a pagar: $%.2f\n", dni, tipo, monto)
	}
}

func Ejercicio6() {
	reader := bufio.NewReader(os.Stdin)
	var mejorTiempo float64
	var mejorVehiculo int
	for i := 0; i < 12; i++ {
		fmt.Print("\nIngrese el número de vehículo: ")
		vehiculoStr, _ := reader.ReadString('\n')
		vehiculo, _ := strconv.Atoi(strings.TrimSpace(vehiculoStr))
		fmt.Print("Ingrese el tiempo (segundos): ")
		tiempoStr, _ := reader.ReadString('\n')
		tiempo, _ := strconv.ParseFloat(strings.TrimSpace(tiempoStr), 64)
		if i == 0 || tiempo < mejorTiempo {
			mejorTiempo = tiempo
			mejorVehiculo = vehiculo
		}
	}
	fmt.Printf("\nEl vehículo con mejor tiempo es el %d con %.2f segundos.\n", mejorVehiculo, mejorTiempo)
}

func Ejercicio7() {
	reader := bufio.NewReader(os.Stdin)
	contTenis := 0
	contFutbol := 0
	deportes := map[int][2]int{1: {0, 0}, 2: {0, 0}, 3: {0, 0}, 4: {0, 0}, 5: {0, 0}}
	fmt.Print("\nIngrese la cantidad de socios: ")
	sociosStr, _ := reader.ReadString('\n')
	socios, _ := strconv.Atoi(strings.TrimSpace(sociosStr))
	for i := 0; i < socios; i++ {
		fmt.Print("\nNúmero de socio: ")
		_, _ = reader.ReadString('\n') // num_socio, not used
		fmt.Print("Edad: ")
		edadStr, _ := reader.ReadString('\n')
		edad, _ := strconv.Atoi(strings.TrimSpace(edadStr))
		fmt.Print("Deporte (1-tenis,2-rugby,3-voley,4-hockey,5-futbol): ")
		deporteStr, _ := reader.ReadString('\n')
		deporte, _ := strconv.Atoi(strings.TrimSpace(deporteStr))
		if deporte == 1 {
			contTenis++
		} else if deporte == 5 {
			contFutbol++
		}
		if datos, ok := deportes[deporte]; ok {
			deportes[deporte] = [2]int{datos[0] + edad, datos[1] + 1}
		}
	}
	fmt.Printf("\nSocios que practican tenis: %d\n", contTenis)
	fmt.Printf("Socios que practican fútbol: %d\n", contFutbol)
	for dep, datos := range deportes {
		sumaEdad, cant := datos[0], datos[1]
		if cant > 0 {
			promedio := float64(sumaEdad) / float64(cant)
			fmt.Printf("Deporte %d: promedio de edad = %.2f\n", dep, promedio)
		} else {
			fmt.Printf("Deporte %d: sin jugadores\n", dep)
		}
	}
}

func Ejercicio8() {
	reader := bufio.NewReader(os.Stdin)
	totalPersonas := 0
	cantVarones := 0
	cantMujeres := 0
	varones1665 := 0
	mayorEdad := -1
	var personaMayor struct {
		dni  int
		edad int
		sexo string
	}
	for {
		fmt.Print("\nIngrese el número de documento (0 para terminar): ")
		dniStr, _ := reader.ReadString('\n')
		dni, _ := strconv.Atoi(strings.TrimSpace(dniStr))
		if dni == 0 {
			break
		}
		fmt.Print("Ingrese la edad: ")
		edadStr, _ := reader.ReadString('\n')
		edad, _ := strconv.Atoi(strings.TrimSpace(edadStr))
		fmt.Print("Ingrese el sexo (M/F): ")
		sexoStr, _ := reader.ReadString('\n')
		sexo := strings.ToUpper(strings.TrimSpace(sexoStr))
		totalPersonas++
		if sexo == "M" {
			cantVarones++
			if edad >= 16 && edad <= 65 {
				varones1665++
			}
		} else if sexo == "F" {
			cantMujeres++
		}
		if edad > mayorEdad {
			mayorEdad = edad
			personaMayor.dni = dni
			personaMayor.edad = edad
			personaMayor.sexo = sexo
		}
	}
	var porcentajeVarones float64
	if cantVarones > 0 {
		porcentajeVarones = float64(varones1665) / float64(cantVarones) * 100
	}
	fmt.Printf("\nTotal de personas censadas: %d\n", totalPersonas)
	fmt.Printf("Cantidad de varones: %d\n", cantVarones)
	fmt.Printf("Cantidad de mujeres: %d\n", cantMujeres)
	fmt.Printf("Porcentaje de varones (16-65 años): %.2f%%\n", porcentajeVarones)
	if mayorEdad != -1 {
		fmt.Printf("Persona de mayor edad → DNI: %d, Edad: %d, Sexo: %s\n", personaMayor.dni, personaMayor.edad, personaMayor.sexo)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nSeleccione el ejercicio a ejecutar:")
		fmt.Println("1. Mostrar números del 1 al 35")
		fmt.Println("2. Mostrar números hasta un límite dado")
		fmt.Println("3. Números del 200 al 250 de 2 en 2")
		fmt.Println("4. Cuenta regresiva de año nuevo (10 a 1)")
		fmt.Println("5. Calcular monto de internet para 5 clientes")
		fmt.Println("6. Mejor tiempo en carrera (12 competidores)")
		fmt.Println("7. Club de socios")
		fmt.Println("8. Censo provincial")
		fmt.Println("0. Salir")
		fmt.Print("Opción: ")
		opcionStr, _ := reader.ReadString('\n')
		opcionStr = strings.TrimSpace(opcionStr)
		switch opcionStr {
		case "1":
			Ejercicio1()
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
		case "8":
			Ejercicio8()
		case "0":
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción inválida.")
		}
	}
}
