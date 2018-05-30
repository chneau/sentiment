package main

import (
	"bufio"
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
	toAnalyse := []string{}
	if len(os.Args) == 1 {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter text then press enter")
		text, _ := reader.ReadString('\n')
		toAnalyse = []string{text}
	} else {
		toAnalyse = os.Args[1:]
	}
	result := sentiment.Evaluate(toAnalyse...)
	fmt.Println("Score:", result)
}
