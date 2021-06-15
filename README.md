# Conway's Game of Life

The Game of Life, also known simply as Life, is a cellular automaton devised by the British mathematician John Horton Conway in 1970. It is a zero-player game, meaning that its evolution is determined by its initial state, requiring no further input. One interacts with the Game of Life by creating an initial configuration and observing how it evolves. It is Turing complete and can simulate a universal constructor or any other Turing machine.

## Rules

1. Any live cell with fewer than two live neighbours dies, as if by underpopulation.
2. Any live cell with two or three live neighbours lives on to the next generation.
3. Any live cell with more than three live neighbours dies, as if by overpopulation.
4. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

## Usage

This repository is built with gtk3 binding for go making it a cross-platform gui. To use it make sure you have go installed in your machine. Using it is simple by just building the project and running the excecutable

```bash
go build
./main
```

### Results

![ezgif com-gif-maker](https://user-images.githubusercontent.com/31564734/122069601-f0d29d80-ce12-11eb-9445-d054c790e73e.gif)
