package checo

import (
	"fmt"
	"github.com/fatih/color"
	"strconv"
	"strings"
	"unicode"
)

type Piece uint8

const (
	Empty Piece = iota
	Pawn
	Knight
	Bishop
	Rook
	Queen
	King
)

var BlackRook = Piece(Black) | Rook
var BlackKnight = Piece(Black) | Knight
var BlackBishop = Piece(Black) | Bishop
var BlackQueen = Piece(Black) | Queen
var BlackKing = Piece(Black) | King
var BlackPawn = Piece(Black) | Pawn

var WhiteRook = Piece(White) | Rook
var WhiteKnight = Piece(White) | Knight
var WhiteBishop = Piece(White) | Bishop
var WhiteQueen = Piece(White) | Queen
var WhiteKing = Piece(White) | King
var WhitePawn = Piece(White) | Pawn

var fenNotationToPiece = map[rune]Piece {
	'k': King,
	'q': Queen,
	'r': Rook,
	'b': Bishop,
	'n': Knight,
	'p': Pawn,
}

var fenPieceToNotation = map[Piece]rune {
	King:'k',
	Queen :'q',
	 Rook:'r',
	Bishop:'b',
	Knight:'n',
	Pawn: 'p',
	Empty: ' ',
}

type ChessSet map[Piece]rune

var ChessSetFEN = map[Piece]rune {
	WhiteKing: 'K',
	WhiteQueen: 'Q',
	WhiteRook: 'R',
	WhiteBishop: 'B',
	WhiteKnight: 'N',
	WhitePawn: 'P',
	BlackKing: 'k',
	BlackQueen: 'q',
	BlackRook: 'r',
	BlackBishop: 'b',
	BlackKnight: 'n',
	BlackPawn: 'p',
	Empty: ' ',
}

var ChessSetUnicode = map[Piece]rune {
	WhiteKing: '♔',
	WhiteQueen: '♕',
	WhiteRook: '♖',
	WhiteBishop: '♗',
	WhiteKnight: '♘',
	WhitePawn: '♙',
	BlackKing: '♚',
	BlackQueen: '♛',
	BlackRook: '♜',
	BlackBishop: '♝',
	BlackKnight: '♞',
	BlackPawn: '♟',
	Empty: ' ',
}

var colorizeWhite = color.New(color.FgHiWhite).SprintFunc()
var colorizeBlack = color.New(color.FgBlack).SprintFunc()

var ChessSetUnicodeBW = map[Piece]rune {
	WhiteKing: '♚',
	WhiteQueen: '♛',
	WhiteRook: '♜',
	WhiteBishop: '♝',
	WhiteKnight: '♞',
	WhitePawn: '♟',
	BlackKing: '♚',
	BlackQueen: '♛',
	BlackRook: '♜',
	BlackBishop: '♝',
	BlackKnight: '♞',
	BlackPawn: '♟',
	Empty: ' ',
}

type Color uint8

func (c Color) SetPieceFromNotation(r rune) Piece {
	return Piece(c) + fenNotationToPiece[r]
}

const (
	White Color = 8
	Black Color = 16
)

const (
	pieceMask = 0b111
	colorMask = 0b11000
)

func (p Piece) SetColor(color Color) Piece {
	piece := p & pieceMask
	return piece | Piece(color)
}

func (p Piece) Kind() Piece {
	return p & pieceMask
}

func (p Piece) Color() Color {
	return Color(p & colorMask)
}

func (p Piece) String() string {
	piece := fenPieceToNotation[p & pieceMask]
	if Color(p & colorMask) == White {
		piece = unicode.ToUpper(piece)
	}
	return string(piece)
}

type Board [8][8]Piece

func (b *Board) SetFENPieces(fenPieces string) error {
	ranks := strings.Split(strings.TrimSpace(fenPieces), "/")
	if len(ranks) != 8 {
		return fmt.Errorf("expected 8 ranks in FEN notation; found %d", len(ranks))
	}
	for i := 7; i >= 0; i-- {
		j := 0
		for _, r := range ranks[i] {
			switch r {
			case 'r','n','b','q','k','p':
				b[i][j] = White.SetPieceFromNotation(r)
				j++
			case 'R','N','B','Q','K','P':
				b[i][j] = Black.SetPieceFromNotation(unicode.ToLower(r))
				j++
			default:
				// otherwise we must have an integer between 1 and 8
				skip, err := strconv.Atoi(string(r))
				if err != nil {
					return fmt.Errorf("invalid skip files number format: %w", err)
				}
				if skip < 1 || skip > 8 {
					return fmt.Errorf("can skip between 1 and 8 files; got %d", skip)
				}
				j += skip
			}
		}
		if j < 7 {
			return fmt.Errorf("incomplete rank %d; expected 8 files, got %d", i+1, j+1)
		}
	}
	return nil
}

func (b *Board) Print(set ChessSet) {
	fmt.Print("  ╔════════════════════════╗\n")
	for i := 7; i >= 0; i-- {
		fmt.Print(strconv.Itoa(i+1))
		fmt.Print(" ║")
		for j := 0; j < 8; j++ {
			if (i + j) % 2 == 0 {
				// dark
				color.Set(color.BgBlue)
			} else {
				// light
				color.Set(color.BgHiYellow)
			}
			piece := b[i][j]
			if piece.Color() == White {
				color.Set(color.FgWhite)
			} else {
				color.Set(color.FgBlack)
			}
			fmt.Print(" ")
			fmt.Print(string(set[piece]))
			fmt.Print(" ")
			color.Unset()
		}
		fmt.Print("║\n")
	}
	fmt.Print("  ╚════════════════════════╝\n")
	fmt.Println("    A  B  C  D  E  F  G  H ")
}

func NewBoard() Board {
	b := Board{}
	_ = b.SetFENPieces("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")
	return b
}