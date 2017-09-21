package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

const (
	Basso     = "Basso"
	Blow      = "Blow"
	Bottle    = "Bottle"
	Frog      = "Frog"
	Funk      = "Funk"
	Glass     = "Glass"
	Hero      = "Hero"
	Morse     = "Morse"
	Ping      = "Ping"
	Pop       = "Pop"
	Purr      = "Purr"
	Sosumi    = "Sosumi"
	Submarine = "Submarine"
	Tink      = "Tink"
)

const (
	sets           = 2
	repsPerSet     = 5
	prepSeconds    = 10
	hangSeconds    = 12
	restRepMinutes = 3
	restSetMinutes = 5
)

func main() {
	reps := sets * repsPerSet
	for i := 0; i < reps; i++ {
		log.Printf("Get ready for set %d, rep %d!\n", i/repsPerSet+1, i%5+1)
		alert(Sosumi, time.Second*prepSeconds)
		log.Printf("Start hanging for %d seconds\n", hangSeconds)
		alert(Purr, time.Second*hangSeconds)
		if i < reps-1 {
			if (i+1)%repsPerSet == 0 {
				log.Printf("Rest for %d minutes", restSetMinutes)
				time.Sleep(time.Minute * restSetMinutes)
			} else {
				log.Printf("Rest for %d minutes", restRepMinutes)
				time.Sleep(time.Minute * restRepMinutes)
			}
			alert(Ping)
		} else {
			log.Println("Done!")
		}
	}
}

func alert(kind string, runFor ...time.Duration) {
	done := make(chan struct{})

	go func() {
		if len(runFor) > 0 {
			<-time.After(runFor[0])
			close(done)
		} else {
			reader := bufio.NewReader(os.Stdin)
			reader.ReadString('\n')
			close(done)
		}
	}()

	for {
		select {
		case <-time.After(time.Millisecond * 100):
			exec.Command("afplay", fmt.Sprintf("/System/Library/Sounds/%s.aiff", kind)).Run()
		case <-done:
			return
		}
	}
}
