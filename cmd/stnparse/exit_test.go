package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

/** TestExitCodes builds stnparse and checks the exit codes of DR-0003.
 *
 * Example:
 *   go test ./cmd/stnparse
 */
func TestExitCodes(t *testing.T) {
	bin := filepath.Join(t.TempDir(), "stnparse")
	if out, err := exec.Command("go", "build", "-o", bin, ".").CombinedOutput(); err != nil {
		t.Fatalf("build: %v\n%s", err, out)
	}
	cases := []struct {
		name string
		args []string
		want int
	}{
		{"help", []string{"-h"}, 0},
		{"license", []string{"-l"}, 0},
		{"version", []string{"-v"}, 0},
		{"bad flag", []string{"-bogus"}, 2},
		{"missing input", []string{filepath.Join(t.TempDir(), "missing.txt")}, 66},
		{"uncreatable output", []string{"-", filepath.Join(t.TempDir(), "no", "dir", "out.txt")}, 73},
	}
	for _, c := range cases {
		cmd := exec.Command(bin, c.args...)
		cmd.Stdin, _ = os.Open(os.DevNull)
		err := cmd.Run()
		got := 0
		if ee, ok := err.(*exec.ExitError); ok {
			got = ee.ExitCode()
		} else if err != nil {
			t.Fatalf("%s: %v", c.name, err)
		}
		if got != c.want {
			t.Errorf("%s: exit %d, want %d", c.name, got, c.want)
		}
	}
}
