Algoritmo TP2_Act4_Escobar_Rojas
    Definir codigos, horasNormales, horasExtras, i, j, codigo, numEmpleados Como Entero
    Definir pagoNormal, pagoExtra, totalNormal, totalExtra, totalEmpleado, pagoTotalNormal, pagoTotalExtra Como Real
    Definir dias Como Cadena
    
    pagoNormal <- 2350
    pagoExtra <- 3500

    Dimension dias[7]
    dias[1] <- "Lunes"
    dias[2] <- "Martes"
    dias[3] <- "Miércoles"
    dias[4] <- "Jueves"
    dias[5] <- "Viernes"
    dias[6] <- "Sábado"
    dias[7] <- "Domingo"

    Escribir "Empresa La Huelga S.A."
    Escribir "Ingrese la cantidad de empleados (máximo 100): "
    Leer numEmpleados
    
    Mientras numEmpleados < 1 O numEmpleados > 100 Hacer
        Escribir "Error: El número de empleados debe estar entre 1 y 100"
        Escribir "Ingrese la cantidad de empleados: "
        Leer numEmpleados
    FinMientras
    
    Dimension codigos[numEmpleados]
    Dimension horasNormales[numEmpleados, 7]  
    Dimension horasExtras[numEmpleados, 7]    
    
    Escribir ""
    Escribir "Carga de datos:"
    Para i <- 1 Hasta numEmpleados Hacer
        Escribir ""
        Escribir "Empleado ", i, ":"
        
        Escribir "Ingresar código del empleado: "
        Leer codigos[i]
        
        Escribir "Ingresar horas trabajadas para el empleado ", codigos[i], ":"
        Para j <- 1 Hasta 7 Hacer
            Escribir dias[j], ":"
            Escribir "  Horas normales: "
            Leer horasNormales[i,j]
            Escribir "  Horas extras: "
            Leer horasExtras[i,j]
        FinPara
    FinPara
    
    Escribir ""
    Escribir "                    Informe de pagos - La huelga S.A."
    Escribir "Código | Horas Normales | $ Normales  | Horas Extras | $ Extras   | Total $"
    
    Para i <- 1 Hasta numEmpleados Hacer
        totalNormal <- 0
        totalExtra <- 0
        
        Para j <- 1 Hasta 7 Hacer
            totalNormal <- totalNormal + horasNormales[i,j]
            totalExtra <- totalExtra + horasExtras[i,j]
        FinPara

        pagoTotalNormal <- totalNormal * pagoNormal
        pagoTotalExtra <- totalExtra * pagoExtra
        totalEmpleado <- pagoTotalNormal + pagoTotalExtra

        Escribir codigos[i], "     | ", totalNormal, "           | $", pagoTotalNormal, "  | ", totalExtra, "          | $", pagoTotalExtra, "  | $", totalEmpleado
    FinPara

    
FinAlgoritmo