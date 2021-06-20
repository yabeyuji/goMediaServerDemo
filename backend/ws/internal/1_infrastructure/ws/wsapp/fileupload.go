package wsapp

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

// FileUpload ...
func (wd *WsApp) FileUpload(c echo.Context) error {
	// multipart.FileHeaderはHTTPで扱うファイル形式なので、
	// external以外で扱うのにbytes.Bufferに変換
	uploadFile, err := c.FormFile("file")
	if err != nil {
		myErr.Logging(err)
	}

	// ファイル名取得
	fileName := strings.Split(uploadFile.Filename, ".")[0]

	// multipart ファイルボディ取得
	multipartFileBody, err := uploadFile.Open()
	if err != nil {
		myErr.Logging(err)
	}
	defer multipartFileBody.Close()

	// multipart を bytes.Buffer に変換
	fileBody := bytes.NewBuffer(nil)
	_, err = io.Copy(fileBody, multipartFileBody)
	if err != nil {
		myErr.Logging(err)
	}

	wd.Controller.FileUpload(fileName, fileBody)

	return c.JSON(http.StatusOK, "ok")
}
