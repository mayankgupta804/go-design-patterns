package main

type DataSubject struct {
	Observers []DataListener
	Field     string
}

func (ds *DataSubject) ChangeItem(data string) {
	ds.Field = data
	ds.NotifyAll()
}

func (ds *DataSubject) RegisterObserver(obs DataListener) {
	ds.Observers = append(ds.Observers, obs)
}

func (ds *DataSubject) UnRegisterObserver(obs DataListener) {
	var observers []DataListener
	for _, o := range ds.Observers {
		if o.Name != obs.Name {
			observers = append(observers, o)
		}
	}
	ds.Observers = observers
}

func (ds *DataSubject) NotifyAll() {
	for _, o := range ds.Observers {
		o.OnUpdate(ds.Field)
	}
}
