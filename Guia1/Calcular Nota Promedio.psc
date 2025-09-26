Algoritmo CalcularNotaPromedio
	Definir nota1, nota2, nota3, nota4 Como Real
    Definir promedio, parteEntera, parteDecimal, notaFinal Como Real
	
    Escribir "Ingrese la primera nota (1 a 10):"
    Leer nota1
    Escribir "Ingrese la segunda nota (1 a 10):"
    Leer nota2
    Escribir "Ingrese la tercera nota (1 a 10):"
    Leer nota3
    Escribir "Ingrese la cuarta nota (1 a 10):"
    Leer nota4
	
    promedio <- (nota1 + nota2 + nota3 + nota4) / 4
	
    parteEntera <- Trunc(promedio)
    parteDecimal <- promedio - parteEntera
	
    Si parteDecimal < 0.5 Entonces
        notaFinal <- parteEntera + 0.5
    Sino
        notaFinal <- parteEntera + 1
    FinSi
	
    Escribir "La nota promedio final es: ", notaFinal
FinAlgoritmo
