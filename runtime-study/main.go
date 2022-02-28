package main

import "runtime"

var GOOS_LIST = []string{"aix", "android", "darwin", "freebsd", "illumos", "ios", "js", "linux", "netbsd", "openbsd", "plan9", "solaris", "windows"}
var GOARCH_LIST = []string{"ppc64", "386", "amd64", "arm", "arm64", "wasm", "mips", "mips64", "mips64le", "ppc64le", "riscv64", "s390x"}

type OsInfo = struct {
	IsAix     bool
	IsAndroid bool
	IsDarwind bool
	IsMac     bool
	IsFreeBsd bool
	IsIllumos bool
	IsIos     bool
	IsJs      bool
	IsLinux   bool
	Isnetbsd  bool
	IsOpenbsd bool
	IsPlan9   bool
	IsSolaris bool
	IsWindows bool
}

type ArchInfo = struct {
	IsPpc64    bool
	Is386      bool
	IsAmd64    bool
	IsArm      bool
	IsArm64    bool
	IsWasm     bool
	IsMips     bool
	IsMips64   bool
	IsMips64le bool
	IsPpc64le  bool
	IsRiscv64  bool
	IsS390x    bool
}

type Info struct {
	OsInfo   OsInfo
	ArchInfo ArchInfo
}

// GetInfo() return runtime information
// ex: what target's os, architechur
func GetInfo() Info {
	goos := runtime.GOOS
	goarch := runtime.GOARCH

	return Info{
		OsInfo: OsInfo{
			IsAix:     goos == "aix",
			IsAndroid: goos == "android",
			IsDarwind: goos == "darwin",
			IsMac:     goos == "darwin",
			IsFreeBsd: goos == "freebsd",
			IsIllumos: goos == "illumos",
			IsIos:     goos == "ios",
			IsJs:      goos == "js",
			IsLinux:   goos == "linux",
			Isnetbsd:  goos == "netbsd",
			IsOpenbsd: goos == "openbsd",
			IsPlan9:   goos == "plan9",
			IsSolaris: goos == "solaris",
			IsWindows: goos == "windows",
		},
		ArchInfo: ArchInfo{
			IsPpc64:    goarch == "ppc64",
			Is386:      goarch == "386",
			IsAmd64:    goarch == "amd64",
			IsArm:      goarch == "arm",
			IsArm64:    goarch == "arm64",
			IsWasm:     goarch == "wasm",
			IsMips:     goarch == "mips",
			IsMips64:   goarch == "mips64",
			IsMips64le: goarch == "mips64le",
			IsPpc64le:  goarch == "ppc64le",
			IsRiscv64:  goarch == "riscv64",
			IsS390x:    goarch == "s390x",
		},
	}
}
