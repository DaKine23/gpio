package gpio

import (
	"errors"
)

const (
	OUT  = "out"
	IN   = "in"
	ONE  = "1"
	ZERO = "0"
)

type GPIO_Pin struct {
	Port      string
	Direction Direction
	Value     Value
}

type Value bool
type Direction bool

var nilPinErr = errors.New("GPIO_Pin was not initialized (nil)")

func (o *GPIO_Pin) Write(flowID string) error {

	if o == nil {
		return nilPinErr
	}
	var llgpio GPIO_ll
	pin := llgpio.Pin(flowID, o.Port)

	// set Direction
	if o.Direction {
		pin.Output(flowID)
	} else {
		pin.Input(flowID)
	}

	// set Value
	if o.Value {
		pin.High(flowID)
	} else {
		pin.Low(flowID)
	}

	return nil

}

type GPIO struct {
}
type GPIO_LedSet struct {
	Set []GPIO_Pin
}

func (o *GPIO_LedSet) Oostring() (result string) {
	for _, v := range o.Set {
		if v.Value {
			result += "O"
		} else {
			result += "o"
		}
	}
	return
}

func (o *GPIO) NewLedSet(ports ...string) *GPIO_LedSet {

	ledSet := GPIO_LedSet{}
	ledSet.Set = make([]GPIO_Pin, 0, len(ports))
	for _, v := range ports {
		ledSet.Set = append(ledSet.Set, GPIO_Pin{v, true, false})
	}
	return &ledSet
}

func (o *GPIO_LedSet) AllOn(flowID string) {
	for i, _ := range o.Set {
		o.Set[i].SetValue(true)
	}
	for _, v := range o.Set {
		v.Write(flowID)
	}
}
func (o *GPIO_LedSet) AllOff(flowID string) {
	for i, _ := range o.Set {
		o.Set[i].SetValue(false)

	}
	for _, v := range o.Set {
		v.Write(flowID)
	}
}
func (o *GPIO_LedSet) AllSwitch(flowID string) {
	for i, _ := range o.Set {
		o.Set[i].SwitchValue()

	}
	for _, v := range o.Set {
		v.Write(flowID)
	}
}
func (o *GPIO_LedSet) SingleOn(flowID, port string) {
	for i, v := range o.Set {
		if v.Port == port {
			o.Set[i].SetValue(true)
			o.Set[i].Write(flowID)
		}
	}
}
func (o *GPIO_LedSet) SingleOff(flowID, port string) {
	for i, v := range o.Set {
		if v.Port == port {
			o.Set[i].SetValue(false)
			o.Set[i].Write(flowID)
		}
	}
}
func (o *GPIO_LedSet) SingleSwitch(flowID, port string) {
	for i, v := range o.Set {
		if v.Port == port {
			o.Set[i].SwitchValue()
			o.Set[i].Write(flowID)
		}
	}
}
func (o *GPIO_LedSet) Move(flowID string, direction bool) {

	var j int
	var temp Value
	if direction {
		temp = o.Set[len(o.Set)-1].Value
	} else {
		temp = o.Set[0].Value
	}

	for i := 0; i < len(o.Set); i++ {
		if direction {
			j = i
		} else {
			j = len(o.Set) - i - 1
		}
		temp, o.Set[j].Value = o.Set[j].Value, temp
		o.Set[j].Write(flowID)
	}
}

func (o *GPIO_Pin) SwitchValue() *GPIO_Pin {

	o.Value = !o.Value
	return o
}
func (o *GPIO_Pin) SwitchDirection() *GPIO_Pin {

	o.Direction = !o.Direction
	return o
}

func (o *GPIO_Pin) SetValue(value bool) *GPIO_Pin {
	o.Value = Value(value)
	return o
}

func (o *GPIO_Pin) SetOutput() *GPIO_Pin {

	o.Direction = true
	return o
}
func (o *GPIO_Pin) SetInput() *GPIO_Pin {

	o.Direction = false
	return o
}

func (d Direction) String() string {
	if d {
		return OUT
	}
	return IN
}
func (v Value) String() string {
	if v {
		return ONE
	}
	return ZERO

}
