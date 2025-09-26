Algoritmo MayorYMenorDeTresNumeros
	Definir A, B, C Como Entero
	
    Escribir "Ingrese el primer número (A):"
    Leer A
    Escribir "Ingrese el segundo número (B):"
    Leer B
    Escribir "Ingrese el tercer número (C):"
    Leer C
	
    Si (A = B) Y (B = C) Entonces
        Escribir "Los tres números son iguales."
    Sino
        Si (A > B) Y (A > C) Entonces
            mayor <- A
        Sino
            Si (B > A) Y (B > C) Entonces
                mayor <- B
            Sino
                mayor <- C
            FinSi
        FinSi
		
        Si (A < B) Y (A < C) Entonces
            menor <- A
        Sino
            Si (B < A) Y (B < C) Entonces
                menor <- B
            Sino
                menor <- C
            FinSi
        FinSi
		
        Escribir "El mayor es: ", mayor
        Escribir "El menor es: ", menor
    FinSi
FinAlgoritmo
