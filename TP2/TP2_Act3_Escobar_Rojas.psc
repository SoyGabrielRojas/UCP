Algoritmo TP2_Act3_Escobar_Rojas
	Definir C Como Entero
    Definir P, i, j Como Entero
	
    Escribir "Ingresar el orden de la matriz cuadrada P: "
    Leer P
    Dimension C[P,P]
    Escribir "Ingresar los valores de la matriz (", P, "x", P, "):"
    Para i <- 1 Hasta P Con Paso 1 Hacer
        Para j <- 1 Hasta P Con Paso 1 Hacer
            Escribir "Elemento (", i, ",", j, "): "
            Leer C[i,j]
        FinPara
    FinPara
    Escribir ""
    Escribir "Matriz que ingresamos:"
    Para i <- 1 Hasta P Con Paso 1 Hacer
        Para j <- 1 Hasta P Con Paso 1 Hacer
            Escribir Sin Saltar C[i,j], " "
        FinPara
        Escribir ""
    FinPara
    Escribir ""
    Escribir "Diagonal principal:"
    Para i <- 1 Hasta P Con Paso 1 Hacer
        Escribir C[i,i], " "
    FinPara
    Escribir ""
    Escribir "Diagonal secundaria:"
    Para i <- 1 Hasta P Con Paso 1 Hacer
        Escribir C[i, P - i + 1], " "
    FinPara
    Escribir ""
    Escribir "Triangular superior:"
    Para i <- 1 Hasta P Con Paso 1 Hacer
        Para j <- i+1 Hasta P Con Paso 1 Hacer
            Escribir C[i,j], " "
        FinPara
    FinPara
    Escribir ""
    Escribir "Triangular inferior:"
    Para i <- 2 Hasta P Con Paso 1 Hacer
        Para j <- 1 Hasta i-1 Con Paso 1 Hacer
            Escribir C[i,j], " "
        FinPara
    FinPara
FinAlgoritmo
