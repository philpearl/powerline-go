package main

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"fmt"

	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v2"
)

type KubeContext struct {
	Context struct {
		Cluster   string
		Namespace string
		User      string
	}
	Name string
}

type KubeConfig struct {
	Contexts       []KubeContext `yaml:"contexts"`
	CurrentContext string        `yaml:"current-context"`
}

func homePath() string {
	env := "HOME"
	if runtime.GOOS == "windows" {
		env = "USERPROFILE"
	}
	return os.Getenv(env)
}

func readKubeConfig(config *KubeConfig, path string) (err error) {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return
	}
	fileContent, err := os.ReadFile(absolutePath)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(fileContent, config)
	if err != nil {
		return
	}

	return
}

func readGcloudAccount(basePath string) (string, error) {
	data, err := os.ReadFile(filepath.Join(basePath, "active_config"))
	if err != nil {
		return "", err
	}
	filename := filepath.Join(basePath, "configurations", "config_"+strings.TrimSpace(string(data)))

	cfg, err := ini.Load(filename)
	if err != nil {
		return "", err
	}
	account := cfg.Section("core").Key("account").String()
	return account, nil
}

func segmentKube(p *powerline) {
	paths := append(strings.Split(os.Getenv("KUBECONFIG"), ":"), filepath.Join(homePath(), ".kube", "config"))
	config := &KubeConfig{}
	for _, configPath := range paths {
		if readKubeConfig(config, configPath) == nil {
			break
		}
	}

	// We also read the gcloud config to determine the current account. In gke the kubernetes account is determined
	// via gcloud. We set up service accounts with a sudo- prefix with additional permissions
	gcloudBase := filepath.Join(homePath(), ".config", "gcloud")
	account, err := readGcloudAccount(gcloudBase)
	if err != nil {
		fmt.Println(err)
	}
	sudo := strings.HasPrefix(account, "sudo-")

	name := config.CurrentContext
	namespace := ""
	for _, context := range config.Contexts {
		if context.Name == config.CurrentContext {
			namespace = context.Context.Namespace
			break
		}
	}

	// Only draw the icon once
	kubeIconHasBeenDrawnYet := false
	if name != "" {
		kubeIconHasBeenDrawnYet = true
		fg, bg := p.theme.KubeClusterFg, p.theme.KubeClusterBg
		if sudo {
			fg = 9  // very red
			bg = 51 // cyan
			name = name + "-sudo"
		}
		p.appendSegment("kube-cluster", segment{
			content:    fmt.Sprintf("⎈ %s", name),
			foreground: fg,
			background: bg,
		})
	}

	if namespace != "" {
		content := namespace
		if !kubeIconHasBeenDrawnYet {
			content = fmt.Sprintf("⎈ %s", content)
		}
		p.appendSegment("kube-namespace", segment{
			content:    content,
			foreground: p.theme.KubeNamespaceFg,
			background: p.theme.KubeNamespaceBg,
		})
	}
}
