package messages

type LoadAction struct {
	Action
}

func (this *LoadAction) Name() string {
	return "Load"
}
