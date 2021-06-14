package main

import (
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
    "math/rand"
)

const (
    WIDTH int = 400
    HEIGHT int = 200
    GRIDSIZE int = 10
)

var board[HEIGHT/GRIDSIZE][WIDTH/GRIDSIZE] int;


func main() {
	gtk.Init(nil)

	// gui boilerplate
	win, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	da, _ := gtk.DrawingAreaNew()
    da.SetSizeRequest(WIDTH,HEIGHT)
	win.Add(da)
	win.SetTitle("Game of Life")
	win.Connect("destroy", gtk.MainQuit)
	win.ShowAll()

    for i := 0; i<HEIGHT/GRIDSIZE; i++ {
        for j := 0; j<WIDTH/GRIDSIZE; j++ {
            p := rand.Float64()
            if p < 0.2 {
                board[i][j] = 1
            }
        }
    }

    board[2][7] = 1;

	// Event handlers
	da.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
        // backgroud
        cr.SetSourceRGB(1, 1, 1)
		cr.Rectangle(0, 0, float64(WIDTH), float64(HEIGHT))
		cr.Fill()

		// draw gird
        cr.SetLineWidth(1.0);
        cr.SetSourceRGB(0, 0, 0)
        for i := 1; i<HEIGHT/GRIDSIZE; i++ {
            cr.MoveTo(0, float64(i*GRIDSIZE));
            cr.LineTo(float64(WIDTH), float64(i*GRIDSIZE));
            cr.Stroke();
        }
        for i := 1; i<WIDTH/GRIDSIZE; i++ {
            cr.MoveTo(float64(i*GRIDSIZE), 0);
            cr.LineTo(float64(i*GRIDSIZE), float64(HEIGHT));
            cr.Stroke();
        }

        // map array
        for i := 0; i<HEIGHT/GRIDSIZE; i++ {
            for j := 0; j<WIDTH/GRIDSIZE; j++ {
                if(board[i][j] == 1) {
                    cr.Rectangle(float64(j*GRIDSIZE), float64(i*GRIDSIZE), float64(GRIDSIZE), float64(GRIDSIZE))
		            cr.Fill()
                }
            }
        }
	})

	gtk.Main()
}