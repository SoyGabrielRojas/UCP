Algoritmo CalculoSueldosDespensa
    Definir categoria Como Entero
    Definir sueldoBruto, jubilacion, obraSocial, sueldoNeto, bono Como Real
    
    Escribir "Categorías disponibles:"
    Escribir "1 - Repositor"
    Escribir "2 - Cajero"
    Escribir "3 - Supervisor"
    Escribir ""
    
    Escribir "Ingrese la categoría del empleado (1-3): "
    Leer categoria
    
    Segun categoria Hacer
        1:
            sueldoBruto <- 32335
            Escribir "Empleador: Repositor"
            Escribir "Sueldo bruto: $", sueldoBruto
            
            // Calcular bono para repositor
            bono <- sueldoBruto * 0.08
            Escribir "Bono por compras (8%): $", bono
            
        2:
            sueldoBruto <- 38630.89
            Escribir "Empleado: Cajero"
            Escribir "Sueldo bruto: $", sueldoBruto
            
        3:
            sueldoBruto <- 45560.20
            Escribir "Empleado: Supervisor"
            Escribir "Sueldo bruto: $", sueldoBruto
            
        De Otro Modo:
            Escribir "ERROR: Categoría no válida. Debe ser 1, 2 o 3."
	FinSegun

	// Calcular deducciones
	jubilacion <- sueldoBruto * 0.11
	obraSocial <- sueldoBruto * 0.03
	sueldoNeto <- sueldoBruto - jubilacion - obraSocial

	Escribir ""
	Escribir "Descuento por jubilación (11%): $", jubilacion
	Escribir "Descuento por obra social (3%): $", obraSocial
	Escribir "Total deducciones: $", jubilacion + obraSocial
	Escribir ""
	Escribir "SUELDO NETO: $", sueldoNeto

FinAlgoritmo