package usecase

import (
	"os"

	"file/internal/4_domain/domain"
	"file/pkg/shared"
)

var myErr *shared.MyErr

func init() {
	myErr = shared.NewMyErr("file", "usecase")
}

type (
	// UseCase ...
	UseCase struct {
		ToService ToService
		ToDomain  ToDomain
	}
)

// ToService ...
type ToService interface {
	// grpc
	SvSendContent(address string, cc *shared.CommonContent) (string, error)

	// file
	SvExecCommand(value []string) ([]byte, error)
	SvUploadFile(filePath string, chunks *[]byte) error
	SvWriteJSONFile(filePath string, data []byte, perm os.FileMode) error
	SvReadJSONFile(filePath string) ([]byte, error)
}

// ToDomain ...
type ToDomain interface {
	// 標準関数
	StrconvFormatBool(value bool) string
	JSONUnmarshal(raw []byte, files *[]domain.File) error
	JSONMarshal(files *[]domain.File) ([]byte, error)
	JSONMarshalIndent(files *[]domain.File) ([]byte, error)
	FilepathJoin(strings ...string) string

	// 標準関数以外
	IsExistVid(files *[]domain.File, vid string) bool
	GetValidFiles(files *[]domain.File) string
	IsExistFilePath(filePath string) bool
	ParseMessage(message string) (string, error)
	AddtoFiles(files []domain.File, vid, name string) []domain.File
	ToggleFiles(files []domain.File, vid string) []domain.File
}
