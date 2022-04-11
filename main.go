package main

import (
	"fmt"
	"time"

	progressbar "github.com/schollz/progressbar/v3"
)

func main() {
	fmt.Println("開始執行進度條測試。")

	progressBar := progressbar.NewOptions(1000,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription("[cyan]測試進度[reset] 請稍等..."),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))
		
	for i := 0; i < 1000; i++ {
		progressBar.Add(1)
		time.Sleep(5 * time.Millisecond)
	}

	progressBar.Clear()
	progressBar.Close()

	fmt.Println("進度條結束。")
}
