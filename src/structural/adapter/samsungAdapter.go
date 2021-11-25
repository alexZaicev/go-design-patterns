package main

type samsungAdapter struct {
	tv *SamsungTV
}

func (adapter *samsungAdapter) turnOn() {
	adapter.tv.setOnState(true)
}

func (adapter *samsungAdapter) turnOff() {
	adapter.tv.setOnState(false)
}

func (adapter *samsungAdapter) volumeUp() int {
	adapter.tv.setVolume(adapter.tv.getVolume() + 1)
	return adapter.tv.getVolume()
}

func (adapter *samsungAdapter) volumeDown() int {
	adapter.tv.setVolume(adapter.tv.getVolume() - 1)
	return adapter.tv.getVolume()
}

func (adapter *samsungAdapter) channelUp() int {
	adapter.tv.setChannel(adapter.tv.getChannel() + 1)
	return adapter.tv.getChannel()
}

func (adapter *samsungAdapter) channelDown() int {
	adapter.tv.setChannel(adapter.tv.getChannel() - 1)
	return adapter.tv.getChannel()
}

func (adapter *samsungAdapter) goToChannel(ch int) {
	adapter.tv.setChannel(ch)
}
