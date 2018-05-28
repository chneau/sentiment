package main

import (
	"fmt"
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
	result := sentiment.Evaluate(os.Args[1:]...)
	fmt.Println(result)
}
