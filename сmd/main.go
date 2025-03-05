package main

import (
	app "https://github.com/DinnerDer/iCalc/internal_application"
)

func main() {
	a := app.New()
	a.RunServer() // Запускаем приложение
}
