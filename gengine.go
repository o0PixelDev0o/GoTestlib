package gengine

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func Greet(){
	fmt.Println("GENGINE")
}

func NewWindow(Title string){
	cfg := pixelgl.WindowConfig{
		//Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	
	win.Clear(colornames.Skyblue)

	for !win.Closed() {
		win.Update()
	}

	pixelgl.Run(run)
}

