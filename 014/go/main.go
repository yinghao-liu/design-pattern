package main

import "fmt"

// 请求者
type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

/*****************命令接口*****************************/
type command interface {
	execute()
}

// 具体接口-on
type CommandOn struct {
	device device
}

func (c *CommandOn) execute() {
	c.device.on()
}

// 具体接口-off
type CommandOff struct {
	device device
}

func (c *CommandOff) execute() {
	c.device.off()
}

// 接收者接口
type device interface {
	on()
	off()
}

// 具体接收者
type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

// client
func main() {
	tv := &tv{}

	Commandon := &CommandOn{
		device: tv,
	}

	Commandoff := &CommandOff{
		device: tv,
	}

	ButtonOn := &button{
		command: Commandon,
	}
	ButtonOn.press()

	ButtonOff := &button{
		command: Commandoff,
	}
	ButtonOff.press()
}
