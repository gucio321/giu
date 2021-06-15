package giu

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Shortcut struct {
	Callback func()
	Key      Key
	Modifier Modifier
}

var shortcuts map[KeyCombo]*CallbackGroup

func init() {
	shortcuts = make(map[KeyCombo]*CallbackGroup)
}

type KeyCombo struct {
	Key      glfw.Key
	Modifier glfw.ModifierKey
}

type CallbackGroup struct {
	Global func()
	Window func()
}

func createKeyCombo(key Key, modifier Modifier) KeyCombo {
	return KeyCombo{
		Key:      glfw.Key(key),
		Modifier: glfw.ModifierKey(modifier),
	}
}

func RegisterShortcut(cb func(), key Key, mod Modifier, isGlobal bool) {
	combo := createKeyCombo(key, mod)

	shortcut, isRegistered := shortcuts[combo]

	if !isRegistered {
		shortcut = &CallbackGroup{}
	}

	if isGlobal {
		shortcut.Global = cb
	} else {
		shortcut.Window = cb
	}

	shortcuts[combo] = shortcut
}

func UnregisterWindowShortcuts() {
	for _, cbs := range shortcuts {
		cbs.Window = nil
	}
}

func InputHandler(key glfw.Key, mod glfw.ModifierKey, action glfw.Action) {
	for combo, cb := range shortcuts {
		if key != combo.Key || mod != combo.Modifier || action != glfw.Press {
			continue
		}

		if cb.Window != nil {
			cb.Window()
		} else if cb.Global != nil {
			cb.Global()
		}

		return
	}
}
