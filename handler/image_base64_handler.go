/**
 * @Author: entere@126.com
 * @Description:
 * @File:  image_local_handler
 * @Version: 1.0.0
 * @Date: 2020/5/22 08:51
 */

package handler

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/geiqin/poster/core"
	"image"
	"image/jpeg"
	"image/png"
)

// ImageBase64Handler 根据Base64设置图片
type ImageBase64Handler struct {
	// 合成复用Next
	Next
	X       int
	Y       int
	Weight  int
	Height  int
	Content string //Base64
	IsJpeg  bool
}

// Do 地址逻辑
func (h *ImageBase64Handler) Do(c *Context) (err error) {
	var srcImage image.Image
	var imgErr error
	imgData, err := base64.StdEncoding.DecodeString(h.Content) //成图片文件并把文件写入到buffer
	if err != nil {
		fmt.Errorf("ImageBase64 image.Decode err：%v", err)
		return
	}

	bbb := bytes.NewBuffer(imgData)
	if h.IsJpeg {
		srcImage, imgErr = jpeg.Decode(bbb)
	} else {
		srcImage, imgErr = png.Decode(bbb)
	}

	if imgErr != nil {
		fmt.Errorf("img.Decode err：%v", imgErr)
		return
	}
	srcPoint := image.Point{
		X: h.X,
		Y: h.Y,
	}

	if h.Weight > 0 && h.Height > 0 {
		srcImage = imaging.Resize(srcImage, h.Weight, h.Height, imaging.Lanczos)
	}

	core.MergeImage(c.PngCarrier, srcImage, srcImage.Bounds().Min.Sub(srcPoint))
	return
}
