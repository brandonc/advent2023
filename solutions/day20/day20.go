package day20

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/brandonc/advent2023/internal/maths"
	"github.com/brandonc/advent2023/internal/ui"
	"github.com/brandonc/advent2023/solutions/solution"
)

type day20 struct{}

type Pulse rune

const (
	None Pulse = 0
	High Pulse = 'H'
	Low  Pulse = 'L'
)

func Factory() solution.Solver {
	return day20{}
}

type Module interface {
	Receive(pulse Pulse, from string)
	Init(inputs []string)
}

var _ Module = (*BroadcasterModule)(nil)
var _ Module = (*ConjunctionModule)(nil)
var _ Module = (*FlipFlopModule)(nil)

var ErrNotEnoughButtonPresses = errors.New("not enough button presses have been made yet")

type Common struct {
	Name         string
	Destinations []string
	SendBuffer   Pulse
	System       *System
}

type FlipFlopModule struct {
	Common
	On bool
}

type ConjunctionModule struct {
	Common
	Memory        map[string]Pulse
	AllLowPresses int
}

type BroadcasterModule struct {
	Common
}

type System struct {
	Modules       map[string]Module
	TxLow, TxHigh int
	Presses       int

	sendBuffer           []SendBuffer
	rxConjunctionPresses map[string]int
}

func (s *System) CalcuatatePressesUntilRxReceivesLow() (int, error) {
	allPresses := true
	for _, p := range s.rxConjunctionPresses {
		if p == 0 {
			allPresses = false
			break
		}
	}

	if !allPresses {
		return 0, ErrNotEnoughButtonPresses
	}

	return maths.LCM(
		s.rxConjunctionPresses["js"],
		s.rxConjunctionPresses["qs"],
		s.rxConjunctionPresses["dt"],
		s.rxConjunctionPresses["ts"],
	), nil
}
func (s *System) PressButton() {
	s.sendBuffer = append(s.sendBuffer, SendBuffer{From: "button", To: []string{"broadcaster"}, Pulse: Low})
	s.Presses += 1
	s.process()
}

type SendBuffer struct {
	From  string
	To    []string
	Pulse Pulse
}

func (s *System) process() {
	for len(s.sendBuffer) > 0 {
		// Process the buffer of messages to send, but
		sendNow := make([]SendBuffer, len(s.sendBuffer))
		copy(sendNow, s.sendBuffer)
		s.sendBuffer = make([]SendBuffer, 0)

		for i := 0; i < len(sendNow); i++ {
			buffer := sendNow[i]
			for _, destination := range buffer.To {
				dest, ok := s.Modules[destination]
				ui.Debugf("%s -%c-> %s", buffer.From, buffer.Pulse, destination)

				if buffer.Pulse == Low {
					s.TxLow += 1
				} else {
					s.TxHigh += 1
				}

				if !ok {
					continue
				}

				dest.Receive(buffer.Pulse, buffer.From)
			}
		}
	}
}

func (s *System) Send(buffer SendBuffer) {
	s.sendBuffer = append(s.sendBuffer, buffer)
}

func (b *BroadcasterModule) Receive(pulse Pulse, from string) {
	b.System.Send(SendBuffer{From: b.Name, To: b.Destinations, Pulse: pulse})
}

func (b *BroadcasterModule) Init(inputs []string) {
}

func (b *ConjunctionModule) Init(inputs []string) {
	for _, input := range inputs {
		b.Memory[input] = Low
	}
}

func (c *ConjunctionModule) Receive(pulse Pulse, from string) {
	c.Memory[from] = pulse

	send := Low
	for _, mem := range c.Memory {
		if mem == Low {
			send = High
			break
		}
	}

	p, ok := c.System.rxConjunctionPresses[c.Name]
	if c.System.Presses > 0 && send == High && ok && p == 0 {
		c.System.rxConjunctionPresses[c.Name] = c.System.Presses
	}

	c.System.Send(SendBuffer{From: c.Name, To: c.Destinations, Pulse: send})
}

func (f *FlipFlopModule) Receive(pulse Pulse, from string) {
	if pulse != Low {
		return
	}

	send := Low
	if !f.On {
		f.On = true
		send = High
	} else {
		f.On = false
	}

	f.System.Send(SendBuffer{From: f.Name, To: f.Destinations, Pulse: send})
}

func (f *FlipFlopModule) Init(inputs []string) {
}

func parseSystem(reader io.Reader) *System {
	scanner := bufio.NewScanner(reader)
	result := &System{
		Modules: make(map[string]Module),
		rxConjunctionPresses: map[string]int{
			"js": 0,
			"qs": 0,
			"dt": 0,
			"ts": 0,
		},
		sendBuffer: make([]SendBuffer, 0),
	}

	inputs := make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()

		fromTo := strings.Split(line, " -> ")
		destinations := strings.Split(fromTo[1], ", ")

		name := ""
		if fromTo[0] == "broadcaster" {
			name = "broadcaster"
			result.Modules[fromTo[0]] = &BroadcasterModule{
				Common: Common{
					Name:         fromTo[0],
					Destinations: destinations,
					System:       result,
				},
			}
		} else {
			name = fromTo[0][1:]
			if fromTo[0][0] == '&' {
				result.Modules[name] = &ConjunctionModule{
					Common: Common{
						Name:         name,
						Destinations: destinations,
						System:       result,
					},
					Memory: make(map[string]Pulse),
				}
			}

			if fromTo[0][0] == '%' {
				result.Modules[name] = &FlipFlopModule{
					Common: Common{
						Name:         name,
						Destinations: destinations,
						System:       result,
					},
				}
			}
		}

		for _, dest := range destinations {
			inputs[dest] = append(inputs[dest], name)
		}
	}

	for name, module := range result.Modules {
		module.Init(inputs[name])
	}

	return result
}

func (d day20) Part1(reader io.Reader) int {
	system := parseSystem(reader)

	for i := 0; i < 1000; i++ {
		system.PressButton()
	}
	return system.TxHigh * system.TxLow
}

func (d day20) Part2(reader io.Reader) int {
	system := parseSystem(reader)

	for {
		system.PressButton()

		if system.Presses%1_000_000 == 0 {
			fmt.Printf("%d presses...\n", system.Presses)
		}

		result, err := system.CalcuatatePressesUntilRxReceivesLow()
		if errors.Is(err, ErrNotEnoughButtonPresses) {
			continue
		}
		ui.Die(err)

		return result
	}
}
