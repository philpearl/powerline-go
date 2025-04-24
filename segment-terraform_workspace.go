package main

import (
	"os"
)

const wsFile = "./.terraform/environment"

func segmentTerraformWorkspace(p *powerline) {
	stat, err := os.Stat(wsFile)

	if err == nil && !stat.IsDir() {
		workspace, err := os.ReadFile(wsFile)
		if err == nil {
			p.appendSegment("terraform-workspace", segment{
				content:    string(workspace),
				foreground: p.theme.TFWsFg,
				background: p.theme.TFWsBg,
			})

		}
	}

}
