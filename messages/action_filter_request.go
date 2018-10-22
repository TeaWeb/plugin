package messages

import (
	"bufio"
	"bytes"
	"net/http"
)

type FilterRequestAction struct {
	Action

	Continue bool
	Data     []byte
}

func (this *FilterRequestAction) Name() string {
	return "FilterRequest"
}

func (this *FilterRequestAction) Request() (*http.Request, error) {
	return http.ReadRequest(bufio.NewReader(bytes.NewBuffer(this.Data)))
}
