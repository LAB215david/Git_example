package main

import (
	"fmt"
	"time"
  	"github.com/betable/bark"
)


func main() {
	watcher := bark.NewWatchdog(nil, "/usr/share/pphook/PPHookServerSample")
	watcher.Start() // launchs proc, automatically restarts it

	// give it a second to get started
	time.Sleep(2 * time.Millisecond)

	// current Process ID is available here; if 0 then
	// it is between restarts and you should poll again.
	pid := <-watcher.CurrentPid
	fmt.Print(pid)

	// ready to stop both child and watchdog
	// watcher.TermChildAndStopWatchdog <- true

	// wait for shutdown of child and watchdog to complete
	<-watcher.Done
	fmt.Print("watchdog and child pid shutdown.\n")
}