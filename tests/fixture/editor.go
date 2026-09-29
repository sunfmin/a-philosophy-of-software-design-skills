package store

// debugMode is read by several packages at startup.
var debugMode bool

type BaseView struct {
	lines  []string
	scroll int
}

func (b *BaseView) Render() string { return b.lines[b.scroll] }

// ScrollView overrides Render and mutates the parent's fields directly.
type ScrollView struct{ BaseView }

func (s *ScrollView) Render() string {
	if s.scroll >= len(s.lines) {
		s.scroll = len(s.lines) - 1
	}
	return s.BaseView.Render()
}

type Text struct{ BaseView }

// Backspace deletes the char before the cursor.
func (t *Text) Backspace(line, col int) {
	l := t.lines[line]
	t.lines[line] = l[:col-1] + l[col:]
}

// Lines returns the internal slice.
func (t *Text) Lines() []string { return t.lines }

func copyBlock(block int, fileBlocks map[int]int) int {
	block = fileBlocks[block] // now a disk block
	return block
}
