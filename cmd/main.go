package main

import (
	"log"

	"remote_db/gui"
)

func main() {
	window := gui.CreateMainWindow()

	// Запускаем приложение
	if window == nil {
		log.Fatal("Error: window not created.")
	}

}
