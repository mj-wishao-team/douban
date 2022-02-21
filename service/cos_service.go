package service

import (
	"context"
	"douban/tool"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
	"net/url"
	"time"
)

//cos
//上传头像
func UploadAvatar(file io.Reader, filePath string) error {
	cfg := tool.GetCfg().Cos
	u, _ := url.Parse(cfg.AvatarUrl)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		// 设置超时时间
		Timeout: 30 * time.Second,
		Transport: &cos.AuthorizationTransport{
			SecretID:  cfg.SecretId,
			SecretKey: cfg.SecretKey,
		},
	})

	_, err := c.Object.Put(context.Background(), filePath, file, nil)
	if err != nil {
		return err
	}
	return nil
}
