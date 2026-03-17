package mcp

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// pluginConfig describes how to start a third-party MCP server over stdio.
// It is discovered from <repo-root>/mcp/plugins/*.{yaml,yml,json}.
type pluginConfig struct {
	Name        string            `json:"name" yaml:"name"`
	Description string            `json:"description" yaml:"description"`
	Command     string            `json:"command" yaml:"command"`
	Args        []string          `json:"args" yaml:"args"`
	Env         map[string]string `json:"env" yaml:"env"`
}

func loadPluginConfigs(repoRoot string) ([]pluginConfig, error) {
	pluginsDir := filepath.Join(repoRoot, "mcp", "plugins")
	if _, err := os.Stat(pluginsDir); err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("stat plugins dir: %w", err)
	}

	entries, err := os.ReadDir(pluginsDir)
	if err != nil {
		return nil, fmt.Errorf("read plugins dir: %w", err)
	}

	var out []pluginConfig
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(e.Name()))
		if ext != ".yaml" && ext != ".yml" && ext != ".json" {
			continue
		}
		path := filepath.Join(pluginsDir, e.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("read plugin config %s: %w", e.Name(), err)
		}

		var cfg pluginConfig
		switch ext {
		case ".json":
			if err := json.Unmarshal(data, &cfg); err != nil {
				return nil, fmt.Errorf("parse plugin config %s: %w", e.Name(), err)
			}
		default:
			if err := yaml.Unmarshal(data, &cfg); err != nil {
				return nil, fmt.Errorf("parse plugin config %s: %w", e.Name(), err)
			}
		}

		if cfg.Name == "" {
			cfg.Name = strings.TrimSuffix(e.Name(), filepath.Ext(e.Name()))
		}
		cfg.Name = strings.TrimSpace(cfg.Name)
		if cfg.Name == "" {
			return nil, fmt.Errorf("plugin config %s: name must be set", e.Name())
		}
		if cfg.Command == "" {
			return nil, fmt.Errorf("plugin %q: command must be set", cfg.Name)
		}

		// Expand environment variables in config values.
		cfg.Command = os.ExpandEnv(cfg.Command)
		for i := range cfg.Args {
			cfg.Args[i] = os.ExpandEnv(cfg.Args[i])
		}
		if cfg.Env != nil {
			for k, v := range cfg.Env {
				cfg.Env[k] = os.ExpandEnv(v)
			}
		}

		out = append(out, cfg)
	}

	return out, nil
}

