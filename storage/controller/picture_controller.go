package controller

import (
	"github.com/kataras/iris"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"stouch_server/common/er"
	"stouch_server/common/utils"
	"stouch_server/conf"
	"stouch_server/storage/model"
	"stouch_server/storage/service"
	"strings"
)

type PictureController struct{
	Ctx iris.Context
}

func (c *PictureController) GetBy(id int64) interface{}{
	picture := model.Picture{Id: id}
	if ok, _ := conf.Orm.Get(&picture); ok {
		return er.Data(map[string]model.Picture{"picture": picture})
	} else  {
		return er.SourceNotExistError
	}
}

func (c *PictureController) Post() interface{}{
	file, fileHeader, err := c.Ctx.FormFile("file")
	if err != nil {
		conf.Logger.Error("file is no exist!!!")
	}
	img, _, err := image.DecodeConfig(file)
	width, height, size := img.Width, img.Height, fileHeader.Size
	file.Seek(0, 0)
	md5 := utils.GetMD5(file)
	file.Seek(0, 0)
	sr := strings.Split(fileHeader.Filename, ".")
	picture := &model.Picture{Width: width, Height: height, Size:size, Md5:md5, Format: sr[len(sr) - 1]}
	if service.GetOrSave(md5 + "." + string(sr[len(sr) - 1]), file){
		if _, err := conf.Orm.Get(picture); err!= nil {
			conf.Logger.Error(err)
		}
	} else {
		if _, err := conf.Orm.Insert(picture); err!= nil {
			conf.Logger.Error(err)
		}
	}
	return er.Data(map[string]model.Picture{"picture": *picture})
}
