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

// ImageHandler 根据URL地址设置图片
type ImageHandler struct {
	// 合成复用Next
	Next
	Reader
	X      int
	Y      int
	Weight int
	Height int
	Path   string //本地路径
	URL    string //网络地址
}

// Do 地址逻辑
func (h *ImageHandler) Do(c *Context) (err error) {
	var srcImage image.Image
	if h.Path != "" {
		srcImage, err = h.GetLocalImage(h.Path)
	} else if h.URL != "" {
		srcImage, err = h.GetRemoteImage(h.URL)
	}

	if err != nil {
		fmt.Errorf("core.GetResourceReader err：%v", err)
		return err
	}

	srcPoint := image.Point{
		X: h.X,
		Y: h.Y,
	}

	if h.Weight > 0 && h.Height > 0 {
		srcImage = imaging.Resize(srcImage, h.Weight, h.Height, imaging.Lanczos)
	}

	core.MergeImage(c.PngCarrier, srcImage, srcImage.Bounds().Min.Sub(srcPoint))
	return nil
}
