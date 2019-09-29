package config

import "flag"

type config struct {
	HttpAddr          string
	EncryptedKey      string
	DataSourceName    string
	AuthTokenStart    string
	AuthTokenDeadline int64
	QiniuAccessKey    string
	QiniuSecretKey    string
	QiniuBucket       string
	GaoDeKey          string
}

var CONF *config

func init() {
	httpAddr := flag.String("httpAddr", ":50001", "服务端口")
	encryptedKey := flag.String("encryptedKey", "key:zoolon872112", "加解密token的私钥")
	dataSourceName := flag.String("dataSourceName", "root:123.com@tcp(quickex.com.cn:3306)/ixd?charset=utf8", "mysql数据源")
	authTokenStart := flag.String("authTokenStart", "Bearer ", "token类型")
	authTokenDeadline := flag.Int64("authTokenDeadline", 604800, "token有效期")
	qiniuAccessKey := flag.String("qiniuAccessKey", "zobcXoKuteQUmUFgRk6FR85kgRRslmqx4BZeEdrc", "七牛 accessKey")
	qiniuSecretKey := flag.String("qiniuSecretKey", "0enF1nuLkZAqBztJfbhdxIwcE59FTuqtkbazGf3J", "七牛 secretKey")
	qiniuBucket := flag.String("qiniuBucket", "quickex-storage", "qiniu bucket")
	gaoDeKey := flag.String("gaoDeKey", "b0f725fc97052704dc65557051ceee6f", "高德key")
	flag.Parse()

	CONF = new(config)
	CONF.HttpAddr = *httpAddr
	CONF.EncryptedKey = *encryptedKey
	CONF.DataSourceName = *dataSourceName
	CONF.AuthTokenStart = *authTokenStart
	CONF.AuthTokenDeadline = *authTokenDeadline
	CONF.QiniuAccessKey = *qiniuAccessKey
	CONF.QiniuSecretKey = *qiniuSecretKey
	CONF.QiniuBucket = *qiniuBucket
	CONF.GaoDeKey = *gaoDeKey
}
