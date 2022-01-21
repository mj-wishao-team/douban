package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppPort  string         `json:"app_port"`
	AppHost  string         `json:"app_host"`
	AppHttps bool           `json:"app_https"`
	Database DatabaseConfig `json:"database"`
	Email    EmailConfig    `json:"email"`
	Redis    RedisConfig    `json:"redis_config"`
	Jwt      JwtCfg         `json:"jwt"`
	Sms      SmsCfg         `json:"sms"`
	Oss      OssCfg         `json:"oss"`
}

type OssCfg struct {
	EndPoint     string `json:"end_point"`
	AppKey       string `json:"app_key"`
	AppSecret    string `json:"app_secret"`
	AvatarBucket string `json:"avatar_bucket"`
	AvatarUrl    string `json:"avatar_url"`
	VideosBucket string `json:"videos_bucket"`
	VideosUrl    string `json:"videos_url"`
}

type SmsCfg struct {
	LongBucketUrl string `json:"tencent_could_long_bucket_url"`
	TemplateId    string `json:"tencent_could_template_id"`
	SignName      string `json:"tencent_could_sign_name"`
	SmsSdkAppId   string `json:"tencent_could_sms_sdk_app_id"`
	AppId         string `json:"tencent_could_app_id"`
	SecretKey     string `json:"tencent_could_secret_key"`
	SecretId      string `json:"tencent_could_secret_id"`
	RegionAddr    string `json:"tencent_could_region_addr"`
	HttpEndpoint  string `json:"tencent_could_http_endpoint"`
}

type JwtCfg struct {
	MySecret string `json:"JwtSecret"`
}

type RedisConfig struct {
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

type EmailConfig struct {
	ServiceEmail string `json:"service_email"`
	ServicePwd   string `json:"service_pwd"`
	SmtpPort     string `json:"smtp_port"`
	SmtpHost     string `json:"smtp_host"`
}

type DatabaseConfig struct {
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"db_name"`
}

var cfg *Config

func GetCfg() *Config {
	return cfg
}

func init() {
	err := ParseCfg("./config/app.json")
	if err != nil {
		panic(err)
	}
}

func ParseCfg(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&cfg)
	if err != nil {
		return err
	}

	return nil
}
