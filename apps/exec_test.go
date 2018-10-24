package apps

import "testing"

func TestExec(t *testing.T) {
	output, err := Exec("/opt/nginx/sbin/nginx", "-v")
	t.Log(output, err)
}
