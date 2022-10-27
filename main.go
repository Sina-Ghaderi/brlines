package main

import (
	"brlines/lines"
	"flag"
	"fmt"
	"log"
	"os"
)

const imagesize = 0x1f4

func main() {

	log.SetFlags(0)
	flag.Usage = flagUsage

	xa := flag.Int("xa", 0, "x coordinate of point a")
	ya := flag.Int("ya", 0, "y coordinate of point a")
	xb := flag.Int("xb", 0, "x coordinate of point b")
	yb := flag.Int("yb", 0, "y coordinate of point b")
	ot := flag.String("pt", "br_resualt.png", "output png file to save results")

	flag.Parse()

	if *xa < 0 || *xa > 19 || *xb < 0 || *xb > 19 || *ya < 0 || *ya > 19 || *yb < 0 || *yb > 19 {
		log.Fatalf("fatal error: %v", "bad input, please use coordinates 0 upto 19")
	}

	img, err := lines.NewBreLine(imagesize, *ot)
	if err != nil {
		log.Fatalf("fatal error: %v", err)
	}

	img.DrawMesh()
	img.BresenhamLine(*xa, *ya, *xb, *yb)

	if err := img.WriteToFile(); err != nil {
		log.Fatalf("fatal error: %v", err)
	}
}

func flagUsage() {
	fmt.Printf(`usage of bresenham line simulator:
%v options...

options:
  --pt  <file>      named png file to save bresenham output
                    default path for this file is br_resualt.png

  --xa  <uint>      x coordinate of starting point
  --xb  <uint>      x coordinate of ending point

  --ya  <uint>      y coordinate of starting point
  --yb  <uint>      y coordinate of ending point

example:
   %v --pt mytest.png --xa 8 --ya 3 --xb 2 --yb 10


Copyright (c) 2022 snix.ir, All rights reserved.
Developed BY <Sina Ghaderi> sina@snix.ir
This work is licensed under the terms of the MIT licence
Github: github.com/sina-ghaderi and Source: git.snix.ir
`, os.Args[0], os.Args[0])
}
