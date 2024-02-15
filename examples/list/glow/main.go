package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
	humanize "github.com/dustin/go-humanize"
)

type Document struct {
	Name string
	Date time.Time
}

func (d Document) String() string {
	return d.Name + "\n" + lipgloss.NewStyle().Faint(true).Render(humanize.Time(d.Date))
}

var docs = []Document{
	{"README.md", time.Now().Add(-time.Minute * 2)},
	{"Example.md", time.Now().Add(-time.Hour)},
	{"secrets.md", time.Now().Add(-time.Hour * 24 * 7)},
}

const selectedIndex = 1

func main() {
	baseStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("250")).MarginBottom(1).MarginLeft(1)
	highlightStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#EE6FF8")).MarginBottom(1).MarginLeft(1)

	l := list.New().Enumerator(func(_ *list.List, i int) string {
		if i == selectedIndex {
			return "│\n│"
		}
		return ""
	}).
		ItemStyleFunc(func(_ *list.List, i int) lipgloss.Style {
			if selectedIndex == i {
				return highlightStyle
			}
			return baseStyle
		}).
		EnumeratorStyleFunc(func(_ *list.List, i int) lipgloss.Style {
			if selectedIndex == i {
				return highlightStyle
			}
			return baseStyle
		})

	for _, d := range docs {
		l.Item(d.String())
	}

	fmt.Println()
	fmt.Println(l)
}