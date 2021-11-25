package main

type observable interface {
	registerObserver(obs observer)
	unregisterObserver(obs observer)
	notifyAll()
}

type DataSubject struct {
	observers []*DataListener
	field     string
}

func (ds *DataSubject) ChangeItem(data string) {
	ds.field = data
	ds.notifyAll()
}

func (ds *DataSubject) registerObserver(o *DataListener) {
	ds.observers = append(ds.observers, o)
}

func (ds *DataSubject) unregisterObserver(o *DataListener) {
	var newObs []*DataListener
	for _, listener := range ds.observers {
		if listener.Name != o.Name {
			newObs = append(newObs, listener)
		}
	}
	ds.observers = newObs
}

func (ds *DataSubject) notifyAll() {
	for _, listener := range ds.observers {
		listener.onUpdate(ds.field)
	}
}
