package main

type subject interface {
	registerObserver(o observer)
	deregister(o observer)
	notifyObserver()
}
