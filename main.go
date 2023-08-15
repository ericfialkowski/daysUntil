package main

import (
	"daysUntil/keys"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/buger/goterm"
	"github.com/hako/durafmt"
)

const DATEFORMAT = "01-02-2006"
const DATETIMEFORMAT = "01-02-2006 15:04"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Include at least one date in the form of mm-dd-yyyy as a program argument")
		os.Exit(1)
	}

	keys.ExitOnAnyKey()

	ticker := time.NewTicker(1 * time.Second)

	clearCounter := 0
	linesPrinted := 0
	go func() {
		for t := range ticker.C {
			clearCounter--
			if clearCounter < 0 {
				goterm.Clear()
				clearCounter = rand.Intn(5) + 5
			}
			goterm.MoveCursor(1, 1)
			_, _ = goterm.Printf("Current time: %v", t.Format(time.ANSIC))

			y := 2
			for _, s := range os.Args[1:] {
				format := DATEFORMAT
				if len(s) == len(DATETIMEFORMAT) {
					format = DATETIMEFORMAT
				}
				date, err := time.ParseInLocation(format, s, t.Location())
				y = y + 1
				goterm.MoveCursor(1, y)
				if err != nil {
					_, _ = goterm.Printf("%s is not a date in the form of mm-dd-yyyy or mm-dd-yyyy hh:mm\n", s)
					goterm.Flush()
					os.Exit(2)
				}

				c := date.Compare(t)

				switch c {
				case 1:
					diff := date.Sub(t)
					_, _ = goterm.Printf("%30v until %v                  ", durafmt.Parse(diff).LimitFirstN(3), date.Format(time.ANSIC))
					linesPrinted++
				case -1:
					diff := t.Sub(date)
					_, _ = goterm.Printf("%30v since %v                  ", durafmt.Parse(diff).LimitFirstN(3), date.Format(time.ANSIC))
					linesPrinted++
				case 0:
					_, _ = goterm.Printf("%30v is now                  ", date.Format(time.ANSIC))
					linesPrinted++
				}
			}

			if linesPrinted > y {
				for z := y + 1; z < linesPrinted; z++ {
					goterm.MoveCursor(1, z)
					_, _ = goterm.Print("                                                             ")
				}
			}
			linesPrinted = y
			_, _ = goterm.Println()
			goterm.Flush()
		}
	}()

	select {}
}
