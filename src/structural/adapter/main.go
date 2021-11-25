package main

import "fmt"

func main() {

	tv1 := &SamsungTV{
		currentChan: 13,
		currentVolume: 56,
		tvOn: true,
	}

	tv2 := &SonyTV{
		channel: 20,
		vol: 23,
		isOn: true,
	}

	tv2.turnOn()
	tv2.volumeUp()
	tv2.volumeDown()
	tv2.channelUp()
	tv2.channelDown()
	tv2.goToChannel(65)
	tv2.turnOff()

	fmt.Println("########################################")

	adapter := &samsungAdapter{tv1}
	adapter.turnOn()
	adapter.volumeUp()
	adapter.volumeDown()
	adapter.channelUp()
	adapter.channelDown()
	adapter.goToChannel(65)
	adapter.turnOff()
}
