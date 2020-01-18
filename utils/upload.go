package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelzx/paladin-go/errs"
	"github.com/michaelzx/paladin-go/logger"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

// 文件上传
func UploadToWeb(c *gin.Context, formFileName, webPath, staticDir string, allowTypes ...FileType) (fileWebPathList []string) {
	files := uploadGetFiles(c, formFileName)
	fileWebPathList = make([]string, 0, len(files))
	if !strings.HasSuffix(webPath, "/") {
		webPath = webPath + "/"
	}
	serverPath := staticDir + webPath
	if !PathExists(serverPath) {
		err := os.MkdirAll(serverPath, os.ModePerm)
		if err != nil {
			panic(errs.NewBadRequestError(err.Error()))
		}
	}
	uploadCheckFileType(files, allowTypes...)
	for _, file := range files {
		fileName := getRandomFileName() + path.Ext(file.Filename)
		fileWebPath := webPath + fileName
		fileServerPath := serverPath + fileName
		fileWebPathList = append(fileWebPathList, fileWebPath)
		err := c.SaveUploadedFile(file, fileServerPath)
		if err != nil {
			panic(errs.NewBadRequestError(err.Error()))
		} else {
			logger.Debug("file in: ", fileServerPath)
		}
	}

	return
}

func UploadOneToServer(c *gin.Context, formFileName, fileServerPath string, allowTypes ...FileType) error {
	files := uploadGetFiles(c, formFileName)
	if len(files) == 0 {
		return errs.NewBadRequestError("请选择要上传的文件")
	}
	if len(files) != 1 {
		return errs.NewBadRequestError("仅支持单文件上传")
	}
	fileDir := path.Dir(fileServerPath)
	if !PathExists(fileDir) {
		err := os.MkdirAll(fileDir, os.ModePerm)
		if err != nil {
			return errs.NewBadRequestError(err.Error())
		}
	}
	uploadCheckFileType(files, allowTypes...)
	for _, file := range files {
		err := c.SaveUploadedFile(file, fileServerPath)
		if err != nil {
			return errs.NewBadRequestError(err.Error())
		} else {
			logger.Debug("file in: ", fileServerPath)
		}
	}
	return nil
}

func uploadGetFiles(c *gin.Context, formFileName string) []*multipart.FileHeader {
	form, err := c.MultipartForm()
	if err != nil {
		panic(errs.NewBadRequestError(err.Error()))
	}
	return form.File[formFileName]
}
func uploadCheckFileType(files []*multipart.FileHeader, allowTypes ...FileType) {
	for _, file := range files {
		f, err := file.Open()
		if err != nil {
			panic(errs.NewBadRequestError(err.Error()))
		}
		fSrc, err := ioutil.ReadAll(f)
		_ = f.Close()
		ft := GetFileType(fSrc[:10])
		allow := false
		for _, at := range allowTypes {
			if ft == string(at) {
				allow = true
				break
			}
		}
		if !allow {
			panic(errs.NewBadRequestError("不允许上传 " + ft + " 类型的文件"))
		}
	}
}
func getRandomFileName() string {
	return strconv.FormatInt(time.Now().UnixNano()/1000000, 10) + RandomNumStr(3)
}
