package main

import (
	. "github.com/gizak/termui/v3"
	"image"
)

// Modified termui table, added custom style per cell
type ColorfulTable struct {
	Block
	Rows          [][]string
	ColumnWidths  []int
	TextStyle     Style
	RowSeparator  bool
	TextAlignment Alignment
	Styles        map[int]map[int]Style // [row][col]
	FillRow       bool

	// ColumnResizer is called on each Draw. Can be used for custom column sizing.
	ColumnResizer func()
}

func NewColorfulTable() *ColorfulTable {
	return &ColorfulTable{
		Block:         *NewBlock(),
		TextStyle:     Theme.Table.Text,
		RowSeparator:  true,
		Styles:        make(map[int]map[int]Style),
		ColumnResizer: func() {},
	}
}

func (self *ColorfulTable) StyleCell(row int, col int, style Style) {
	if self.Styles[row] == nil {
		self.Styles[row] = make(map[int]Style)
	}

	self.Styles[row][col] = style
}

func (self *ColorfulTable) Draw(buf *Buffer) {
	self.Block.Draw(buf)

	self.ColumnResizer()

	columnWidths := self.ColumnWidths
	if len(columnWidths) == 0 {
		columnCount := len(self.Rows[0])
		columnWidth := self.Inner.Dx() / columnCount
		for i := 0; i < columnCount; i++ {
			columnWidths = append(columnWidths, columnWidth)
		}
	}

	yCoordinate := self.Inner.Min.Y

	// draw rows
	for i := 0; i < len(self.Rows) && yCoordinate < self.Inner.Max.Y; i++ {
		row := self.Rows[i]
		colXCoordinate := self.Inner.Min.X

		if self.FillRow {
			blankCell := NewCell(' ', self.TextStyle)
			buf.Fill(blankCell, image.Rect(self.Inner.Min.X, yCoordinate, self.Inner.Max.X, yCoordinate+1))
		}

		// draw row cells
		for j := 0; j < len(row); j++ {
			cellStyle, exists := self.Styles[i][j]
			if !exists {
				cellStyle = self.TextStyle
			}

			col := ParseStyles(row[j], cellStyle)

			// draw row cell
			if len(col) > columnWidths[j] || self.TextAlignment == AlignLeft {
				for _, cx := range BuildCellWithXArray(col) {
					k, cell := cx.X, cx.Cell
					if k == columnWidths[j] || colXCoordinate+k == self.Inner.Max.X {
						cell.Rune = ELLIPSES
						buf.SetCell(cell, image.Pt(colXCoordinate+k-1, yCoordinate))
						break
					} else {
						buf.SetCell(cell, image.Pt(colXCoordinate+k, yCoordinate))
					}
				}
			} else if self.TextAlignment == AlignCenter {
				xCoordinateOffset := (columnWidths[j] - len(col)) / 2
				stringXCoordinate := xCoordinateOffset + colXCoordinate
				for _, cx := range BuildCellWithXArray(col) {
					k, cell := cx.X, cx.Cell
					buf.SetCell(cell, image.Pt(stringXCoordinate+k, yCoordinate))
				}
			} else if self.TextAlignment == AlignRight {
				stringXCoordinate := MinInt(colXCoordinate+columnWidths[j], self.Inner.Max.X) - len(col)
				for _, cx := range BuildCellWithXArray(col) {
					k, cell := cx.X, cx.Cell
					buf.SetCell(cell, image.Pt(stringXCoordinate+k, yCoordinate))
				}
			}
			colXCoordinate += columnWidths[j] + 1
		}

		// draw vertical separators
		separatorStyle := self.Block.BorderStyle

		separatorXCoordinate := self.Inner.Min.X
		verticalCell := NewCell(VERTICAL_LINE, separatorStyle)
		for i, width := range columnWidths {
			if self.FillRow && i < len(columnWidths)-1 {
				verticalCell.Style.Bg = self.TextStyle.Bg
			} else {
				verticalCell.Style.Bg = self.Block.BorderStyle.Bg
			}

			separatorXCoordinate += width
			buf.SetCell(verticalCell, image.Pt(separatorXCoordinate, yCoordinate))
			separatorXCoordinate++
		}

		yCoordinate++

		// draw horizontal separator
		horizontalCell := NewCell(HORIZONTAL_LINE, separatorStyle)
		if self.RowSeparator && yCoordinate < self.Inner.Max.Y && i != len(self.Rows)-1 {
			buf.Fill(horizontalCell, image.Rect(self.Inner.Min.X, yCoordinate, self.Inner.Max.X, yCoordinate+1))
			yCoordinate++
		}
	}
}
