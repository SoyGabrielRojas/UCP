def ejercicio1():
    filas, columnas = 5, 10
    matriz = []
    for i in range(filas):
        fila = []
        for j in range(columnas):
            valor = int(input(f"Ingrese el valor para [{i}][{j}]: "))
            fila.append(valor)
        matriz.append(fila)

    max_val = matriz[0][0]
    min_val = matriz[0][0]
    max_pos = (0, 0)
    min_pos = (0, 0)
    for i in range(filas):
        for j in range(columnas):
            if matriz[i][j] > max_val:
                max_val = matriz[i][j]
                max_pos = (i, j)
            if matriz[i][j] < min_val:
                min_val = matriz[i][j]
                min_pos = (i, j)
    print(f"Máximo: {max_val} en posición {max_pos}")
    print(f"Mínimo: {min_val} en posición {min_pos}")
    print("Matriz ingresada:")
    for fila in matriz:
        print(" ".join(str(x) for x in fila))
    return matriz

def ejercicio2(matriz):
    filas = len(matriz)
    columnas = len(matriz[0])
    traspuesta = [[matriz[i][j] for i in range(filas)] for j in range(columnas)]
    print("Matriz traspuesta:")
    for fila in traspuesta:
        print(" ".join(str(x) for x in fila))

def ejercicio3():
    p = int(input("Ingrese el orden de la matriz cuadrada: "))
    matriz = []
    for i in range(p):
        fila = []
        for j in range(p):
            valor = int(input(f"Ingrese el valor para [{i}][{j}]: "))
            fila.append(valor)
        matriz.append(fila)

    print("Diagonal principal:", end=" ")
    for i in range(p):
        print(matriz[i][i], end=" ")
    print()

    print("Diagonal secundaria:", end=" ")
    for i in range(p):
        print(matriz[i][p-1-i], end=" ")
    print()

    print("Triangular superior:", end=" ")
    for i in range(p):
        for j in range(i+1, p):
            print(matriz[i][j], end=" ")
    print()

    print("Triangular inferior:", end=" ")
    for i in range(1, p):
        for j in range(i):
            print(matriz[i][j], end=" ")
    print()

def ejercicio4():
    empleados = 3  # Puedes cambiar la cantidad de empleados
    codigos = []
    horas_normales = []
    horas_extras = []

    for i in range(empleados):
        codigo = int(input(f"Ingrese el código del empleado {i+1} (1-100): "))
        horas_n = int(input("Ingrese horas normales trabajadas: "))
        horas_e = int(input("Ingrese horas extras trabajadas: "))
        codigos.append(codigo)
        horas_normales.append(horas_n)
        horas_extras.append(horas_e)

    print("\n--- Informe de empleados ---")
    for i in range(empleados):
        pago_normal = horas_normales[i] * 2350
        pago_extra = horas_extras[i] * 3500
        total = pago_normal + pago_extra
        print(f"Código: {codigos[i]} | Horas normales: {horas_normales[i]} | $ normales: {pago_normal:.2f} | Horas extras: {horas_extras[i]} | $ extras: {pago_extra:.2f} | Total: {total:.2f}")

def main():
    matriz_ej1 = None
    while True:
        print("\nSeleccione una opción:")
        print("1 - Matriz 5x10: máximo, mínimo y posiciones")
        print("2 - Matriz traspuesta (requiere completar ejercicio 1)")
        print("3 - Diagonales y triangulares de matriz cuadrada")
        print("4 - Informe de empleados La Huelga S.A.")
        print("0 - Salir")
        opcion = input("Opción: ")

        if opcion == "1":
            matriz_ej1 = ejercicio1()
        elif opcion == "2":
            if matriz_ej1 is None:
                print("Primero debe completar el ejercicio 1 para poder ejecutar el ejercicio 2.")
            else:
                ejercicio2(matriz_ej1)
        elif opcion == "3":
            ejercicio3()
        elif opcion == "4":
            ejercicio4()
        elif opcion == "0":
            print("Saliendo...")
            break
        else:
            print("Opción inválida.")

if __name__ == "__main__":
    main()