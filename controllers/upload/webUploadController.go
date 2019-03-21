package upload

import (
	"fmt"
	"auroraLZDF/beegoBBS/utils"
	"os"
)

type WebUploadController struct {
	Controller
}

func (this *WebUploadController) Upload () {
	file, fHeader, err := this.GetFile("file")
	defer file.Close()

	if err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	dir := uploadDir + utils.Date("2006/01/02" )
	if err := os.MkdirAll(dir, 0775);err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	path := dir + "/" +fHeader.Filename

	fmt.Println(file, fHeader.Filename,fHeader.Size)	//
	if err := this.SaveToFile("file", path); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	data["file_path"] = domain + "/" + path
	this.JsonMessage(1, "图片上传成功！", data)
}