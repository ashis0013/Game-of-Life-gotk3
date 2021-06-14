package main

import (
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
    "github.com/gotk3/gotk3/glib"
    "math/rand"
)

const (
    WIDTH int = 800
    HEIGHT int = 500
    GRIDSIZE int = 5
)

var board[HEIGHT/GRIDSIZE][WIDTH/GRIDSIZE] int;
var run bool;

func gameOfLife()  {
    if len(board) == 0 || len(board[0]) == 0 {
        return
    }
    m := len(board)
    n := len(board[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            lives := 0
            lives += isLive(i-1, j-1)
            lives += isLive(i-1, j)
            lives += isLive(i-1, j+1)
            lives += isLive(i, j-1)
            lives += isLive(i, j+1)
            lives += isLive(i+1, j-1)
            lives += isLive(i+1, j)
            lives += isLive(i+1, j+1)
            if board[i][j] == 0 {
                if lives == 3 {
                    board[i][j] = 2
                }
            } else {
                if lives < 2 || lives > 3 {
                    board[i][j] = 3
                }
            }
        }
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == 2 {
                board[i][j] = 1
            } else if board[i][j] == 3 {
                board[i][j] = 0
            }
        }
    }
}

func isLive(i int, j int) int {
    if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] == 0 || board[i][j] == 2 {
        return 0
    }
    return 1
}


func main() {
	gtk.Init(nil)

	// gui boilerplate
	win, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	da, _ := gtk.DrawingAreaNew()
    da.SetSizeRequest(WIDTH,HEIGHT)
	grid,_ := gtk.GridNew()
    win.Add(grid)
    button1,_ := gtk.ButtonNewWithLabel("B1")
    button2,_ := gtk.ButtonNewWithLabel("B2")
    grid.Attach(da,0,0,2,1)
    grid.Attach(button1,0,1,1,1)
    grid.Attach(button2,1,1,1,1)
	win.SetTitle("Game of Life")
	win.Connect("destroy", gtk.MainQuit)
	win.ShowAll()
    run = false

    for i := 0; i<HEIGHT/GRIDSIZE; i++ {
        for j := 0; j<WIDTH/GRIDSIZE; j++ {
            p := rand.Float64()
            if p < 0.2 {
                board[i][j] = 1
            }
        }
    }

	// Event handlers
	da.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
        // backgroud
        cr.SetSourceRGB(1, 1, 1)
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
        cr.SetSourceRGB(0, 0, 0)

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

    button1.Connect("clicked", func(b *gtk.Button) {
        run = true
        id = glib.TimeoutAdd(100,func() bool{
            gameOfLife()
            win.QueueDraw()
            return true
        })
        
    })
    button2.Connect("clicked", func(b *gtk.Button) {
        glib.SourceRemove(id)
    })

    
	gtk.Main()
}