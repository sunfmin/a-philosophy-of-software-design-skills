package store

// History records undoable actions for the editor.
type History struct {
	actions []action
	pos     int
}

type action struct {
	kind      string // "insert", "delete", "selection"
	text      string
	offset    int
	selStart  int
	selEnd    int
}

func (h *History) Add(a action) { h.actions = append(h.actions[:h.pos], a); h.pos++ }

// Undo reverts the last action.
func (h *History) Undo(doc *Document, ui *EditorUI) {
	if h.pos == 0 {
		return
	}
	h.pos--
	a := h.actions[h.pos]
	switch a.kind {
	case "insert":
		doc.text = doc.text[:a.offset] + doc.text[a.offset+len(a.text):]
	case "delete":
		doc.text = doc.text[:a.offset] + a.text + doc.text[a.offset:]
	case "selection":
		ui.selStart, ui.selEnd = a.selStart, a.selEnd
	}
}

type Document struct{ text string }
type EditorUI struct{ selStart, selEnd int }
