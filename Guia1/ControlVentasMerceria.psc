Algoritmo ControlVentasMerceria
	Definir cantidadPaquetes Como Entero
	Definir costoEnvio Como Real

	Escribir "MERCER�A MAYORISTA"
	Escribir "VENTA DE CANUTILLOS Y MOSTACILLAS"
	Escribir ""

	Escribir "Ingrese la cantidad de paquetes: "
	Leer cantidadPaquetes

	Segun cantidadPaquetes Hacer
		1,2,3,4:  // Casos menores a 5
			Escribir "NO se permiten compras inferiores a 5 productos."
		5,6,7,8,9,10,11,12,13,14,15:  // Casos entre 5 y 15
			costoEnvio <- 200
			Escribir "Costo de env�o: $", costoEnvio
			Escribir ""
			Escribir "Resumen de la compra:"
			Escribir "Cantidad de paquetes: ", cantidadPaquetes
			Escribir "Costo de env�o: $", costoEnvio
		De Otro Modo:
			Si cantidadPaquetes > 15 Entonces
				Escribir "�ENV�O GRATUITO!"
				costoEnvio <- 0
				Escribir ""
				Escribir "Resumen de la compra:"
				Escribir "Cantidad de paquetes: ", cantidadPaquetes
				Escribir "Costo de env�o: $", costoEnvio
			Sino
				Escribir "ERROR: Cantidad no v�lida."
			FinSi
	FinSegun

FinAlgoritmo
