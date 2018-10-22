package messages

type FilterResponseAction struct {
	Action
}

func (this *FilterResponseAction) Name() string {
	return "FilterResponse"
}
