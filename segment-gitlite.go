package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func segmentGitLite(p *powerline) {
	if len(p.ignoreRepos) > 0 {
		out, err := runGitCommand("git", "rev-parse", "--show-toplevel")
		if err != nil {
			return
		}
		out = strings.TrimSpace(out)
		if p.ignoreRepos[out] {
			return
		}
	}

	out, err := runGitCommand("git", "rev-parse", "--abbrev-ref", "HEAD")
	if err != nil {
		return
	}

	status := strings.TrimSpace(out)
	var branch string

	if status != "HEAD" {
		branch = status
	} else {
		branch = getGitDetachedBranch(p)
	}

	p.appendSegment("git-branch", segment{
		content:    branch,
		foreground: p.theme.RepoCleanFg,
		background: p.theme.RepoCleanBg,
	})
}

func gitProcessEnv() []string {
	home, _ := os.LookupEnv("HOME")
	path, _ := os.LookupEnv("PATH")
	env := map[string]string{
		"LANG": "C",
		"HOME": home,
		"PATH": path,
	}
	result := make([]string, 0)
	for key, value := range env {
		result = append(result, fmt.Sprintf("%s=%s", key, value))
	}
	return result
}

func runGitCommand(cmd string, args ...string) (string, error) {
	command := exec.Command(cmd, args...)
	command.Env = gitProcessEnv()
	out, err := command.Output()
	return string(out), err
}

func getGitDetachedBranch(p *powerline) string {
	out, err := runGitCommand("git", "rev-parse", "--short", "HEAD")
	if err != nil {
		out, err := runGitCommand("git", "symbolic-ref", "--short", "HEAD")
		if err != nil {
			return "Error"
		} else {
			a, _, _ := strings.Cut(out, "\n")
			return a
		}
	} else {
		detachedRef, _, _ := strings.Cut(out, "\n")
		return fmt.Sprintf("%s %s", p.symbolTemplates.RepoDetached, detachedRef)
	}
}
