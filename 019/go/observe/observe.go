package observe

type Observer interface {
	Notify()
}

type EventA struct {
	obs map[Observer]int
}

func (ai *EventA) Subscribe(ob Observer) {
	if nil == ai.obs {
		ai.obs = make(map[Observer]int)
	}
	ai.obs[ob] = 0
}
func (ai *EventA) UnSubscribe(ob Observer) {
	delete(ai.obs, ob)
}
func (ai EventA) NotifyAll() {
	for i, _ := range ai.obs {
		go i.Notify()
	}
}
