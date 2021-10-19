package handler

import (
	"fmt"
	"github.com/geiqin/poster/core"
	"image"
	"os"
)

type Reader struct {
}

func (h *Reader) GetRemoteImage(url string) (image.Image, error) {
	srcReader, err := core.GetResourceReader(url)
	if err != nil {
		fmt.Errorf("core.GetResourceReader err：%v", err)
		return nil, err
	}
	srcImage, imageType, err := image.Decode(srcReader)
	_ = imageType
	if err != nil {
		fmt.Errorf("SetRemoteImage image.Decode err：%v", err)
		return nil, err
	}
	return srcImage, err
}

func (h *Reader) GetLocalImage(path string) (image.Image, error) {
	//获取背景 必须是PNG图
	imageFile, err := os.Open(path)

	if err != nil {
		fmt.Errorf("os.Open err：%v", err)
		return nil, err
	}

	srcImage, imageType, srcErr := image.Decode(imageFile)
	_ = imageType

	if srcErr != nil {
		fmt.Errorf("png.Decode err：%v", err)
	}
	return srcImage, srcErr
}
