Algoritmo e
	
	Definir Var1, Var2, temp Como Entero
	Definir  res Como Entero
	
	Escribir "Ingrese la primer variable: "
	Leer Var1
	
	Escribir "Ingrese la segunda variable: "
	Leer Var2
	
	Escribir "Quiere cambiar las variables?: "
	Escribir "1 para si "
	Escribir "2 para no "
	Leer res
	
	Si res = 1 Entonces
		temp = Var1
		Var1 = Var2
		Var2 = temp
		
		Escribir "La primer variable es: " Var1
		Escribir "La segunda variable es: " Var2
	SiNo
		Si res = 2 Entonces
			Escribir "La primer variable es: " Var1
			Escribir "La segunda variable es: " Var2
		FinSi
		
	Fin Si
	

	
FinAlgoritmo
