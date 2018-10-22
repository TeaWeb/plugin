package messages

type EmptyAction struct {
	Action
}

func (this *EmptyAction) Name() string {
	return "empty"
}
