package proFile

import (
	"github.com/pkg/profile"
)

const pprofPath = "/tmp"

func Start(mode string) (pPROF interface{ Stop() }) {

	switch mode {
	case "cpu":
		pPROF = profile.Start(profile.CPUProfile, profile.ProfilePath(pprofPath))
	case "mem":
		pPROF = profile.Start(profile.MemProfile, profile.ProfilePath(pprofPath))
	case "mutex":
		pPROF = profile.Start(profile.MutexProfile, profile.ProfilePath(pprofPath))
	case "block":
		pPROF = profile.Start(profile.BlockProfile, profile.ProfilePath(pprofPath))
	}

	return
}
