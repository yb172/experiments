package load

import (
	"fmt"

	termbox "github.com/nsf/termbox-go"
)

const rpsIncrement = 2

func readKeyboard(exit chan<- interface{}) error {
	err := read()
	exit <- true
	if err != nil {
		return fmt.Errorf("error while reading from keyboard: %v", err)
	}
	return nil
}

var current string
var curev termbox.Event

func read() error {
	if err := termbox.Init(); err != nil {
		return fmt.Errorf("error while initializing console: %v", err)
	}
	defer termbox.Close()

	redrawAll()

	data := make([]byte, 0, 64)
mainloop:
	for {
		if cap(data)-len(data) < 32 {
			newdata := make([]byte, len(data), len(data)+32)
			copy(newdata, data)
			data = newdata
		}
		beg := len(data)
		d := data[beg : beg+32]
		switch ev := termbox.PollRawEvent(d); ev.Type {
		case termbox.EventRaw:
			data = data[:beg+ev.N]
			current = fmt.Sprintf("%s", data)
			if current == "q" {
				break mainloop
			}

			for {
				ev := termbox.ParseEvent(data)
				if ev.N == 0 {
					break
				}
				curev = ev
				copy(data, data[curev.N:])
				data = data[:len(data)-curev.N]
			}
			switch curev.Key {
			case termbox.KeyArrowUp:
				upRps()
			case termbox.KeyArrowDown:
				downRps()
			}

		case termbox.EventError:
			return fmt.Errorf("termbox error: %v", ev.Err)
		}
		redrawAll()
	}
	return nil
}

func redrawAll() {
	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, coldef)
	row := 0
	tbprint(0, row, coldef, coldef, fmt.Sprintf("Generating load on service @ %v RPS", rate))
	row++
	tbprint(0, row, termbox.ColorGreen, coldef, "Use arrows to increase / decrease RPS")
	row++
	tbprint(0, row, termbox.ColorMagenta, coldef, "Press 'q' to quit")
	row++
	termbox.Flush()
}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func upRps() {
	rate += rpsIncrement
}

func downRps() {
	if rate > 0 {
		rate -= rpsIncrement
	}
	if rate < 0 {
		rate = 0
	}
}
