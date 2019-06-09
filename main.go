package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/chromedp"
)

func main() {
	// create chrome instance
	userAgent := "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36"
	allocContext, _ := chromedp.NewExecAllocator(
		context.Background(),
		chromedp.UserAgent(userAgent),
	)

	ctx, cancel := chromedp.NewContext(
		allocContext,
		chromedp.WithLogf(log.Printf),
	)

	defer cancel()

	loginURL := "https://accounts.google.com/signin/v2/identifier?service=mail&flowName=GlifWebSignIn&flowEntry=ServiceLogin"
	registerElement := "#ow243"
	err := chromedp.Run(ctx, LoginPage(loginURL, registerElement))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("logged")

}

// https://github.com/TestingPens/SwarmIt/blob/master/register.go
// LoginPage ..
func LoginPage(urlstr, elem string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitReady(elem),
		chromedp.Submit(elem),
	}
}
