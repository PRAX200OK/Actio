package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"actio/internal/project"

	"golang.org/x/term"
)

// ResolvePreset returns the preset to use. If flagValue is non-empty it is parsed (error if invalid).
// Otherwise, if stdin is a TTY, the user is prompted; if not a TTY, PresetStandard is used without prompting.
func ResolvePreset(flagValue string, stderr io.Writer) (project.Preset, error) {
	if flagValue != "" {
		p, err := project.ParsePreset(flagValue)
		if err != nil {
			return project.PresetStandard, err
		}
		return p, nil
	}
	if !isTTY(os.Stdin) {
		return project.PresetStandard, nil
	}
	return promptPreset(stderr, os.Stdin), nil
}

func isTTY(f *os.File) bool {
	return term.IsTerminal(int(f.Fd()))
}

func promptPreset(stderr io.Writer, stdin *os.File) project.Preset {
	fmt.Fprintln(stderr, "")
	fmt.Fprintln(stderr, "  Project structure")
	fmt.Fprintln(stderr, "  ----------------")
	fmt.Fprintln(stderr, "    1  minimal   Core only (router, architecture, interfaces, rules, tasks). No scripts.")
	fmt.Fprintln(stderr, "    2  standard  Core + scripts/ with manifest and example  (default)")
	fmt.Fprintln(stderr, "    3  full      Standard + actio/plugins/ and mcp/plugins/")
	fmt.Fprintln(stderr, "")
	fmt.Fprint(stderr, "  Choice [2]: ")

	sc := bufio.NewScanner(stdin)
	if !sc.Scan() {
		return project.PresetStandard
	}
	line := strings.TrimSpace(sc.Text())
	switch line {
	case "1":
		return project.PresetMinimal
	case "3":
		return project.PresetFull
	default:
		return project.PresetStandard
	}
}
