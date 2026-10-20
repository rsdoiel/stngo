package cli

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func runParse(a string, args []string, in *strings.Reader, out, errOut *bytes.Buffer) int {
	return Parse(a, args, in, out, errOut)
}
func runFilter(a string, args []string, in *strings.Reader, out, errOut *bytes.Buffer) int {
	return Filter(a, args, in, out, errOut)
}
func runReport(a string, args []string, in *strings.Reader, out, errOut *bytes.Buffer) int {
	return Report(a, args, in, out, errOut)
}

func TestStandardOptionsInEverySpelling(t *testing.T) {
	progs := map[string]func(string, []string, *strings.Reader, *bytes.Buffer, *bytes.Buffer) int{
		"stnparse": runParse, "stnfilter": runFilter, "stnreport": runReport,
	}
	for name, f := range progs {
		for _, opt := range []string{"-h", "-help", "--help"} {
			out, _, code := callNamed(t, f, name, "", opt)
			if code != 0 || !strings.Contains(out, name+"(1) user manual") || !strings.Contains(out, "# EXIT STATUS") {
				t.Errorf("%s %s: exit %d\n%s", name, opt, code, out)
			}
			if strings.Contains(out, "{app_name}") || strings.Contains(out, "{version}") || strings.Contains(out, "{release_date}") || strings.Contains(out, "{release_hash}") {
				t.Errorf("%s %s: a help token was not replaced\n%s", name, opt, out)
			}
		}
		for _, opt := range []string{"-v", "-version", "--version"} {
			out, _, code := callNamed(t, f, name, "", opt)
			if code != 0 || !strings.HasPrefix(out, name+" 0.") {
				t.Errorf("%s %s: exit %d %q", name, opt, code, out)
			}
		}
		for _, opt := range []string{"-l", "-license", "--license"} {
			out, _, code := callNamed(t, f, name, "", opt)
			if code != 0 || !strings.Contains(out, "Redistribution and use") {
				t.Errorf("%s %s: exit %d %q", name, opt, code, out)
			}
		}
	}
}

func callNamed(t *testing.T, f func(string, []string, *strings.Reader, *bytes.Buffer, *bytes.Buffer) int, name string, stdin string, args ...string) (string, string, int) {
	t.Helper()
	var out, errOut bytes.Buffer
	code := f(name, args, strings.NewReader(stdin), &out, &errOut)
	return out.String(), errOut.String(), code
}

func TestExitCodes(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "missing.txt")
	noDir := filepath.Join(t.TempDir(), "no", "dir", "out.txt")
	notStn := "8:00 - 9:00; A; b\n"
	cases := []struct {
		name  string
		f     func(string, []string, *strings.Reader, *bytes.Buffer, *bytes.Buffer) int
		stdin string
		args  []string
		want  int
	}{
		{"parse bad flag", runParse, "", []string{"-bogus"}, 2},
		{"parse surplus argument", runParse, "", []string{"a", "b", "c"}, 2},
		{"parse missing input", runParse, "", []string{missing}, 66},
		{"parse cannot create output", runParse, "", []string{"-", noDir}, 73},
		{"filter bad flag", runFilter, "", []string{"-bogus"}, 2},
		{"filter bad start", runFilter, "", []string{"-start=yesterday"}, 2},
		{"filter bad end", runFilter, "", []string{"-end=soon"}, 2},
		{"filter missing input", runFilter, "", []string{missing}, 66},
		{"filter not stnparse output", runFilter, notStn, nil, 65},
		{"report bad flag", runReport, "", []string{"-bogus"}, 2},
		{"report bad columns", runReport, "", []string{"-columns=a"}, 2},
		{"report negative column", runReport, "", []string{"-columns=-1"}, 2},
		{"report missing input", runReport, "", []string{missing}, 66},
		{"report not stnparse output", runReport, notStn, nil, 65},
	}
	for _, c := range cases {
		if _, _, code := callNamed(t, c.f, "prog", c.stdin, c.args...); code != c.want {
			t.Errorf("%s: exit %d, want %d", c.name, code, c.want)
		}
	}
}

func TestParseWritesJSONAndFiles(t *testing.T) {
	src := "2026-10-13\n8:00 - 9:00; Acme; Widgets; one\n9:00 - 10:00; Acme; Widgets; two\n"
	out, _, code := callNamed(t, runParse, "stnparse", src, "-json")
	if code != 0 || !strings.HasPrefix(out, "[{") || !strings.HasSuffix(out, "}]") || strings.Count(out, "},{") != 1 {
		t.Errorf("json: exit %d %q", code, out)
	}
	dir := t.TempDir()
	outName := filepath.Join(dir, "out.tsv")
	if _, _, code := callNamed(t, runParse, "stnparse", src, "-o", outName); code != 0 {
		t.Fatalf("-o: exit %d", code)
	}
}

func TestFilterOpenEndedRanges(t *testing.T) {
	loc := time.Local
	defer func() { time.Local = loc }()
	time.Local = time.UTC
	src := "2026-10-12T08:00:00Z\t2026-10-12T09:00:00Z\tA\tb\n2026-10-14T08:00:00Z\t2026-10-14T09:00:00Z\tA\tc\n"
	out, _, _ := callNamed(t, runFilter, "stnfilter", src, "-end=2026-10-13")
	if strings.Count(out, "\n") != 1 || !strings.Contains(out, "10-12") {
		t.Errorf("-end alone should keep what is before it:\n%s", out)
	}
	out, _, _ = callNamed(t, runFilter, "stnfilter", src, "-start=2026-10-13")
	if strings.Count(out, "\n") != 1 || !strings.Contains(out, "10-14") {
		t.Errorf("-start alone should keep what is after it:\n%s", out)
	}
	out, _, _ = callNamed(t, runFilter, "stnfilter", src, "-match=c")
	if strings.Count(out, "\n") != 1 || !strings.Contains(out, "10-14") {
		t.Errorf("-match:\n%s", out)
	}
}
