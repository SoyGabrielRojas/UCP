Algoritmo edad
	Definir diaNac, mesNac, ANONAC Como Entero
	Definir diaActual, mesActual, ANOACTUAL Como Entero
	Definir EDADANOS, edadMeses Como Entero
	Escribir 'Ingrese el día de nacimiento: '
	Leer diaNac
	Escribir 'Ingrese el mes de nacimiento: '
	Leer mesNac
	Escribir 'Ingrese el año de nacimiento: '
	Leer ANONAC
	Escribir 'Ingrese el día actual: '
	Leer diaActual
	Escribir 'Ingrese el mes actual: '
	Leer mesActual
	Escribir 'Ingrese el año actual: '
	Leer ANOACTUAL
	EDADANOS <- ANOACTUAL-ANONAC
	edadMeses <- mesActual-mesNac
	Si edadMeses<0 Entonces
		EDADANOS <- EDADANOS-1
		edadMeses <- edadMeses+12
	FinSi
	Escribir 'Usted tiene ', edadAños, ' años y ', edadMeses, ' meses.'
FinAlgoritmo
