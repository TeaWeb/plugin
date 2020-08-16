package apps

import "testing"

func TestLsofPid(t *testing.T)  {
	results, _ := Lsof(58682)
	for _, result := range results {
		t.Logf("%#v", result)
		t.Log("isSocket", result.IsIPv4(), result.IsIPv6(), result.IsTCP(), result.LAddr(), "->", result.RAddr())
	}
}
