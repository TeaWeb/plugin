package messages

type StartAction struct {
	Action
}

func (this *StartAction) Name() string {
	return "Start"
}
