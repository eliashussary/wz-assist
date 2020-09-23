# wz-assist

> COD Warzone recoil, rapid fire, and auto ping assist.
> Built for learning purposes.

## How to Use

Download the latest release from https://github.com/eliashussary/wz-assist/releases/latest

Open a `powershell` window or `cmd` and run:

```sh
wz-assist.exe run
```

## Features & Config

## Features

- **Recoil Compensation**

  - Recoil compensation is on while `[NUM LOCK]` is off
  - Recoil compensation only occurs while aiming down sights.
  - Recoil compensation speed is based on the `recoilsense` settings
  - Recoil compensation speed is toggled through `[SCROLL LOCK]`

- **Auto Ping**

  - While firing down sight, auto ping will ping where you're shooting and spot an enemy

- **Rapid Fire**
  - For single shot weapons, you can fire rapidly with the `[BROWSER FORWARD]` button on your mouse

### Config

Make changes to `.wz-assist.yaml` as you see necessary. The defaults are mostly sane.

> Tip: if you're uncertain what your key maps to, you can run `wz-assist.exe getKey` to print the key number. You can take this number and add it to your config.

```yaml
# keybinds are virtual keys from the win32 api
# you can supply them as an integer or hex value found below
# use `wz-assist getKey` to help with getting a key int
# https://docs.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes
keybinds:
  # VK_RBUTTON: [RIGHT CLICK]
  ads: 0x02

  # VK_LMENU: [LEFT ALT]
  ping: 0xA4

  # VK_XBUTTON2: [MOUSE BROWSER FORWARD]
  rapidfiretoggle: 0x06

  # toggles recoil handling
  # when NUM LOCK is on recoil will not be handled
  # VK_NUMLOCK: [NUM LOCK]
  recoiltoggle: 0x90

  # toggles between high and low sensitivity
  # when SCROLL LOCK is on high sense is used
  # VK_SCROLL: [SCROLL LOCK]
  recoilsensetoggle: 0x91

# recoilsense controls how fast the mouse moves down the Y axis while firing
# higher = slower
# lower = faster
recoilsense:
  low: 20ms
  high: 14ms

# delays are values that determine how fast a routine runs
delays:
  # this value should not change as it serves as a global routine delay
  # change at your own risk
  standard: 20ms

  # the frequency of how often a ping occurs after firing your weapon
  # it will ping continually while firing
  autopingdelay: 750ms

  # how fast a single fire weapon fires when rapid fire is triggered
  # change at your own risk
  rapidfiredelay: 20ms
```
