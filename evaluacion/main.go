package main

import "fmt"

// Variables globales
var productosVendidos []string
var subtotales []float64

// Función para registrar una venta
func RegistrarVenta(nombre string, precio float64, cantidad int) {
	subtotal := precio * float64(cantidad)

	productosVendidos = append(productosVendidos, nombre)
	subtotales = append(subtotales, subtotal)

	fmt.Println("Venta registrada correctamente.")
}

// Función para mostrar estadísticas
func MostrarEstadisticas() {
	if len(subtotales) == 0 {
		fmt.Println("No existen ventas registradas.")
		return
	}

	var total float64

	for _, subtotal := range subtotales {
		total += subtotal
	}

	fmt.Printf("Total recaudado: $%.2f\n", total)
}

func main() {
	// Productos y precios definidos en el sistema
	productos := []string{"Arroz", "Leche", "Pan"}
	precios := []float64{1.25, 0.95, 0.50}

	var opcion int

	for {
		fmt.Println("\n===== MENÚ =====")
		fmt.Println("1. Registrar nueva venta")
		fmt.Println("2. Mostrar estadísticas")
		fmt.Println("3. Salir")
		fmt.Print("Seleccione una opción: ")
		fmt.Scan(&opcion)

		switch opcion {

		case 1:
			fmt.Println("\nProductos disponibles:")
			for i := 0; i < len(productos); i++ {
				fmt.Printf("%d. %s - $%.2f\n", i+1, productos[i], precios[i])
			}

			var seleccion int
			fmt.Print("Seleccione un producto: ")
			fmt.Scan(&seleccion)

			if seleccion < 1 || seleccion > len(productos) {
				fmt.Println("Producto no válido.")
				continue
			}

			var cantidad int
			fmt.Print("Ingrese la cantidad vendida: ")
			fmt.Scan(&cantidad)

			RegistrarVenta(
				productos[seleccion-1],
				precios[seleccion-1],
				cantidad,
			)

		case 2:
			MostrarEstadisticas()

		case 3:
			fmt.Println("Saliendo del programa...")
			return

		default:
			fmt.Println("Opción inválida.")
		}
	}
}