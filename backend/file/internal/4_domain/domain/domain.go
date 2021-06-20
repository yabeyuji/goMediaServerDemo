package domain

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"file/pkg/shared"
)

var myErr *shared.MyErr

func init() {
	myErr = shared.NewMyErr("file", "domain")
}

type (
	// File ...
	File struct {
		Vid   string
		Name  string
		Valid bool
	}
)

type (
	domain struct{}
)

// NewDomain ...
func NewDomain() *domain {
	return &domain{}
}

// JSONUnmarshal ...
func (domain *domain) JSONUnmarshal(raw []byte, files *[]File) error {
	err := json.Unmarshal(raw, files)
	if err != nil {
		myErr.Logging(err, raw, files)
		return err
	}

	return nil
}

// JSONMarshal ...
func (domain *domain) JSONMarshal(files *[]File) ([]byte, error) {
	raw, err := json.Marshal(files)
	if err != nil {
		myErr.Logging(err, raw, files)
		return nil, err
	}

	return raw, nil
}

// JSONMarshalIndent ...
func (domain *domain) JSONMarshalIndent(files *[]File) ([]byte, error) {
	raw, err := json.MarshalIndent(files, "", "    ")
	if err != nil {
		myErr.Logging(err, raw, files)
		return nil, err
	}

	return raw, nil
}

// IsExistVid ...
func (domain *domain) IsExistVid(files *[]File, vid string) bool {
	for _, file := range *files {
		if file.Vid == vid {
			return true
		}
	}
	return false
}

// GetValidFiles ...
func (domain *domain) GetValidFiles(files *[]File) string {
	// https://knsh14.github.io/translations/go-codereview-comments/#declaring-empty-slices
	// jsonで[]をしたいので下記実装
	validFiles := []string{}
	for _, file := range *files {
		if file.Valid {
			validFiles = append(validFiles, file.Vid)
		}
	}
	return strings.Join(validFiles, ",")
}

// IsExistFilePath ...
func (domain *domain) IsExistFilePath(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}

	return false
}

func (domain *domain) ParseMessage(message string) (string, error) {
	var err error

	switch message {
	case shared.TargetGetValidFiles:
	case shared.TargetGetFiles:
	default:
		err = errors.New("not found target")
	}

	if err != nil {
		myErr.Logging(err, message)
		return "", err
	}

	return message, nil
}

// AddtoFiles ...
func (domain *domain) AddtoFiles(files []File, vid, name string) []File {
	file := &File{
		Vid:   vid,
		Name:  name,
		Valid: false,
	}

	files = append([]File{*file}, files[0:]...)
	return files
}

// StrconvFormatBool ...
func (domain *domain) StrconvFormatBool(value bool) string {
	return strconv.FormatBool(value)
}

// ToggleFiles ...
func (domain *domain) ToggleFiles(files []File, vid string) []File {
	for index, file := range files {
		if file.Vid == vid {
			files[index].Valid = !files[index].Valid
		}
	}
	return files
}

// FilepathJoin ...
func (domain *domain) FilepathJoin(strings ...string) string {
	return filepath.Join(strings...)
}
