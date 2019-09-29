package service

import (
	. "IxDServer/config"
	"IxDServer/param/qiniu"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"log"
)

//七牛的认证
func QiniuKey() map[string]string {
	putPolicy := storage.PutPolicy{
		Scope: CONF.QiniuBucket,
	}
	mac := qbox.NewMac(CONF.QiniuAccessKey, CONF.QiniuSecretKey)
	upToken := putPolicy.UploadToken(mac)
	return map[string]string{
		"upToken": upToken,
	}
}

func QiniuFileInfo(p *qiniu.FileInfo) (interface{}, error) {
	mac := qbox.NewMac(CONF.QiniuAccessKey, CONF.QiniuSecretKey)
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: false,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)
	bucket := CONF.QiniuBucket
	key := p.Etag
	fileInfo, err := bucketManager.Stat(bucket, key)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return fileInfo, nil
}
