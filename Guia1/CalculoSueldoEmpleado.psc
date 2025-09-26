Algoritmo CalculoSueldoEmpleado
    Definir dni, categoria Como Entero
    Definir sueldoBruto, descJubilacion, descObraSocial, descClub, sueldoNeto Como Real
    Definir nombreCategoria Como Cadena
    
    Escribir "SISTEMA DE CÁLCULO DE SUELDOS"
    Escribir ""
    
    // Ingreso de datos del empleado
    Escribir "Ingrese el DNI del empleado: "
    Leer dni
    Escribir ""
    
    Escribir "Categorías disponibles:"
    Escribir "0 - Maestranza"
    Escribir "1 - Administración"
    Escribir "2 - Gerencia"
    Escribir ""
    Escribir "Ingrese la categoría del empleado (0-2): "
    Leer categoria
    
    // Validar categoría
    Si categoria < 0 O categoria > 2 Entonces
        Escribir "ERROR: Categoría no válida. Debe ser 0, 1 o 2."
	FinSi

	// Asignar valores según categoría
	Segun categoria Hacer
		0:
			nombreCategoria <- "Maestranza"
			sueldoBruto <- 23600
			descJubilacion <- sueldoBruto * 0.11
			descObraSocial <- sueldoBruto * 0.03
			descClub <- 0
			
		1:
			nombreCategoria <- "Administración"
			sueldoBruto <- 35800
			descJubilacion <- sueldoBruto * 0.11
			descObraSocial <- sueldoBruto * 0.05
			descClub <- 0
			
		2:
			nombreCategoria <- "Gerencia"
			sueldoBruto <- 60420
			descJubilacion <- sueldoBruto * 0.11
			descObraSocial <- sueldoBruto * 0.05
			descClub <- sueldoBruto * 0.04
	FinSegun

	// Calcular sueldo neto
	sueldoNeto <- sueldoBruto - descJubilacion - descObraSocial - descClub

	// Mostrar resultados
	Escribir ""
	Escribir "LIQUIDACIÓN DE SUELDO"
	Escribir "DNI: ", dni
	Escribir "Categoría: ", nombreCategoria
	Escribir ""
	Escribir "Sueldo Bruto: $", sueldoBruto
	Escribir "Descuento Jubilación (11%): $", descJubilacion
	Escribir "Descuento Obra Social: $", descObraSocial

	Si categoria = 2 Entonces
		Escribir "Descuento Club (4%): $", descClub
	FinSi

	Escribir "SUELDO NETO: $", sueldoNeto
	Escribir ""

FinAlgoritmo