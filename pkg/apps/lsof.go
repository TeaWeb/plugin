package apps

import (
	"errors"
	"fmt"
	"github.com/TeaWeb/plugin/internal/utils/types"
	"strings"
)

func Lsof(pid int32) (results []*LsofResult, err error) {
	resultString, err := Exec("lsof", "-Pan", "-FctnLpPT0", "-p", fmt.Sprintf("%d", pid))
	if err != nil {
		return
	}

	if len(resultString) == 0 {
		err = errors.New("lsof fail")
		return
	}

	for _, line := range strings.Split(resultString, "\n") {
		pieces := strings.Split(line, string([]byte{0}))
		result := &LsofResult{}
		for _, piece := range pieces {
			if len(piece) == 0 {
				continue
			}
			switch piece[0] {
			case 'p':
				result.Pid = types.Int32(piece[1:])
			case 'f':
				result.Fd = piece[1:]
			case 'L':
				result.User = piece[1:]
			case 'P':
				result.Protocol = piece[1:]
			case 'c':
				result.Command = piece[1:]
			case 't':
				result.Type = piece[1:]
			case 'n':
				result.Name = piece[1:]
			case 'T':
				if strings.HasPrefix(piece[1:], "ST=") {
					result.ConnectionState = piece[4:]
				}
			}
		}
		results = append(results, result)
	}

	if len(results) == 0 {
		err = errors.New("no process matched")
	}

	return
}

type LsofResult struct {
	Command         string
	Pid             int32
	User            string
	Fd              string
	Type            string
	Protocol        string
	Name            string
	ConnectionState string
}

func (this *LsofResult) IsCwd() bool {
	return this.Fd == "cwd"
}

func (this *LsofResult) IsRegularFile() bool {
	return this.Type == "REG"
}

func (this *LsofResult) IsIPv4() bool {
	return this.Type == "IPv4"
}

func (this *LsofResult) IsIPv6() bool {
	return this.Type == "IPv6"
}

func (this *LsofResult) IsTCP() bool {
	return this.Protocol == "TCP"
}

func (this *LsofResult) IsListening() bool {
	return this.ConnectionState == "LISTEN"
}

func (this *LsofResult) Listen() string {
	if !this.IsListening() {
		return ""
	}
	return this.Name
}

func (this *LsofResult) IsEstablished() bool {
	return this.ConnectionState == "ESTABLISHED"
}

func (this *LsofResult) LAddr() string {
	pieces := strings.Split(this.Name, "->")
	if len(pieces) == 2 {
		return pieces[0]
	}
	return ""
}

func (this *LsofResult) RAddr() string {
	pieces := strings.Split(this.Name, "->")
	if len(pieces) == 2 {
		return pieces[1]
	}
	return ""
}
