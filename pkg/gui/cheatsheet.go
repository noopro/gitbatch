package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func (gui *Gui) openCheatSheetView(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := g.Size()
	v, err := g.SetView(cheatSheetViewFeature.Name, maxX/2-25, maxY/2-10, maxX/2+25, maxY/2+10)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = cheatSheetViewFeature.Title
		fmt.Fprintln(v, " ")
		for _, k := range gui.KeyBindings {
			if k.View == mainViewFeature.Name || k.View == "" {
				binding := " " + k.Display + ": " + k.Description
				fmt.Fprintln(v, binding)
			}
		}
	}
	gui.updateKeyBindingsView(g, cheatSheetViewFeature.Name)
	if _, err := g.SetCurrentView(cheatSheetViewFeature.Name); err != nil {
		return err
	}
	return nil
}

func (gui *Gui) closeCheatSheetView(g *gocui.Gui, v *gocui.View) error {
	if err := g.DeleteView(v.Name()); err != nil {
		return nil
	}
	if _, err := g.SetCurrentView(mainViewFeature.Name); err != nil {
		return err
	}
	gui.updateKeyBindingsView(g, mainViewFeature.Name)
	return nil
}
