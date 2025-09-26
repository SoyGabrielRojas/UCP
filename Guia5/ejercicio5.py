tabla = [
    [8, 8, 7, 0],
    [7, 9, 10, 0],
    [10, 9, 5, 0],
    [4, 9, 8.5, 0]
]

for i in range(4):
    tabla[i][3] = (tabla[i][0] + tabla[i][1] + tabla[i][2]) / 3

print("Ejercicio 5:")
print("         Nota 1  Nota 2  Nota 3  Prom.")
for i in range(4):
    print(f"Alumno {i+1}: {tabla[i][0]:5.2f}   {tabla[i][1]:5.2f}   {tabla[i][2]:5.2f}   {tabla[i][3]:5.2f}")