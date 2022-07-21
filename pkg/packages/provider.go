package packages

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// Describe package managers.

/// Set of arguments for package manager.
type Provider struct {
	bin           string // Binary name.
	info          string // Show info about package.
	install       string // Install package.
	yes           string // Answer YES to manager.
	remove        string // Remove package.
	upgrade       string // Upgrade package.
	upgradeAll    string // Upgrade all packages.
	search        string // Search package.
	updateIndex   string // Update index DB.
	listUpdates   string // List all upgradable packages.
	listInstalled string // List all installed packages.
	provides      string // Find which package provides resource.
}

var Providers map[string]Provider

func init() {
	p := make(map[string]Provider)

	// yum.
	p["yum"] = Provider{
		bin:           "yum",
		info:          "info",
		install:       "install",
		yes:           "-y",
		remove:        "remove",
		upgrade:       "update",
		upgradeAll:    "update",
		search:        "search",
		updateIndex:   "makecache",
		listUpdates:   "list updates",
		listInstalled: "list installed",
		provides:      "provides",
	}

	// For cut&paste.
	p["DEMO"] = Provider{
		bin:           "",
		info:          "",
		install:       "",
		yes:           "",
		remove:        "",
		upgrade:       "",
		upgradeAll:    "",
		search:        "",
		updateIndex:   "",
		listUpdates:   "",
		listInstalled: "",
		provides:      "",
	}

	// brew.
	p["brew"] = Provider{
		bin:           "brew",
		info:          "info",
		install:       "install",
		yes:           "", // Not provided.
		remove:        "uninstall",
		upgrade:       "upgrade",
		upgradeAll:    "upgrade",
		search:        "search",
		updateIndex:   "update",
		listUpdates:   "outdated",
		listInstalled: "list",
		provides:      "", // Not provided.
	}

	Providers = p
}

func GetProvider() (*Provider, error) {
	pname, err := PMName()
	cobra.CheckErr(err)
	p, ok := Providers[pname]
	if !ok {
		return nil, fmt.Errorf("Provider %s does not supported", pname)
	}
	return &p, nil
}

// Provider methods.

func (p *Provider) Info(args ...string) {
	args = append([]string{p.info}, args...)
	cmd := exec.Command(p.bin, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
