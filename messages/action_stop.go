package messages

type StopAction struct {
	Action
}

func (this *StopAction) Name() string {
	return "Stop"
}
