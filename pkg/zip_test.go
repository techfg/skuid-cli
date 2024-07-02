package pkg_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/skuid/skuid-cli/pkg"
	"github.com/skuid/skuid-cli/pkg/logging"
	"github.com/skuid/skuid-cli/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestZip(t *testing.T) {
	util.SkipIntegrationTest(t)
	cd, _ := os.Getwd()
	relpath := filepath.Join(cd, ".", ".", "_deploy")
	bb, err := pkg.Archive(relpath, nil)
	if err != nil {
		logging.Get().Fatal(err)
	}
	logging.Get().Info(len(bb))

}

func TestArchiveWithFilterFunc(t *testing.T) {
	mockDirectoryWalker := func(root string, visit filepath.WalkFunc) (err error) {
		return fmt.Errorf("Force a failure")
	}

	_, err := pkg.ArchiveWithFilterFunc(".", mockDirectoryWalker, func(relativePath string) bool { return true })

	assert.NotNil(t, err)
}
