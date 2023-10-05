package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	hook "github.com/robotn/gohook"
	"github.com/getlantern/systray"
	"github.com/tajtiattila/vjoy"
)

//go:embed build/icon-systray.ico
var icon []byte

var joy *vjoy.Device
var joyNum = uint(1)
var pullRange = 1000

func main() {
	initializeArgs()
	initializeJoy()
	systray.Run(onSystrayReady, onSystrayExit)
}

func initializeArgs() {
	flag.UintVar(&joyNum, "joyNum", 1, "Joy number")
	flag.IntVar(&pullRange, "pullDistance", 1000, "Pull distance")

	flag.Parse()
}

func onSystrayReady() {
	systray.SetIcon(icon)
	systray.SetTitle("Mouse Plunger")
	mQuit := systray.AddMenuItem("Quit", "Quit application")
	

	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
		os.Exit(0)
	}()
	
	registerHooks()
}

func onSystrayExit() {
	cleanup()
}

func cleanup() {
	joy.Reset()
	joy.Relinquish()
}

func initializeJoy() {
	var err error
	joy, err = vjoy.Acquire(joyNum)
	if err != nil {
		panic(fmt.Sprintf("Couldn't find vJoy device %d. Is it configured?", joyNum))
	}
	setJoyZ(0)
	update()
}
func registerHooks() {
	var initialY int16
	var active = false
	hook.Register(hook.MouseHold, []string{}, func(e hook.Event) {
		if e.Button == hook.MouseMap["left"] {
			active = true
			initialY = e.Y
		}
	})

	hook.Register(hook.MouseDrag, []string{}, func(e hook.Event) {
		if active {
			yDistance := initialY - e.Y
			setJoyZ(int(float32(-yDistance)/float32(pullRange) * 0x7fff))

			update()
		}
	})

	hook.Register(hook.MouseUp, []string{}, func(e hook.Event) {
		if e.Button == hook.MouseMap["left"] {
			setJoyZ(0)
			active = false

			update()
		}
	})

	s := hook.Start()
	<-hook.Process(s)
}

func setJoyZ(val int) {
	joy.Axis(vjoy.AxisZ).Setc(val)
}

func update() {
	if err := joy.Update(); err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}
}