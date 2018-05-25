package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/chneau/sentiment/pkg/sentiment"
)

func init() {
	gracefulExit()
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func ce(err error, msg string) {
	if err != nil {
		log.Panicln(msg, err)
	}
}

func gracefulExit() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		os.Exit(0)
	}()
}

func printMapWords() {
	for word, result := range sentiment.Words {
		log.Println(result, word)
	}
}

func main() {
	// TODO: command line to check a phrase
	// TODO: feed more data into the map (like yougster language "lol" and so on)
	word := "hello i'm super happy haha"
	result := sentiment.Evaluate(word)
	log.Println("word", word)
	log.Println("result", result)
}
