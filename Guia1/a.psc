Algoritmo a
    Definir eleccion Como Entero
    Definir per, area, radio, lado, apotema, largo, ancho Como Real
    Definir continuar Como Caracter
	
    Repetir
        Escribir "Por favor, indique: "
        Escribir "1: Para calcular el perimetro y area de un circulo dado su radio"
        Escribir "2: Para calcular el perimetro y area de un pentagono (le pedire que indique uno de sus lados)"
        Escribir "3: Para calcular el perimetro y area de un rectangulo (le pedire que indique la base y la altura)"
        Escribir " "
        Escribir "A continuacion indique que operacion desea realizar (1, 2 o 3):"
        Leer eleccion
		
        Si eleccion = 1 Entonces
            Escribir "Calculemos el perimetro y area de un circulo! :)"
            Escribir ""
            Escribir "Por favor indique el radio del circulo: "
            Leer radio
            Escribir "Calculemos el perimetro! :)"
            Escribir  " "
            per = (2 * PI * radio)
            Escribir  " "
            Escribir "El permitro de nuestro circulo es de: " per "cm ;)"
            Escribir  " "
            Escribir "Calculemos el area! :)"
            Escribir  " "
            area = PI * (radio^2)
            Escribir "El area de nuestro circulo es de: " area "cm² ;)"
            Escribir  " "
            Escribir "Respuesta el perimetro es de: " per, "cm y el area de este circulo es de: " area "cm² :)"
			
        SiNo
            Si eleccion = 2 Entonces
                Escribir "Calculemos el perimetro y area de un pentagono! :)"
                Escribir  " "
                Escribir "Por favor, indique uno de los lados: "
                Leer lado
                Escribir  " "
                Escribir "Calculemos el perimetro! :)"
                per = 5 * lado
                Escribir  " "
                Escribir "El permitro de nuestro pentagono es de: " per " ;)"
                Escribir  " "
                Escribir "Calculemos el area! :)"
                Escribir  " "
                Escribir "Para calcular el area necesitamos saber la apotema del pentagono que se realiza conociendo uno de sus lados (que ya lo tenemos yupii)"
                Escribir " "
                apotema = lado / (2 * tan(36 * PI / 180))
                Escribir "La apotema de nuestro pentagono es de: " apotema "cm ;)"
                Escribir " "
                area = (per * apotema) / 2
                Escribir "El area de nuestro pentagono es de: " area "cm² ;)"
                Escribir  " "
                Escribir "Respuesta el perimetro es de: " per, "cm y el area de este pentagono es de: " area "cm² :)"
				
            SiNo
                Si eleccion = 3 Entonces
                    Escribir "Calculemos el perimetro y area de un rectangulo! :)"
                    Escribir  " "
                    Escribir "Para calcular el perimetro de un rectangulo necesitaremos saber la longitud de sus 4 lados :("
                    Escribir  " "
                    Escribir "Para ahorrarnos tiempo y que german no escriba tanto, le pediremos que indique 2 lados: el del largo y el del ancho: "
                    Escribir  " "
                    Escribir "Por favor, indique el largo: "
                    Leer largo
                    Escribir  " "
                    Escribir "Por favor, indique el ancho: "
                    Leer ancho
                    Escribir  " "
                    Escribir "Ahora si! Calculemos el perimetro! :)"
                    per = 2 * (largo + ancho)
                    Escribir  " "
                    Escribir "El permitro de nuestro rectangulo es de: " per "cm ;)"
                    Escribir  " "
                    Escribir "Calculemos el area! :)"
                    Escribir  " "
                    area = largo * ancho
                    Escribir "El area de nuestro rectangulo es de: " area "cm² ;)"
                    Escribir  " "
                    Escribir "Respuesta el perimetro es de: " per, "cm y el area de este rectangulo es de: " area "cm² :)"
					
                SiNo
                    Escribir "Algo salio mal :( "
                Fin Si
            Fin Si
        Fin Si
		
        Escribir ""
        Escribir "¿Desea calcular otra cosa? (S/N): "
        Leer continuar
    Hasta Que continuar = "N" o continuar = "n"
	
    Escribir "Gracias por tanto, perdon por tan poco"
FinAlgoritmo
