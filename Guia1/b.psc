Algoritmo b
    Definir edad Como Entero 
    Definir sexo, decir Como Caracter
	
    Escribir "Vamos a ver si usted puede votar! :)"
    Escribir " "
	
    Repetir
        Escribir "Por favor, ingrese su sexo (mujer u hombre): "
        Leer sexo
        Si No (sexo = "Mujer" o sexo = "mujer" o sexo = "Hombre" o sexo = "hombre") Entonces
            Escribir "Revise si escribio bien su sexo (Hombre/hombre, Mujer/mujer)"
            Escribir " "
        Fin Si
    Hasta Que sexo = "Mujer" o sexo = "mujer" o sexo = "Hombre" o sexo = "hombre"
	
    Escribir " "
    Escribir "Ahora, por favor, ingrese su edad (sin que sea, por ejemplo, como 18 y 3 meses, que sea solo 18 por favor): "
    Leer edad
    Escribir " "
	
    Si (sexo = "Mujer" o sexo = "mujer") y edad >= 16 Entonces
        decir = "Se�orita" //siempre se le dice se�orita a una dama
        Escribir decir " usted es mujer y puede votar ;)"
    SiNo
        Si (sexo = "Mujer" o sexo = "mujer") y edad < 16 Entonces
            decir = "Se�orita" //siempre se le dice se�orita a una dama
            Escribir decir " usted es mujer y no puede votar :("
        SiNo
            Si (sexo = "Hombre" o sexo = "hombre") y edad >= 16 Entonces
                decir = "Se�or"
                Escribir decir " usted es hombre y puede votar ;)"
            SiNo
                decir = "Se�orito"
                Escribir decir " usted es hombre y no puede votar :("
            Fin Si
        Fin Si
    Fin Si
FinAlgoritmo
