package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"random-quote-picture/config"
	"random-quote-picture/service"
)

func init() {
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	config.Init(workPath)
	log.SetPrefix("[LOG] ")
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)
}

func main() {
	category := flag.String("category", "", "input the category, quote or picture")

	key := flag.String("key", "", "When category is equal to 'quote', you could use this option to specified, the key is an integer and the maximum length is 6")

	grayscale := flag.String("grayscale", "", "When category is equal to 'picture', you could use this option to specified, when the grayscale value is 1, you can get a grayscale picture")

	flag.Parse()

	// check parameter
	check(*category)

	// log.Println("category:", *category)
	if *category == "quote" {
		quote, err := service.GetOneQuote(*key)
		if err != nil {
			fmt.Printf("Get Quote Error: %+v\n", err.Error())
			return
		}
		fmt.Println("======Quote======")
		fmt.Printf("Text:%s\nAuthor:%s\nLink:%s\n", quote.QuoteText, quote.QuoteAuthor, quote.QuoteLink)
		fmt.Println("=================")
	} else if *category == "picture" {
		// How to return an image on cli?
		isGrayscale := *grayscale
		fmt.Println("isGrayscale:", isGrayscale)

	}

}

func check(categoryName string) {
	if categoryName == "" {
		fmt.Println("Please provide the 'category' parameter, which can be either 'quote' or 'picture'.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if categoryName != "quote" && categoryName != "picture" {
		fmt.Println("Wrong category value. Please provide the 'category' parameter, which can be either 'quote' or 'picture'.")
		os.Exit(1)
	}
}
