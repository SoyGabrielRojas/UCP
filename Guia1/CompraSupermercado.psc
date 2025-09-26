Algoritmo CompraSupermercado
    Definir precio1, precio2, precio3 Como Real
    Definir cantidad1, cantidad2, cantidad3 Como Real
    Definir monto1, monto2, monto3, total, totalConDescuento Como Real
    
    Escribir "SISTEMA DE COMPRAS"
    Escribir ""
    
    // Ingreso de datos del primer producto
    Escribir "PRODUCTO 1:"
    Escribir "Ingrese el precio por Kg: $"
    Leer precio1
    Escribir "Ingrese la cantidad en Kg: "
    Leer cantidad1
    Escribir ""
    
    // Ingreso de datos del segundo producto
    Escribir "PRODUCTO 2:"
    Escribir "Ingrese el precio por Kg: $"
    Leer precio2
    Escribir "Ingrese la cantidad en Kg: "
    Leer cantidad2
    Escribir ""
    
    // Ingreso de datos del tercer producto
    Escribir "PRODUCTO 3:"
    Escribir "Ingrese el precio por Kg: $"
    Leer precio3
    Escribir "Ingrese la cantidad en Kg: "
    Leer cantidad3
    Escribir ""
    
    // Cálculo de montos por producto
    monto1 <- precio1 * cantidad1
    monto2 <- precio2 * cantidad2
    monto3 <- precio3 * cantidad3
    
    // Cálculo del total
    total <- monto1 + monto2 + monto3
    
    // Mostrar resultados
    Escribir "=== DETALLE DE LA COMPRA ==="
    Escribir "Producto 1: ", cantidad1, " Kg x $", precio1, " = $", monto1
    Escribir "Producto 2: ", cantidad2, " Kg x $", precio2, " = $", monto2
    Escribir "Producto 3: ", cantidad3, " Kg x $", precio3, " = $", monto3
    Escribir "TOTAL DE LA COMPRA: $", total
    
    // Aplicar descuento si corresponde
    Si total > 100 Entonces
        totalConDescuento <- total * 0.90  // 10% de descuento
        Escribir ""
        Escribir "¡DESCUENTO APLICADO! (10% por compra mayor a $100)"
        Escribir "Total con descuento: $", totalConDescuento
        Escribir "Ahorro: $", total - totalConDescuento
    Sino
        Escribir ""
        Escribir "Para obtener un 10% de descuento, compre más de $100"
    FinSi
    
FinAlgoritmo
