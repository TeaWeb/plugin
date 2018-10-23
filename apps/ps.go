package apps

import (
	"errors"
	"fmt"
	"github.com/TeaWeb/plugin/utils/types"
	"log"
	"regexp"
	"strings"
	"time"
)

func PsLookup(lookup string, matchPatterns []string, onlyParent bool) (result []*Process, err error) {
	resultString, err := Exec("pgrep", "-f", lookup)
	if err != nil {
		return
	}

	if len(resultString) == 0 {
		err = errors.New("process not found")
		return
	}

	for _, pidString := range strings.Split(resultString, "\n") {
		p, err := PsPid(types.Int32(pidString))
		if err != nil {
			continue
		}

		// check parent
		if onlyParent {
			if p.Ppid > 128 {
				break
			}
		}

		if len(matchPatterns) > 0 {
			matched := true
			pieces := ParseArgs(p.Cmdline)
			for _, pattern := range matchPatterns {
				reg, err := regexp.Compile(pattern)
				if err != nil {
					log.Println(err.Error())
					matched = false
					break
				}

				found := false
				for _, piece := range pieces {
					if reg.MatchString(piece) {
						found = true
						break
					}
				}

				if !found {
					matched = false
					break
				}
			}
			if !matched {
				continue
			}
		}

		result = append(result, p)
	}

	if len(result) == 0 {
		err = errors.New("process not found")
	}

	return
}

func PsPid(pid int32) (*Process, error) {
	if pid < 0 {
		return nil, errors.New("pid should not small than '0'")
	}

	keywords := "user=,pid=,ppid=,%cpu=,uid=,gid=,lstart=,%mem=,rss=,vsize=,command="
	patterns := `^(?U)(\S+)\s+(\S+)\s+(\S+)\s+(\S+)\s+(\S+)\s+(\S+)\s+(\S+.+\d{4})\s+(\S+)\s+(\S+)\s+(\S+)\s+(.+)\n?$`

	resultString, err := Exec("ps", "-p", fmt.Sprintf("%d", pid), "-o", keywords)
	if err != nil {
		return nil, err
	}

	pidString := fmt.Sprintf("%d", pid)

	if len(resultString) == 0 {
		return nil, errors.New("process '" + pidString + "' not found")
	}

	matches := regexp.MustCompile(patterns).FindStringSubmatch(resultString)
	if len(matches) <= 1 {
		if len(resultString) == 0 {
			return nil, errors.New("process '" + fmt.Sprintf("%d", pid) + "' not found")
		}
	}

	p := NewProcess(pid)
	p.Pid = pid
	p.User = matches[1]
	p.Ppid = types.Int32(matches[3])
	p.CPUUsage = &CPUUsage{
		Percent: types.Float64(matches[4]),
	}
	p.Uid = types.Int32(matches[5])
	p.Gid = types.Int32(matches[6])

	t, err := time.Parse("Mon Jan _2 15:04:05 2006", matches[7])
	if err == nil {
		p.CreateTime = t.Unix()
	}

	p.MemoryUsage = &MemoryUsage{
		Percent: types.Float64(matches[8]),
		RSS:     types.Uint64(matches[9]) * 1024,
		VMS:     types.Uint64(matches[10]) * 1024,
	}
	p.Cmdline = matches[11]

	// name
	{
		resultString, err := Exec("ps", "-c", "-p", pidString, "-o", "command=")
		if err == nil && len(resultString) > 0 {
			p.Name = resultString
		}
	}

	// cwd
	{
		resultString, err := Exec("pwdx", pidString)
		if err == nil && len(pidString) > 0 {
			index := strings.Index(resultString, ":")
			p.Cwd = strings.TrimPrefix(resultString[index+1:], " ")
		} else {
			resultString, err := Exec("lsof", "-a", "-d", "cwd", "-p", pidString, "-Fn")
			if err == nil && len(resultString) > 0 {
				nIndex := strings.Index(resultString, "\nn")
				if nIndex > 0 {
					p.Cwd = resultString[nIndex+2:]
				}
			}
		}
	}

	return p, nil
}

func PsChildren(parentPid int32) (result []*Process, err error) {
	resultString, err := Exec("pgrep", "-P", fmt.Sprintf("%d", parentPid))
	if err != nil {
		return []*Process{}, err
	}
	if len(resultString) == 0 {
		return
	}

	for _, pidString := range strings.Split(resultString, "\n") {
		pid := types.Int32(pidString)
		if pid < 0 {
			continue
		}

		p, err := PsPid(pid)
		if err != nil {
			continue
		}
		result = append(result, p)
	}

	return
}
