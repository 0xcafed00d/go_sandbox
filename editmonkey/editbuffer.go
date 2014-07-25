package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"neomech/lib/neo/geom"
	"unicode/utf8"
)

type EditLine struct {
	line []byte
}

type EditBuffer struct {
	lines []EditLine
}

func (b *EditBuffer) GetLineCount() int {
	return len(b.lines)
}

func (b *EditBuffer) GetLine(idx int) EditLine {
	return b.lines[idx]
}

func (b *EditBuffer) AppendLine(l []byte) {
	b.lines = append(b.lines, EditLine{line: l})
}

func (b *EditBuffer) SetLine(idx int, l []byte) {
	b.lines[idx] = EditLine{line: l}
}

func (b *EditBuffer) InsertBlankLine(idx int) {
	var blank EditLine
	b.lines = append(b.lines, blank)

	from := b.lines[idx : len(b.lines)-1]
	to := b.lines[idx+1 : len(b.lines)]
	copy(to, from)
	b.lines[idx] = blank
}

type ViewSettings struct {
	TabSize int
}

type EditView struct {
	Settings      *ViewSettings
	Buffer        *EditBuffer
	Cursor        geom.Coord
	VisibleOrigin geom.Coord
	Selection     geom.Rectangle
	AreaSelect    bool
	EditableLine  []rune
}

func NextTabStop(pos, tabsize int) (nexttab, tabdist int) {
	tabdist = tabsize - (pos % tabsize)
	nexttab = pos + tabdist
	return
}

func EditLineToRuneSlice(l []byte, rs []rune, tabsize int) []rune {
	for len(l) > 0 {
		r, rlen := utf8.DecodeRune(l)

		if r < ' ' {
			if r == '\t' {
				_, tabdist := NextTabStop(len(rs), tabsize)
				for n := 0; n < tabdist; n++ {
					rs = append(rs, ' ')
				}
			} else {
				rs = append(rs, ' ')
			}

		} else {
			rs = append(rs, r)
		}

		l = l[rlen:]
	}

	return rs
}

func (e *EditView) CopyToEditLine(idx int) {
	l := e.Buffer.GetLine(idx).line
	e.EditableLine = e.EditableLine[0:0]
	e.EditableLine = EditLineToRuneSlice(l, e.EditableLine, e.Settings.TabSize)
}

func (e *EditView) InsertAtCursor(text []byte) {

}

func dumpBuffer(b *EditBuffer) {
	for n := 0; n < b.GetLineCount(); n++ {
		fmt.Printf("[%s]\n", b.GetLine(n).line)
	}
}

func dumpView(v *EditView) {
	for n := 0; n < v.Buffer.GetLineCount(); n++ {
		v.CopyToEditLine(n)
		fmt.Printf("[%s]\n", string(v.EditableLine))
	}
}

func main() {

	b := EditBuffer{}
	settings := ViewSettings{4}
	view := EditView{Settings: &settings, Buffer: &b}

	b.AppendLine([]byte("dd\t\tccc\t£\ttest1   "))
	b.AppendLine([]byte("012345678901234567890123456789"))
	b.AppendLine([]byte("+---+---+---+---+---+---+---+---+---"))
	b.AppendLine([]byte("test22"))
	b.AppendLine([]byte("test333"))

	dumpBuffer(&b)
	dumpView(&view)

	spew.Dump(&view)

	fmt.Printf("<%p>\n", view.EditableLine)
	dumpView(&view)
	fmt.Printf("<%p>\n", view.EditableLine)

	fmt.Printf("<%v, %v>\n", len(view.EditableLine), cap(view.EditableLine))

}
