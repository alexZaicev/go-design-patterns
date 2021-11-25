package main

import "fmt"

type SonyTV struct {
	vol     int
	channel int
	isOn    bool
}

func (tv *SonyTV) turnOn() {
	tv.isOn = true
	fmt.Println("SonyTV is now ON")
}

func (tv *SonyTV) turnOff() {
	tv.isOn = false
	fmt.Println("SonyTV is now OFF")
}

func (tv *SonyTV) volumeUp() int {
	tv.vol += 1
	fmt.Println("Increasing SonyTV volume to", tv.vol)
	return tv.vol
}

func (tv *SonyTV) volumeDown() int {
	tv.vol -= 1
	fmt.Println("Decreasing SonyTV volume to", tv.vol)
	return tv.vol
}

func (tv *SonyTV) channelUp() int {
	tv.channel += 1
	fmt.Println("Increasing SonyTV channel to", tv.channel)
	return tv.channel
}

func (tv *SonyTV) channelDown() int {
	tv.channel -= 1
	fmt.Println("Decreasing SonyTV channel to", tv.channel)
	return tv.channel
}

func (tv *SonyTV) goToChannel(ch int) {
	tv.channel = ch
	fmt.Println("Setting SonyTV channel to", tv.channel)
}
