package messages

type ReloadAction struct {
	Action
}

func (this *ReloadAction) Name() string {
	return "Reload"
}
