package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"time"
)

type UUID [16]byte

// NewUUIDv7 genera un nuevo UUID de versiÃ³n 7 basado en un timestamp proporcionado
func NewUUIDv7(timestamp *uint64) UUID {
	var uuid UUID

	// Establecer timestamp actual en milisegundos
	*timestamp = uint64(time.Now().UnixNano() / int64(time.Millisecond))

	// Rellenar los primeros 6 bytes con el timestamp
	uuid[0] = byte(*timestamp >> 40)
	uuid[1] = byte(*timestamp >> 32)
	uuid[2] = byte(*timestamp >> 24)
	uuid[3] = byte(*timestamp >> 16)
	uuid[4] = byte(*timestamp >> 8)
	uuid[5] = byte(*timestamp)

	// Generar 10 bytes aleatorios
	randomBytes := make([]byte, 10)
	if _, err := rand.Read(randomBytes); err != nil {
		panic(err)
	}
	copy(uuid[6:], randomBytes)

	// Establecer la versiÃ³n (7) y los bits de variante (2 MSB como 01)
	uuid[6] = (uuid[6] & 0x0f) | (7 << 4)
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	return uuid
}

func main() {
	const max_gen = 1_000_000

	now := time.Now()

	// Crear el archivo para guardar los UUIDs
	f, err := os.Create("NewUUIDv7.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	timestamp := uint64(1724411918)
	i := 1
	for i <= max_gen {
		uuidV7 := NewUUIDv7(&timestamp)

		// Formatear UUID en string en formato GUID
		
		guid := fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
			uuidV7[0:4],
			uuidV7[4:6],
			uuidV7[6:8],
			uuidV7[8:10],
			uuidV7[10:16],
		)

		// Escribir el GUID en el archivo
		_, err := fmt.Fprintf(f, "NewUUIDv7 --> %d --> %s\n", i, guid)
		if err != nil {
			fmt.Println("Error al escribir en el archivo:", err)
			return
		}
		i++
	}

	fmt.Println("Tiempo transcurrido:", time.Since(now))
	fmt.Printf( "%d --> %x\n", (7 << 4), (7 << 4) )
}
