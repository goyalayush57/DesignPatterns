package main

type weatherData struct {
	observerList []observer
	name         string
}

func newWeatherData() *weatherData {
	return &weatherData{
		name: "Weather Stats",
	}
}

func (w weatherData) registerObserver(o observer) {
	w.observerList = append(w.observerList, o)
}

func (w weatherData) deregister(o observer) {
	w.observerList = removeFromslice(w.observerList, o)
}

func (w weatherData) notifyObserver() {
	for _, observer := range w.observerList {
		observer.update(w.name)
	}
}

func removeFromslice(observerList []observer, observerToRemove observer) []observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
