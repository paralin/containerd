package containerd

import "fmt"

const VersionMajor = 0
const VersionMinor = 2
const VersionPatch = 2

var Version = fmt.Sprintf("%d.%d.%d", VersionMajor, VersionMinor, VersionPatch)

var GitCommit = ""
