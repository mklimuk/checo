package checo

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBoard_SetFENPieces(t *testing.T) {
	tests := []struct {
		FEN string
		expect Board
		expectError bool
	}{
		{"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR", [8][8]Piece{
			{WhiteRook, WhiteKnight, WhiteBishop, WhiteQueen, WhiteKing, WhiteBishop, WhiteKnight, WhiteRook},
			{WhitePawn,WhitePawn,WhitePawn,WhitePawn,WhitePawn,WhitePawn,WhitePawn,WhitePawn},
			{},
			{},
			{},
			{},
			{BlackPawn,BlackPawn,BlackPawn,BlackPawn,BlackPawn,BlackPawn,BlackPawn,BlackPawn},
			{BlackRook, BlackKnight, BlackBishop, BlackQueen, BlackKing, BlackBishop, BlackKnight, BlackRook},
		}, false},
	}
	for _, test := range tests {
		t.Run(test.FEN, func(t *testing.T) {
			b := Board{}
			err := b.SetFENPieces(test.FEN)
			b.Print(ChessSetUnicodeBW)
			if test.expectError {
				require.NotNil(t, err)
				return
			}
			assert.Equal(t, test.expect, b)
		})
	}
}

