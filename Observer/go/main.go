package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	weather := newWeatherData()
	//observer1 := newForecast("Observer 1")
	// observer2 := newforecast("Observer 2")
	//weather.registerObserver(observer1)
	// weather.registerObserver(observer2)
	weather.notifyObserver()
}
