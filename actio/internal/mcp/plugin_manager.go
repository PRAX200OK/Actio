package mcp

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"sync"
)

const pluginURIPrefix = "plugin://"

type pluginClientFactory func(ctx context.Context, cfg pluginConfig) (pluginClient, error)

type pluginManager struct {
	repoRoot string
	stderr   io.Writer
	factory  pluginClientFactory

	mu      sync.Mutex
	configs map[string]pluginConfig
	clients map[string]pluginClient
}

func newPluginManager(repoRoot string, stderr io.Writer, factory pluginClientFactory) *pluginManager {
	if factory == nil {
		factory = func(ctx context.Context, cfg pluginConfig) (pluginClient, error) {
			return newStdioPluginClient(ctx, cfg)
		}
	}
	return &pluginManager{
		repoRoot: repoRoot,
		stderr:   stderr,
		factory:  factory,
		configs:  make(map[string]pluginConfig),
		clients:  make(map[string]pluginClient),
	}
}

func (pm *pluginManager) loadConfigsLocked() error {
	cfgs, err := loadPluginConfigs(pm.repoRoot)
	if err != nil {
		return err
	}
	pm.configs = make(map[string]pluginConfig, len(cfgs))
	for _, c := range cfgs {
		pm.configs[c.Name] = c
	}
	return nil
}

func (pm *pluginManager) getClient(ctx context.Context, name string) (pluginClient, error) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	if len(pm.configs) == 0 {
		if err := pm.loadConfigsLocked(); err != nil {
			return nil, err
		}
	}
	if c := pm.clients[name]; c != nil {
		return c, nil
	}
	cfg, ok := pm.configs[name]
	if !ok {
		return nil, fmt.Errorf("unknown plugin: %s", name)
	}
	c, err := pm.factory(ctx, cfg)
	if err != nil {
		return nil, err
	}
	pm.clients[name] = c
	return c, nil
}

func (pm *pluginManager) listAllResources(ctx context.Context) ([]map[string]any, error) {
	pm.mu.Lock()
	if err := pm.loadConfigsLocked(); err != nil {
		pm.mu.Unlock()
		return nil, err
	}
	var names []string
	for n := range pm.configs {
		names = append(names, n)
	}
	pm.mu.Unlock()

	var out []map[string]any
	for _, name := range names {
		c, err := pm.getClient(ctx, name)
		if err != nil {
			continue
		}
		res, err := c.ListResources(ctx)
		if err != nil {
			continue
		}
		for _, r := range res {
			origURI, _ := r["uri"].(string)
			if origURI == "" {
				continue
			}
			proxyURI := makePluginURI(name, origURI)
			desc, _ := r["description"].(string)
			out = append(out, map[string]any{
				"uri":         proxyURI,
				"description": fmt.Sprintf("Plugin %q: %s", name, desc),
			})
		}
	}
	return out, nil
}

func makePluginURI(pluginName, origURI string) string {
	v := url.Values{}
	v.Set("uri", origURI)
	return pluginURIPrefix + pluginName + "?" + v.Encode()
}

func parsePluginURI(uri string) (pluginName string, origURI string, ok bool) {
	if len(uri) <= len(pluginURIPrefix) || uri[:len(pluginURIPrefix)] != pluginURIPrefix {
		return "", "", false
	}
	u, err := url.Parse(uri)
	if err != nil {
		return "", "", false
	}
	name := u.Host
	if name == "" {
		return "", "", false
	}
	orig := u.Query().Get("uri")
	if orig == "" {
		return "", "", false
	}
	return name, orig, true
}

