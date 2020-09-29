package assist

import (
	"time"

	"github.com/JamesHovious/w32"
)

const (
	// Keyboard Event
	KEYEVENTF_SCANCODE = 0x0008
	KEYEVENTF_KEYDOWN  = 0x0
	KEYEVENTF_KEYUP    = 0x0002
)

type Keybinds struct {
	ADS               int
	Ping              uint16
	RecoilToggle      int
	RecoilSenseToggle int
	RapidFireToggle   int
}

type Delays struct {
	Standard       time.Duration
	AutoPingDelay  time.Duration
	RapidFireDelay time.Duration
}

type RecoilSense struct {
	Low  time.Duration
	High time.Duration
}

type Assist struct {
	Keybinds    Keybinds
	Delays      Delays
	RecoilSense RecoilSense
}

func NewAssist(keyBinds Keybinds, delays Delays, recoilSense RecoilSense) *Assist {
	return &Assist{
		Keybinds:    keyBinds,
		Delays:      delays,
		RecoilSense: recoilSense,
	}
}

func (a *Assist) stdDelay() {
	time.Sleep(a.Delays.Standard)
}

func (a *Assist) getSense() time.Duration {
	if isToggled(a.Keybinds.RecoilSenseToggle) {
		return a.RecoilSense.High
	}

	return a.RecoilSense.Low
}

func (a *Assist) isFiring() bool {
	return isPressed(w32.VK_LBUTTON)
}

func (a *Assist) isADS() bool {
	return isPressed(a.Keybinds.ADS)
}

func (a *Assist) isADSFiring() bool {
	return a.isFiring() && a.isADS()
}

func (a *Assist) isRapidFire() bool {
	return a.isADS() && isPressed(a.Keybinds.RapidFireToggle)
}

func (a *Assist) handleRecoil() {
	for {
		for a.isADSFiring() || a.isRapidFire() {
			moveMouse(0, 1)
			time.Sleep(a.getSense())
		}
		a.stdDelay()
	}
}

func (a *Assist) handleAutoTag() {
	for {
		for a.isADSFiring() {
			time.Sleep(a.Delays.AutoPingDelay)
			sendKeyPress(a.Keybinds.Ping)
			a.stdDelay()
			sendKeyPress(a.Keybinds.Ping)
			time.Sleep(a.Delays.AutoPingDelay * 2)
		}
		a.stdDelay()
	}
}

func (a *Assist) handleRapidFire() {
	for {
		for a.isRapidFire() {
			click()
			a.stdDelay()
		}
		a.stdDelay()
	}
}

func (a *Assist) Start() {
	go a.handleRecoil()
	go a.handleAutoTag()
	go a.handleRapidFire()
}
