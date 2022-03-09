package flyweight

import (
	"errors"
	"fmt"
)

const (
	chessBlack = "black"
	chessWhite = "white"
)

var chesses chessFactory

/**********************flyweight factory***********************/
type chessFactory struct {
	chessKinds map[string]*chess
}

func (c *chessFactory) Init() {
	c.chessKinds = make(map[string]*chess)

	var black chess
	black.color = chessBlack
	black.size = 10
	black.weight = 10
	c.chessKinds[chessBlack] = &black
	var white chess
	white.color = chessWhite
	white.size = 10
	white.weight = 10
	c.chessKinds[chessWhite] = &white
}

func (c *chessFactory) GetChess(color string) (*chess, error) {
	if cs, exist := c.chessKinds[color]; exist {
		return cs, nil
	}
	return nil, errors.New("not support color")
}

// singleton
func getChess(color string) (*chess, error) {
	if nil == chesses.chessKinds {
		chesses.Init()
	}
	return chesses.GetChess(color)
}

/*****************flyweight interface***********************/
// interface or pointer which not need the whole memory

/*****************concrete flyweight***********************/
type chess struct {
	size   int
	weight int
	color  string
}

/**********************Context***********************/
type Chessboard struct {
	cs *chess
	x  int
	y  int
}

func (c *Chessboard) Location(color string, x int, y int) {
	var err error
	if c.cs, err = getChess(color); nil != err {
		fmt.Printf("err is %s\n", err)
	}
	c.x = x
	c.y = y
}
func (c Chessboard) GetChess() {
	fmt.Printf("---------------flyweight-------------------\n")
	fmt.Printf("weight is %+v\n", c.cs.weight)
	fmt.Printf("size is %+v\n", c.cs.size)
	fmt.Printf("color is %+v\n", c.cs.color)
	fmt.Printf("---------------unique-------------------\n")
	fmt.Printf("x is %+v\n", c.x)
	fmt.Printf("y is %+v\n", c.y)
}
