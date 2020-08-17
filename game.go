package checo

import "errors"

var ErrIllegalMove = errors.New("illegal move")

type View interface {
	Draw(Board) error
}

type Game struct {
	board Board
	frontend View
}


