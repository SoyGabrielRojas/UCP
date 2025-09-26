Algoritmo f
	
	Definir peso, imc, estatura Como Real
	
	Escribir "Ingrese su peso en kg: "
	Leer peso
	
	Escribir "Ingrese su estatura en m: "
	Leer estatura
	
	imc = peso / (estatura * 2)
	
	Escribir "Su imc es: " imc
	
	Escribir "Tabla de valores dentro del rango: "
	Escribir ""
	
	Escribir "Menos de 18.5: Bajo peso."
	Escribir "18.5 - 24.9: Peso normal."
	Escribir "25.0 - 29.9: Sobrepeso."
	Escribir "30.0 - 34.9: Obesidad grado I."
	Escribir "35.0 - 39.9: Obesidad grado II."
	Escribir "40 o más: Obesidad grado III."
	
	Escribir ""

	
	Si imc < 18 Entonces
        Escribir "Usted esta bajo de peso"
    Sino
        Si imc > 18 y imc < 24 Entonces
            Escribir "Usted tiene un peso normal"
        Sino
			Si imc > 20 y imc < 29 Entonces
				Escribir "Usted tiene Sobrepeso"
			SiNo
				Si imc > 30 y imc < 34 Entonces
					Escribir "Usted tiene Obesidad grado I."
				SiNo
					Si imc > 35 y imc < 39 Entonces
						Escribir "Usted tiene Obesidad grado II."
					FinSi
				FinSi
			FinSi
            Escribir "Usted tiene Obesidad grado III."
        FinSi
    FinSi
FinAlgoritmo
