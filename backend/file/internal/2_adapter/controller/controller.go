package controller

import (
	"file/internal/2_adapter/service"
	"file/internal/3_usecase/usecase"
	"file/internal/4_domain/domain"
	"file/pkg/shared"
)

var myErr *shared.MyErr

func init() {
	myErr = shared.NewMyErr("file", "adapter:controller")
}

type (
	// Controller ...
	Controller struct {
		UseCase usecase.UseCase
		Files   []domain.File
	}
)

// NewController ...
func NewController(toGrpcOut service.ToGrpcOut, toFile service.ToFile) *Controller {
	ct := &Controller{
		UseCase: usecase.UseCase{
			ToDomain: domain.NewDomain(),
			ToService: &service.Service{
				ToGrpcOut: toGrpcOut,
				ToFile:    toFile,
			},
		},
	}

	return ct
}

// BootStrap ...
func (ctrl *Controller) BootStrap() {
	files, err := ctrl.UseCase.ReadFilesFromJSON(shared.JSONPath)
	if err != nil {
		myErr.Logging(err, shared.JSONPath)
	}

	ctrl.Files = *files
}

// UploadFile ...
func (ctrl *Controller) UploadFile(name, vid string, chunks *[]byte) error {
	// ファイル保存
	err := ctrl.UseCase.UploadFile(vid, chunks)
	if err != nil {
		myErr.Logging(err, vid, chunks)
		return err
	}

	// ファイル変換
	err = ctrl.UseCase.Convert(vid)
	if err != nil {
		myErr.Logging(err, vid)
		return err
	}

	// ファイル移動
	err = ctrl.UseCase.Move(vid)
	if err != nil {
		myErr.Logging(err, vid)
		return err
	}

	// ファイルが存在するか確認
	isExistFiles := ctrl.UseCase.IsExistFiles(vid)
	if !isExistFiles {
		myErr.Logging(err, vid)
		return err
	}

	// ファイル情報を更新
	ctrl.Files = ctrl.UseCase.AddtoFiles(ctrl.Files, vid, name)

	// 更新したファイル情報をjson化
	filesString, err := ctrl.UseCase.GetFilesString(&ctrl.Files)
	if err != nil {
		myErr.Logging(err, &ctrl.Files)
		return err
	}

	// json情報をブラウザに送信
	_, err = ctrl.UseCase.SendFilesToWs(filesString)
	if err != nil {
		myErr.Logging(err, filesString)
		return err
	}

	// ファイル情報を書き込み
	err = ctrl.UseCase.WriteFilesIntoJSON(&ctrl.Files)
	if err != nil {
		myErr.Logging(err, &ctrl.Files)
		return err
	}

	return nil
}

// ToggleValid ...
func (ctrl *Controller) ToggleValid(cc *shared.CommonContent) error {
	ctrl.Files = ctrl.UseCase.ToggleFiles(ctrl.Files, cc.Value)

	// 更新したファイル情報をjson化
	filesString, err := ctrl.UseCase.GetFilesString(&ctrl.Files)
	if err != nil {
		myErr.Logging(err, &ctrl.Files)
		return err
	}

	// json情報をブラウザに送信
	_, err = ctrl.UseCase.SendFilesToWs(filesString)
	if err != nil {
		myErr.Logging(err, filesString)
		return err
	}

	// // 更新したファイルリストをvlcに送信
	validFiles := ctrl.UseCase.GetValidFiles(&ctrl.Files)
	_, err = ctrl.UseCase.SendFilesToMedia(validFiles)
	if err != nil {
		myErr.Logging(err, validFiles)
		return err
	}

	// ファイル情報を書き込み
	err = ctrl.UseCase.WriteFilesIntoJSON(&ctrl.Files)
	if err != nil {
		myErr.Logging(err, &ctrl.Files)
		return err
	}

	return nil
}

// FileOperation ...
func (ctrl *Controller) FileOperation(message string) (string, error) {
	var value string

	switch message {
	case shared.TargetGetValidFiles:
		value = ctrl.UseCase.GetValidFiles(&ctrl.Files)

	case shared.TargetGetFiles:
		filesString, err := ctrl.UseCase.GetFilesString(&ctrl.Files)
		if err != nil {
			myErr.Logging(err, &ctrl.Files)
			return "", err
		}
		value = filesString
	}

	return value, nil
}
