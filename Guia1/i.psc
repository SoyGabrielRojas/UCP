Algoritmo i
    Definir Num1, Num2, Num3, Num4, promedio, total Como Real
    Definir notas Como Entero
    Definir nombre, situacion Como Caracter
	
    Escribir "Ingrese el nombre del alumno: "
    Leer nombre
	
    Escribir "Ingrese la primera nota: "
    Leer Num1
	
    Escribir "Ingrese la segunda nota: "
    Leer Num2
	
    Escribir "Ingrese la tercera nota: "
    Leer Num3
	
    Escribir "Ingrese la cuarta nota: "
    Leer Num4
	
    notas <- 4
    total <- Num1 + Num2 + Num3 + Num4
    promedio <- total / notas
	
    Si promedio > 10 Entonces
        situacion <- "El alumno es un genio, hay que llamar a la policia"
    Sino
        Si promedio >= 6 Entonces
            situacion <- "Aprobó"
        Sino
            situacion <- "Desaprobó"
        FinSi
    FinSi
	
    Escribir "El promedio de ", nombre, " es de ", promedio, " por lo tanto: ", situacion
FinAlgoritmo
