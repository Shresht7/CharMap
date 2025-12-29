package char_list

import (
	"github.com/charmbracelet/bubbles/list"
)

const (
	HPadding = 4
	VPadding = 2
)

func New() list.Model {
	delegate := newDelegate()
	l := list.New([]list.Item{}, delegate, 0, 0)
	l.Title = ""
	l.SetFilteringEnabled(false)
	l.SetShowStatusBar(false)
	l.SetShowHelp(true)
	return l
}
