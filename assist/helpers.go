package assist

import (
	"github.com/JamesHovious/w32"
)

func moveMouse(x int32, y int32) {
	w32.SendInput([]w32.INPUT{
		{
			Type: w32.INPUT_MOUSE,
			Mi: w32.MOUSEINPUT{
				Dx:      x,
				Dy:      y,
				DwFlags: w32.MOUSEEVENTF_MOVE,
			},
		},
	})
}

func click() {
	w32.SendInput([]w32.INPUT{
		{
			Type: w32.INPUT_MOUSE,
			Mi: w32.MOUSEINPUT{
				DwFlags: w32.MOUSEEVENTF_LEFTDOWN,
			},
		},
	})
}

func sendKeyPress(vkey uint16) {
	scanCode := uint16(w32.MapVirtualKeyEx(uint(vkey), w32.MAPVK_VK_TO_VSC, 0))
	w32.SendInput([]w32.INPUT{
		{
			Type: w32.INPUT_KEYBOARD,
			Ki: w32.KEYBDINPUT{
				WScan:   scanCode,
				DwFlags: KEYEVENTF_SCANCODE | KEYEVENTF_KEYDOWN,
			},
		},
		{
			Type: w32.INPUT_KEYBOARD,
			Ki: w32.KEYBDINPUT{
				WScan:   scanCode,
				DwFlags: KEYEVENTF_SCANCODE | KEYEVENTF_KEYUP,
			},
		},
	})
}

func isPressed(vKey int) bool {
	return w32.GetAsyncKeyState(vKey)&0x8000 > 0
}

func isToggled(vKey int) bool {
	return w32.GetKeyState(vKey) > 0
}

func GetCurrentAsyncKeyState() []int {
	pressedKeys := []int{}
	for i := 1; i < 256; i++ {
		keyState := w32.GetAsyncKeyState(i)
		if keyState == 0x8001 {
			pressedKeys = append(pressedKeys, i)
		}
	}
	return pressedKeys
}
