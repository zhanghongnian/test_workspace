package main

import (
	"fmt"
	"image/jpeg"
	"os"
	"sync"
	"time"

	"github.com/nfnt/resize"
)

// PrintCost 打印函数花费时间
func PrintCost(i int, f func() error) {
	var (
		start = time.Now()
		err   error
	)
	defer func() {
		fmt.Printf("cost miehaha, %v, %v \n", i, time.Now().Sub(start))
	}()
	if err = f(); err != nil {
		fmt.Printf("err: %+v \n", err)
		return
	}
}

func test() (err error) {
	// open "test.jpg"
	file, err := os.Open("images/test_1.jpg")
	if err != nil {
		return
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		return
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(1000, 0, img, resize.Lanczos3)

	out, err := os.Create("test_resized.jpg")
	if err != nil {
		return
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
	return
}

// func main() {
// 	// var wg sync.WaitGroup
// 	c := 50
// 	for i := 0; i < c; i++ {
// 		// wg.Add(1)
// 		// go func(i int) {
// 		// 	defer wg.Done()
// 		PrintCost(i, test)
// 		// }(i)
// 	}
// 	// wg.Wait()
// }

func main() {
	var wg sync.WaitGroup
	c := 15
	for i := 0; i < c; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			PrintCost(i, test)
		}(i)
	}
	wg.Wait()
}
