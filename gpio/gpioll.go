package gpio

import (
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

//GPIO_ll is a handle for low level gpio object creation
type GPIO_ll struct{}

//Pin creates a new Pin for low lever control
func (r GPIO_ll) Pin(flowID, port string) GPIO_Pin_ll {
	pin := GPIO_Pin_ll{port, nil}
	filename := pin.Filename()
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// export gpio pin
		err := ioutil.WriteFile("/sys/class/gpio/export", []byte(pin.Port), 0666)
		pin.addError(err)
	}
	return pin
}

type GPIO_Pin_ll struct {
	Port string
	Err  error
}

func (r GPIO_Pin_ll) addError(err error) {

	if err == nil {
		return
	}
	if r.Err == nil {
		r.Err = err
		return
	}

	r.Err = errors.Wrap(r.Err, err.Error())

}

func (r GPIO_Pin_ll) Filename() string {
	return "/sys/class/gpio/gpio" + r.Port
}
func (r GPIO_Pin_ll) write(flowID, where, what string) GPIO_Pin_ll {
	filename := r.Filename() + "/" + where
	err := ioutil.WriteFile(filename, []byte(what), 0666)
	r.addError(err)
	return r
}
func (r GPIO_Pin_ll) Output(flowID string) GPIO_Pin_ll {
	return r.write(flowID, "direction", "out")
}
func (r GPIO_Pin_ll) Input(flowID string) GPIO_Pin_ll {
	return r.write(flowID, "direction", "in")
}
func (r GPIO_Pin_ll) High(flowID string) GPIO_Pin_ll {
	return r.write(flowID, "value", "1")
}
func (r GPIO_Pin_ll) Low(flowID string) GPIO_Pin_ll {
	return r.write(flowID, "value", "0")
}
