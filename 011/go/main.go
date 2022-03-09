package main

import "flyweight/flyweight"

func main() {
	var cs flyweight.Chessboard
	cs.Location("white", 1, 2)
	cs.GetChess()
}
