package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time" // Paquete esencial para medir tiempo
)

// CalculadorPi representa un "Objeto con Salud"
type CalculadorPi struct {
	Rounds int `tag:"iteraciones_configuradas"`
}

// ValidarSalud asegura que el número de rondas sea procesable
func (c *CalculadorPi) ValidarSalud() error {
	if c.Rounds <= 0 {
		return fmt.Errorf("salud insuficiente: las rondas deben ser mayores a 0")
	}
	return nil
}

func main() {
	// 1. Lectura y preparación
	file, err := os.ReadFile("rounds.txt")
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}

	val, _ := strconv.Atoi(strings.TrimSpace(string(file)))
	calc := CalculadorPi{Rounds: val}

	// Verificamos la "salud" del objeto antes de iniciar la carga pesada
	if err := calc.ValidarSalud(); err != nil {
		fmt.Println(err)
		return
	}

	// --- INICIO DE MEDICIÓN ---
	start := time.Now()

	x := 1.0
	pi := 1.0
	stop := float64(calc.Rounds + 2)

	for i := 2.0; i <= stop; i++ {
		x = -x
		pi += x / (2.0*i - 1.0)
	}

	pi *= 4.0

	// --- FIN DE MEDICIÓN ---
	duration := time.Since(start)

	// 3. Resultados
	fmt.Printf("Valor de PI: %.15f\n", pi)
	fmt.Printf("Tiempo transcurrido: %v\n", duration)
}
