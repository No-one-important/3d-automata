package main

import "log"

func checkPanic(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

func main() {
	var engine Engine
	engine.Init()
	engine.Run()
	engine.Stop()
}
