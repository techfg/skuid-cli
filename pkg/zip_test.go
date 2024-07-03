package pkg_test

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/skuid/skuid-cli/pkg"
	"github.com/skuid/skuid-cli/pkg/logging"
	"github.com/skuid/skuid-cli/pkg/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockZipWriter struct {
	mock.Mock
}

func (m *MockZipWriter) Create(name string) (io.Writer, error) {
	return nil, fmt.Errorf("force zip create error")
}

func (m *MockZipWriter) Close() error {
	return nil
}

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

func TestArchiveWithFilterFuncAbandonsChannel(t *testing.T) {
	createZipWriter := func(w io.Writer) pkg.ZipWriter {
		m := new(MockZipWriter)
		return m
	}

	_, err := pkg.ArchiveWithFilterFunc(".", createZipWriter, func(relativePath string) bool { return true })
	logging.Get().Infof("**** ArchiveWithFilterFunc returned with error ****: %v", err)
	// force a sleep to demonstrate that channel is never closed
	time.Sleep(10 * time.Second)
	assert.NotNil(t, err)
	logging.Get().Info("**** Test completed ****")
}
