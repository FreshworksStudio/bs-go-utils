package board

import (
	"fmt"
	"strings"

	"github.com/FreshworksStudio/bs-go-utils/api"
)

type Board struct {
	Width  int
	Height int
	Grid   [][]Entity
}

// Create a width x height board
func CreateBoard(width int, height int) *Board {
	b := new(Board)
	b.Width = width
	b.Height = height
	grid := make([][]Entity, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]Entity, width)
		for j := 0; j < width; j++ {
			grid[i][j] = Empty()
		}
	}

	b.Grid = grid
	return b
}

func (b Board) copy() *Board {
	copy := CreateBoard(b.Width, b.Height)
	for i := 0; i < copy.Height; i++ {
		for j := 0; j < copy.Width; j++ {
			copy.Grid[i][j] = GetEntity(b.Grid[i][j].EntityType)
		}
	}

	return copy
}

// Return if an x coordinate is in bounds
func (b Board) xInBounds(xpos int) bool {
	return (0 <= xpos && xpos < b.Width)
}

// Return if an x coordinate is in bounds
func (b Board) yInBounds(ypos int) bool {
	return (0 <= ypos && ypos < b.Height)
}

// Return what is in the grid at point p
func (b Board) getTile(p api.Point) Entity {
	if b.tileInBounds(p) {
		return b.Grid[p.Y][p.X]
	}
	return Invalid()
}

// Return if a point p is in bounds
func (b Board) tileInBounds(p api.Point) bool {
	return (b.xInBounds(p.X) && b.yInBounds(p.Y))
}

func (b Board) insert(p api.Point, e Entity) {
	if b.xInBounds(p.X) && b.yInBounds(p.Y) {
		b.Grid[p.Y][p.X] = e
	}
}

func (b Board) getValidTiles(p api.Point) []api.Point {
	validTiles := make([]api.Point, 0)
	potential := []api.Point{
		api.Point{p.X - 1, p.Y},
		api.Point{p.X + 1, p.Y},
		api.Point{p.X, p.Y - 1},
		api.Point{p.X, p.Y + 1},
	}
	for i, p := range potential {
		if b.tileInBounds(p) && b.getTile(p).EntityType != OBSTACLE {
			validTiles = append(validTiles, potential[i])
		}
	}
	return validTiles
}

func (b Board) show() {
	rowDivider := strings.Repeat(" ---", b.Width)
	println(rowDivider)
	for i := 0; i < b.Height; i++ {
		print("| ")
		for j := 0; j < b.Width; j++ {
			fmt.Printf("%s |", b.Grid[i][j].Display)
			print(" ")
		}
		println("\n" + rowDivider)
	}
}