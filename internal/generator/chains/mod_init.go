package chains

import (
	"errors"
	"github.com/spf13/afero"
	entity "goro/internal/config"
	"os/exec"
)

type modInitChain struct{}

func NewModInitChain() *modInitChain {
	return &modInitChain{}
}

func (m *modInitChain) Apply(fs *afero.Afero, data entity.Config) (*afero.Afero, error) {
	cmd := exec.Command("go", "mod", "init", data.App.Name)
	cmd.Dir = data.App.WorkDir

	output, err := cmd.CombinedOutput()

	out := string(output)
	if err != nil && out != "" {
		return fs, errors.New(out)
	}

	return fs, nil
}

func (m *modInitChain) Name() string {
	return "Go mod init"
}

func (m *modInitChain) Rollback() error {
	return nil
}
