Algoritmo TP2_Act1_Escobar_Rojas
	
    Definir matriz Como Entero
    Definir i, j Como Entero
    Definir max_val, min_val Como Entero
    Definir pos_max_fila, pos_max_col Como Entero
    Definir pos_min_fila, pos_min_col Como Entero
	
    Dimension matriz[5,10]
	
    Escribir "Ingresar los valores de la matriz (5 filas x 10 columnas):"
	
    Para i <- 1 Hasta 5 Con Paso 1 Hacer
        Para j <- 1 Hasta 10 Con Paso 1 Hacer
            Escribir "Ingresar el valor para la posición (", i, ",", j, "): "
            Leer matriz[i,j]
        FinPara
    FinPara
    Escribir ""
    Escribir "La matriz que fue ingresada:"
    Para i <- 1 Hasta 5 Con Paso 1 Hacer
        Para j <- 1 Hasta 10 Con Paso 1 Hacer
            Escribir Sin Saltar matriz[i,j], " "
        FinPara
    FinPara
    max_val <- matriz[1,1]
    min_val <- matriz[1,1]
    pos_max_fila <- 1
    pos_max_col <- 1
    pos_min_fila <- 1
    pos_min_col <- 1
    Para i <- 1 Hasta 5 Con Paso 1 Hacer
        Para j <- 1 Hasta 10 Con Paso 1 Hacer
            Si matriz[i,j] > max_val Entonces
                max_val <- matriz[i,j]
                pos_max_fila <- i
                pos_max_col <- j
            FinSi
			
            Si matriz[i,j] < min_val Entonces
                min_val <- matriz[i,j]
                pos_min_fila <- i
                pos_min_col <- j
            FinSi
        FinPara
    FinPara
    Escribir ""
    Escribir "El valor máximo es: ", max_val, " en la posición (", pos_max_fila, ",", pos_max_col, ")"
    Escribir "El valor mínimo es: ", min_val, " en la posición (", pos_min_fila, ",", pos_min_col, ")"
	
FinAlgoritmo
