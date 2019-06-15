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
		chromedp.WithDebugf(log.Printf),
	)

	defer cancel()

	// loginURL := "https://accounts.google.com/signin/v2/identifier?service=mail&flowName=GlifWebSignIn&flowEntry=ServiceLogin"
	// loginURL := "https://google.com"
	loginURL := "https://github.com"
	registerElement := "#ow244.U26fgb.c7fp5b.FS4hgd.nDKKZc.NpwL8d.t29vte"
	personalResgisterElement := `#initialView div.xkfVF.nnGvjf div.JPdR6b.qjTEB div.XvhY1d div.JAPqpe.K0NPx span.z80M1.G3hhxb`
	err := chromedp.Run(ctx, LoginPage(loginURL, registerElement, personalResgisterElement))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("logged")

}

// https://github.com/TestingPens/SwarmIt/blob/master/register.go
// LoginPage ..
func LoginPage(urlstr, elem1, elem2 string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		// chromedp.Sleep(2 * time.Second),
		// chromedp.WaitVisible(elem1, chromedp.ByID),
		// // chromedp.Sleep(time.Duration(8239-rand.Intn(100)) * time.Millisecond),
		// chromedp.Click(elem1, chromedp.NodeVisible),
		// chromedp.WaitVisible(elem2, chromedp.ByID),
		// // chromedp.Sleep(time.Duration(8239-rand.Intn(100)) * time.Millisecond),
		// chromedp.Click(elem2, chromedp.NodeVisible),
	}
}
