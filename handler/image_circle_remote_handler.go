/**
 * @Author: entere@126.com
 * @Description:
 * @File:  image_local_handler
 * @Version: 1.0.0
 * @Date: 2020/5/22 08:51
 */

package handler

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/geiqin/poster/circlemask"
	"github.com/geiqin/poster/core"
	"image"
)

// ImageCircleRemoteHandler 根据URL地址设置圆形图片
type ImageCircleRemoteHandler struct {
	// 合成复用Next
	Next
	X      int
	Y      int
	Weight int
	Height int
	URL    string //http://xxx.png
}

// Do 地址逻辑
func (h *ImageCircleRemoteHandler) Do(c *Context) (err error) {
	srcReader, err := core.GetResourceReader(h.URL)
	if err != nil {
		fmt.Errorf("core.GetResourceReader err：%v", err)
	}
	srcImage, imageType, err := image.Decode(srcReader)
	_ = imageType
	if err != nil {
		fmt.Errorf("SetRemoteImage image.Decode err：%v", err)
	}

	if h.Weight > 0 && h.Height > 0 {
		srcImage = imaging.Resize(srcImage, h.Weight, h.Height, imaging.Lanczos)
	}

	// 算出图片的宽度和高试
	width := srcImage.Bounds().Max.X - srcImage.Bounds().Min.X
	height := srcImage.Bounds().Max.Y - srcImage.Bounds().Min.Y

	//把头像转成Png,否则会有白底
	srcPng := core.NewPNG(0, 0, width, height)
	core.MergeImage(srcPng, srcImage, srcImage.Bounds().Min)

	// 圆的直径以长边为准
	diameter := width
	if width > height {
		diameter = height
	}
	// 遮罩
	srcMask := circlemask.NewCircleMask(srcPng, image.Point{0, 0}, diameter)

	srcPoint := image.Point{
		X: h.X,
		Y: h.Y,
	}
	core.MergeImage(c.PngCarrier, srcMask, srcImage.Bounds().Min.Sub(srcPoint))
	return
}
