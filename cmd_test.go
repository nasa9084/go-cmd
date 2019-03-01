package cmd_test

import (
	"os/exec"
	"strconv"
	"strings"
	"testing"

	cmd "github.com/nasa9084/go-cmd"
)

func TestDf(t *testing.T) {
	out, err := exec.Command("df").Output()
	if err != nil {
		t.Fatal(err)
	}
	vs := strings.Fields(strings.Split(string(out), "\n")[1])
	all, err := strconv.ParseUint(vs[1], 10, 64)
	used, err := strconv.ParseUint(vs[2], 10, 64)
	stat, err := cmd.Df("")
	if err != nil {
		t.Fatal(err)
	}
	if all != stat.All {
		t.Errorf("%d != %d", stat.All, all)
	}
	if used != stat.Used {
		t.Errorf("%d != %d", stat.Used, used)
	}
}
