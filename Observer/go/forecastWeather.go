package main

import "fmt"

type forecast struct {
	id string
}

func newForecast(id string) *forecast {
	return &forecast{
		id: id,
	}
}

func (f forecast) update(data string) {
	fmt.Printf("observer : %s Got Value : %s", f.id, data)
}

func (f forecast) getID() string {
	return f.id
}
