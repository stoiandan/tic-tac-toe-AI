package model

//Move consists of a symbol applied to a GameBoard, at a certain row-column
type Move struct {
	RowIndex    int
	ColumnIndex int
	Symbol      Symbol
}
