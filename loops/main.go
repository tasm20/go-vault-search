package loops

import "github.com/tasm20/go-vault-search/initVault"

type PathStruct struct {
	dirs  []string
	files []string
}

func (p PathStruct) GetDirs() []string {
	return p.dirs
}

func (p PathStruct) GetFiles() []string {
	return p.files
}

var (
	pathStruct  PathStruct
	clientVault = initVault.GetClient()
)
