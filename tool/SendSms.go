package tool

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

func SendSms(phone, code string) error {
	SmsCfg := GetCfg().Sms
	credential := common.NewCredential(
		SmsCfg.SecretId,
		SmsCfg.SecretKey,
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = SmsCfg.HttpEndpoint
	client, _ := sms.NewClient(credential, SmsCfg.RegionAddr, cpf)
	request := sms.NewSendSmsRequest()

	request.PhoneNumberSet = common.StringPtrs([]string{phone})
	request.SmsSdkAppId = common.StringPtr(SmsCfg.SmsSdkAppId)
	request.SignName = common.StringPtr(SmsCfg.SignName)
	request.TemplateId = common.StringPtr(SmsCfg.TemplateId)
	request.TemplateParamSet = common.StringPtrs([]string{code, "5"})
	response, err := client.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return err
	}
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", response.ToJsonString())
	return err
}
