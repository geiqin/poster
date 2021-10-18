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
	"github.com/geiqin/poster/core"
	"image"
)

// ImageRemoteHandler 根据URL地址设置图片
type ImageRemoteHandler struct {
	// 合成复用Next
	Next
	X      int
	Y      int
	Weight int
	Height int
	URL    string //http://xxx.png
}

// Do 地址逻辑
func (h *ImageRemoteHandler) Do(c *Context) (err error) {
	srcReader, err := core.GetResourceReader(h.URL)
	if err != nil {
		fmt.Errorf("core.GetResourceReader err：%v", err)
	}
	srcImage, imageType, err := image.Decode(srcReader)
	_ = imageType
	if err != nil {
		fmt.Errorf("SetRemoteImage image.Decode err：%v", err)
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
