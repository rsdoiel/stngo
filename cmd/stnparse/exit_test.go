package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

/** TestUnparsableLineDoesNotCrash feeds stnparse an entry line that fails to
 * parse. It must report the line, keep going, parse the lines after it and
 * exit 65 (data) instead of dereferencing a nil entry.
 *
 * Example:
 *   go test -run TestUnparsableLineDoesNotCrash ./cmd/stnparse
 */
func TestUnparsableLineDoesNotCrash(t *testing.T) {
	bin := filepath.Join(t.TempDir(), "stnparse")
	if out, err := exec.Command("go", "build", "-o", bin, ".").CombinedOutput(); err != nil {
		t.Fatalf("build: %v\n%s", err, out)
	}
	src := "2015-07-04\n25:61 - 26:00; DLD; Misc; bad\n08:00 - 09:00; DLD; Misc; good\n"
	cmd := exec.Command(bin)
	cmd.Stdin = strings.NewReader(src)
	var stdout, stderr strings.Builder
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	err := cmd.Run()
	ee, ok := err.(*exec.ExitError)
	if !ok || ee.ExitCode() != 65 {
		t.Errorf("exit = %v, want 65\nstderr: %s", err, stderr.String())
	}
	if strings.Contains(stderr.String(), "panic") {
		t.Errorf("stnparse panicked:\n%s", stderr.String())
	}
	if !strings.Contains(stderr.String(), "line 2") {
		t.Errorf("stderr should name line 2, got %q", stderr.String())
	}
	if !strings.Contains(stdout.String(), "good") {
		t.Errorf("the line after the bad one was not parsed: %q", stdout.String())
	}
}
