package probes

import (
	"github.com/TeaWeb/plugin/pkg/apps"
	"regexp"
	"testing"
)

func TestProcessProbe_Run(t *testing.T) {
	probe := NewProcessProbe()
	probe.Name = "PHP-FPM"
	probe.Site = "http://php.net/"
	probe.DocSite = "http://php.net/docs.php"
	probe.Developer = "The PHP Group"
	probe.CommandName = "php-fpm"
	probe.CommandPatterns = []string{}
	probe.CommandVersion = "${commandFile} -v"
	probe.OnParseVersion(func(versionString string) (string, error) {
		reg := regexp.MustCompile(`PHP \d+\.\d+\.\d+`)
		if reg.MatchString(versionString) {
			return reg.FindStringSubmatch(versionString)[0], nil
		}
		return versionString, nil
	})
	probe.OnProcess(func(process *apps.Process) bool {
		return true
	})
	resultApps, _ := probe.Run()
	t.Log(resultApps)

	if len(resultApps) > 0 {
		t.Log(resultApps[0].Version)
	}
}
