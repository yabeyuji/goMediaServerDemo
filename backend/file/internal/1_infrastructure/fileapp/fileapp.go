package fileapp

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"file/internal/2_adapter/service"
	"file/pkg/shared"
)

var myErr *shared.MyErr

func init() {
	myErr = shared.NewMyErr("file", "infrastructure:file")
}

type file struct{}

// NewToFileApp ...
func NewToFileApp() service.ToFile {
	fl := new(file)
	return fl
}

// IsExecCommand ...
func (fl *file) IsExecCommand(command []string) ([]byte, error) {
	res, err := exec.Command(command[0], command[1:]...).Output() // #nosec G204
	if err != nil {
		myErr.Logging(err, command)
		return nil, err
	}

	return res, nil
}

// IsUploadFile ...
func (fl *file) IsUploadFile(filePath string, chunks *[]byte) error {
	file, err := os.Create(filePath)
	if err != nil {
		myErr.Logging(err, filePath)
		return err
	}

	defer func() {
		if err := file.Close(); err != nil {
			myErr.Logging(err)
		}
	}()

	_, err = file.Write(*chunks)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	return nil
}

// IsWriteJSONFile ...
func (fl *file) IsWriteJSONFile(filePath string, data []byte, perm os.FileMode) error {
	err := ioutil.WriteFile(filePath, []byte(data), perm)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	return nil
}

// IsReadJSONFile ...
func (fl *file) IsReadJSONFile(filePath string) ([]byte, error) {
	buf, err := ioutil.ReadFile(filepath.Clean(filePath))
	if err != nil {
		myErr.Logging(err)
		return nil, err
	}

	return buf, nil
}
