/**
 * @Author: entere@126.com
 * @Description:
 * @File:  background_handler.go
 * @Version: 1.0.0
 * @Date: 2020/5/21 12:31
 */

package handler

import (
	"bytes"
	"fmt"
	"github.com/geiqin/poster/core"
	"image/png"
)

// EndHandler 结束，写在最后，把图片合并到一张图上
type EndHandler struct {
	// 合成复用Next
	Next
	Output string // "/tmp/xxx.png"
	IsWriteBuffer bool
	BufferData []byte
}

// Do 地址逻辑
func (h *EndHandler) Do(c *Context) (err error) {
	// 新建文件载体
	//fileName := "poster-" + xid.New().String() + ".png"
	if h.IsWriteBuffer{
		// 创建一个空的字节切片
		// 创建一个缓冲区，并将图像写入其中
		buffer := new(bytes.Buffer)
		bufErr := png.Encode(buffer, c.PngCarrier)
		if bufErr != nil {
			// 处理错误
		}
		// 将缓冲区中的内容复制到字节切片中
		h.BufferData = buffer.Bytes()
	}else{
		merged, err := core.NewMerged(h.Output)
		if err != nil {
			fmt.Errorf("core.NewMerged err：%v", err)
		}
		// 合并
		err = core.Merge(c.PngCarrier, merged)
		if err != nil {
			fmt.Errorf("core.Merge err：%v", err)
		}
	}

	return
}
