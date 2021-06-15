package main

import (
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
    "github.com/gotk3/gotk3/glib"
)

const (
    WIDTH int = 1366
    HEIGHT int = 736
    GRIDSIZE int = 5
)

func main() {
	gtk.Init(nil)

	// gui boilerplate
	win, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	da, _ := gtk.DrawingAreaNew()
    da.SetSizeRequest(WIDTH,HEIGHT)
	grid,_ := gtk.GridNew()
    win.Add(grid)
    butStart,_ := gtk.ButtonNewWithLabel("Start")
    butStop,_ := gtk.ButtonNewWithLabel("Stop")
    butRand,_ := gtk.ButtonNewWithLabel("Random")
    grid.Attach(da,0,0,3,1)
    grid.Attach(butRand,0,1,1,1)
    grid.Attach(butStart,1,1,1,1)
    grid.Attach(butStop,2,1,1,1)
	win.SetTitle("Game of Life")
	win.Connect("destroy", gtk.MainQuit)
	win.ShowAll()

    board := make([][]int, HEIGHT/GRIDSIZE)
    for i := range board {
        board[i] = make([]int, WIDTH/GRIDSIZE)
    }

	// Event handlers
	da.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
        // backgroud
        cr.SetSourceRGB(0.1, 0.1, 0.1)
		cr.Rectangle(0, 0, float64(WIDTH), float64(HEIGHT))
		cr.Fill()

		// draw gird
        cr.SetLineWidth(0.25);
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
        cr.SetSourceRGB(0.25, 1, 0.7)

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

    var id glib.SourceHandle

    butRand.Connect("clicked", func(b *gtk.Button) {
        randBoard(board)
        win.QueueDraw()
    }) 

    butStart.Connect("clicked", func(b *gtk.Button) {
        id = glib.TimeoutAdd(100,func() bool{
            gameOfLife(board)
            win.QueueDraw()
            return true
        })
        
    })
    butStop.Connect("clicked", func(b *gtk.Button) {
        glib.SourceRemove(id)
    })

    
	gtk.Main()
}