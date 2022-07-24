package goiac

import (
	"os"

	"github.com/Tolyar/goiac/pkg/sysinfo"
	"github.com/spf13/cobra"
)

var F Facts

type Facts struct {
	SysInfo SysInfo
}

type SysInfo struct {
	Platform string
	Arch     string
	Linux    map[string]string
	Os       string
	Hostname string
}

func init() {
	var ok bool

	F = Facts{}
	s := SysInfo{}
	h, err := os.Hostname()
	cobra.CheckErr(err)
	s.Hostname = h

	s.Platform = sysinfo.Platform()
	s.Arch = sysinfo.Arch()
	if s.Arch == "linux" {
		s.Linux, err = sysinfo.LinuxRelease()
		cobra.CheckErr(err)
		if s.Os, ok = s.Linux["ID"]; !ok {
			s.Os = s.Platform
		}
	} else {
		s.Os = s.Platform
	}
	F.SysInfo = s
}
