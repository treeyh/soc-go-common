package http_client

import (
	"context"
	"crypto/tls"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/treeyh/soc-go-common/core/config"
	"github.com/treeyh/soc-go-common/core/consts"
	"github.com/treeyh/soc-go-common/core/errors"
	"github.com/treeyh/soc-go-common/core/logger"
	"io/ioutil"
	"net/http"
	neturl "net/url"
	"strings"
	"time"
)

var (
	tarceConfig = &config.TraceConfig{
		Enable: false,
		Server: "",
	}
	log = logger.Logger()
)

func InitTraceConfig(traceConfig *config.TraceConfig) {
	tarceConfig = traceConfig
}

func Get(c context.Context, url string, querys map[string]string) (string, int, errors.AppError) {
	return do(c, "GET", url, querys, nil, nil)
}

func Post(c context.Context, url string, querys map[string]string, body *string) (string, int, errors.AppError) {
	return do(c, "POST", url, querys, nil, body)
}

func Put(c context.Context, url string, querys map[string]string, body *string) (string, int, errors.AppError) {
	return do(c, "PUT", url, querys, nil, body)
}

func Delete(c context.Context, url string, querys map[string]string, body *string) (string, int, errors.AppError) {
	return do(c, "DELETE", url, querys, nil, body)
}

func do(ctx context.Context, method string, url string, querys map[string]string, headers map[string]string, body *string) (string, int, errors.AppError) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Timeout:   time.Second * 3, //默认3秒超时时间
		Transport: tr,
	}

	// 拼接url
	reqUrl := url
	if querys != nil {
		reqUrl += "?" + ConvertToQueryParams(querys)
	}

	var req *http.Request
	var err error

	if body != nil {
		req, err = http.NewRequest(method, reqUrl, strings.NewReader(*body))
	} else {
		req, err = http.NewRequest(method, reqUrl, nil)
	}

	if err != nil {
		return "", 0, errors.NewAppErrorByExistError(errors.HttpCreateRequestFail, err)
	}

	//设置header
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	if tarceConfig.Enable && ctx != nil {

		tracer := ctx.Value(consts.TracerContextKey)
		parentSpanContext := ctx.Value(consts.TraceParentSpanContextKey)

		if tracer != nil && parentSpanContext != nil {
			span := opentracing.StartSpan(
				"call Http "+method,
				opentracing.ChildOf(parentSpanContext.(opentracing.SpanContext)),
				opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
				ext.SpanKindRPCClient,
			)

			span.Finish()

			injectErr := tracer.(opentracing.Tracer).Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
			if injectErr != nil {
				log.Fatalf("%s: Couldn't inject headers", err)
			}
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", 0, errors.NewAppErrorByExistError(errors.HttpRequestFail, err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()

	if err != nil {
		logger.Logger().Error(err)
		return "", resp.StatusCode, errors.NewAppErrorByExistError(errors.HttpRequestFail, err)
	}
	return string(b), resp.StatusCode, nil
}

func ConvertToQueryParams(queryParams map[string]string) string {
	rq := neturl.Values{}
	for k, v := range queryParams {
		rq.Add(k, v)
	}
	return rq.Encode()
}
