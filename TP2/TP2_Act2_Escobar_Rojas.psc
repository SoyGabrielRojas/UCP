Algoritmo TP2_Act2_Escobar_Rojas
		Definir matriz, traspuesta Como Entero
		Definir i, j Como Entero
		Dimension matriz[5,10]
		Dimension traspuesta[10,5]
		
		Escribir "Ingrese los valores de la matriz (5 filas x 10 columnas):"
		Para i <- 1 Hasta 5 Con Paso 1 Hacer
			Para j <- 1 Hasta 10 Con Paso 1 Hacer
				Escribir "Ingrese el valor para la posición (", i, ",", j, "): "
				Leer matriz[i,j]
			FinPara
		FinPara
		Para i <- 1 Hasta 5 Con Paso 1 Hacer
			Para j <- 1 Hasta 10 Con Paso 1 Hacer
				traspuesta[j,i] <- matriz[i,j]
			FinPara
		FinPara
		Escribir "Matriz original:"
		Para i <- 1 Hasta 5 Con Paso 1 Hacer
			Para j <- 1 Hasta 10 Con Paso 1 Hacer
				Escribir Sin Saltar matriz[i,j], " "
			FinPara
			Escribir ""
		FinPara
		Escribir "Matriz traspuesta:"
		Para i <- 1 Hasta 10 Con Paso 1 Hacer
			Para j <- 1 Hasta 5 Con Paso 1 Hacer
				Escribir Sin Saltar traspuesta[i,j], " "
			FinPara
			Escribir ""
		FinPara
		
FinAlgoritmo
