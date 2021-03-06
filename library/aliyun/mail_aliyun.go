package aliyun

import (
	"context"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/treeyh/soc-go-common/core/errors"
	"strings"
)

func (aly *ALiYunProxy) SendSingleMail(ctx context.Context, accountName string, toAddress []string, replyToAddress bool,
	fromAlias, subject, htmlBody string) (*responses.CommonResponse, errors.AppError) {

	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Product = "DirectMail"     // Specify product
	request.Domain = "dm.aliyuncs.com" // Location Service will not be enabled if the host is specified. For example, service with a Certification type-Bearer Token should be specified
	request.Version = "2015-11-23"     // Specify product version
	request.Scheme = "https"           // Set request scheme. Default: http
	request.ApiName = "SingleSendMail" // Specify product interface
	request.AcceptFormat = "json"

	request.QueryParams["Action"] = "SingleSendMail"
	request.QueryParams["AddressType"] = "1"
	request.QueryParams["AccountName"] = accountName
	if replyToAddress {
		request.QueryParams["ReplyToAddress"] = "true"
	} else {
		request.QueryParams["ReplyToAddress"] = "false"
	}
	request.QueryParams["ToAddress"] = strings.Join(toAddress, ",")
	request.QueryParams["Subject"] = subject
	request.QueryParams["FromAlias"] = fromAlias
	request.FormParams["HtmlBody"] = htmlBody
	request.TransToAcsRequest()

	response, err := aly.client.ProcessCommonRequest(request)
	if err != nil {
		log.ErrorCtx2(ctx, err, "aliyun send mail error")
		return response, errors.NewAppError(errors.ALiYunSendMailError, err)
	}
	return response, nil
}
