package main

import (
	"github.com/TeaWeb/plugin/apps"
	"github.com/TeaWeb/plugin/apps/probes"
	"github.com/TeaWeb/plugin/loader"
	"github.com/TeaWeb/plugin/plugins"
	"log"
	"regexp"
	"strings"
)

func main() {
	p := plugins.NewPlugin()
	p.Name = "Apps"
	p.Description = "默认集成的一些本地服务探测器"
	p.Version = "v0.0.1"
	p.Developer = "TeaWeb"
	p.Site = "https://github.com/TeaWeb/build"
	p.Code = "apps.teaweb"
	p.Date = "2018-12-15"

	{
		probe := probes.NewProcessProbe()
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
		resultApps, _ := probe.Run()
		p.AddApp(resultApps ...)
	}

	{
		probe := probes.NewProcessProbe()
		probe.Name = "Redis"
		probe.Site = "https://redis.io/"
		probe.DocSite = "https://redis.io/documentation"
		probe.Developer = "redislabs"
		probe.CommandName = "redis-server"
		probe.CommandPatterns = []string{}
		probe.CommandVersion = "${commandFile} -v"
		probe.OnParseVersion(func(versionString string) (string, error) {
			return versionString, nil
		})
		resultApps, _ := probe.Run()
		p.AddApp(resultApps ...)
	}

	{
		probe := probes.NewProcessProbe()
		probe.Name = "MongoDB"
		probe.Site = "https://www.mongodb.com/"
		probe.DocSite = "https://docs.mongodb.com/"
		probe.Developer = "MongoDB, Inc"
		probe.CommandName = "mongod"
		probe.CommandPatterns = []string{"/mongod"}
		probe.CommandVersion = "${commandFile} --version"
		probe.OnParseVersion(func(versionString string) (string, error) {
			result := regexp.MustCompile("version (v\\S+)").FindStringSubmatch(versionString)
			if len(result) > 0 {
				return result[1], nil
			}
			return versionString, nil
		})
		resultApps, _ := probe.Run()
		p.AddApp(resultApps ...)
	}

	{
		probe := probes.NewProcessProbe()
		probe.Name = "nginx"
		probe.Site = "http://nginx.org/"
		probe.DocSite = "http://nginx.org/en/docs/"
		probe.Developer = "nginx.org"
		probe.CommandName = "nginx"
		probe.CommandPatterns = []string{}
		probe.CommandVersion = "${commandFile} -v"
		probe.OnParseVersion(func(versionString string) (string, error) {
			index := strings.Index(versionString, "nginx version:")
			if index > -1 {
				return versionString[len("nginx version:"):], nil
			}

			return versionString, nil
		})
		resultApps, _ := probe.Run()
		p.AddApp(resultApps ...)
	}

	{
		probe := probes.NewProcessProbe()
		probe.Name = "MySQL"
		probe.Site = "https://www.mysql.com/"
		probe.DocSite = "https://dev.mysql.com/doc/"
		probe.Developer = "Oracle Corporation"
		probe.CommandName = "mysqld_safe"
		probe.CommandPatterns = []string{"mysqld_safe$"}
		probe.CommandVersion = "${commandDir}/mysqld -V"
		probe.OnParseVersion(func(versionString string) (string, error) {
			index := strings.Index(versionString, "Ver ")
			if index > -1 {
				return versionString[index:], nil
			}
			return versionString, nil
		})
		probe.OnProcess(func(p *apps.Process) bool {
			p.ChangeName("mysqld_safe")
			return true
		})
		resultApps, _ := probe.Run()
		p.AddApp(resultApps ...)
	}

	log.Println("[apps]", len(p.Apps), "apps")

	loader.Start(p)
}

func printApps(resultApps []*apps.App) {
	for _, resultApp := range resultApps {
		log.Println(resultApp)
		log.Println(resultApp.Version)
	}
}
