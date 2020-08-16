package messages

import (
	"bufio"
	"bytes"
	"net/http"
)

type FilterResponseAction struct {
	Action

	Continue bool
	Data     []byte
}

func (this *FilterResponseAction) Name() string {
	return "FilterResponse"
}

func (this *FilterResponseAction) Response() (*http.Response, error) {
	return http.ReadResponse(bufio.NewReader(bytes.NewBuffer(this.Data)), nil)
}
