package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/getlantern/systray"
)

var quit chan struct{}

func main() {
	quit = make(chan struct{})

	go func() {
		onExit := make(chan struct{})
		systray.Run(func() {
			// Icon is not currently working and I honestly couldnt tell you why...
			icon := "C:/Users/Hunter/workspace/github.com/systraySleep/assets/icon.ico"
			iconBytes, err := os.ReadFile(icon)
			if err != nil {
				fmt.Printf("Error reading icon file: %v\n", err)
			} else {
				// Ignore the "success" error on Windows
				systray.SetIcon(iconBytes)
			}
			systray.SetTitle("systraySleep")
			systray.SetTooltip("systraySleep")

			menuItemTime := systray.AddMenuItem("Time", "Select time until sleep")
			go func() {
				twoHours := menuItemTime.AddSubMenuItem("2 hours", "Sleep in 2 hours")
				go func() {
					for {
						<-twoHours.ClickedCh
						fmt.Println("Sleeping now...")
						sleepFunc(2.0)
					}
				}()

				oneHour := menuItemTime.AddSubMenuItem("1 hour", "Sleep in 1 hour")
				go func() {
					for {
						<-oneHour.ClickedCh
						fmt.Println("Sleeping in one hour...")
						sleepFunc(1.0)
					}
				}()

				halfHour := menuItemTime.AddSubMenuItem("30 minutes", "Sleep in 30 minuntes")
				go func() {
					for {
						<-halfHour.ClickedCh
						fmt.Println("Sleeping in 30 minutes...")
						sleepFunc(0.5)
					}
				}()

				now := menuItemTime.AddSubMenuItem("Now", "Testing only, not for final app")
				go func() {
					for {
						<-now.ClickedCh
						sleepFunc(0.0)
					}
				}()

				henryStinks := menuItemTime.AddSubMenuItem("Henry stinks", "Henry is so stinky")
				go func() {
					for {
						<-henryStinks.ClickedCh
						fmt.Println("Henry stinks")
					}
				}()

				menuItemQuit := systray.AddMenuItem("Quit", "Quit the application")
				go func() {
					<-menuItemQuit.ClickedCh
					fmt.Println("Exiting...")
					systray.Quit()
					os.Exit(0)
				}()

				<-onExit
			}()
		}, func() {
			close(onExit)
		})
	}()

	// Handle OS signals
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	fmt.Println("Exiting...")
	close(quit)
	os.Exit(0)

}

func sleepFunc(sleepTime float64) {
	kernel32, err := syscall.LoadLibrary("powrprof.dll")
	if err != nil {
		panic(err)
	}
	defer syscall.FreeLibrary(kernel32)

	go func() {
		time.Sleep(time.Duration(time.Hour.Hours() * sleepTime))
	}()

	// Get the SetSuspendState procedure
	SetSuspendStateProc, err := syscall.GetProcAddress(kernel32, "SetSuspendState")
	if err != nil {
		panic(err)
	}

	// Call the SetSuspendState function to put the computer to sleep
	// The parameters are:
	//   - Hiberbate: If true, the system hibernates; if false, it sleeps
	//   - ForceCritical: If true, forces the sleep/hibernate even if applications prevent it
	//   - DisableWakeEvent: If true, wake events are disabled

	ret, _, err := syscall.SyscallN(uintptr(SetSuspendStateProc), 3, 0, 1, 0)
	if err != syscall.Errno(0) {
		fmt.Printf("Error setting suspend state: %v\n", err)
		return
	}
	if ret == 0 {
		fmt.Println("Failed to put system to sleep")
	}
}
