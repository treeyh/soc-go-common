package http_client

import (
	"context"
	"crypto/tls"
	"github.com/SkyAPM/go2sky"
	"github.com/treeyh/soc-go-common/core/config"
	"github.com/treeyh/soc-go-common/core/errors"
	"github.com/treeyh/soc-go-common/core/logger"
	"github.com/treeyh/soc-go-common/library/tracing"
	"io/ioutil"
	"net/http"
	neturl "net/url"
	v3 "skywalking.apache.org/repo/goapi/collect/language/agent/v3"
	"strings"
	"time"
)

const timeOutSecond = 3

var (
	_traceConfig = &config.TraceConfig{
		Enable:    false,
		Type:      "skywalking",
		Namespace: "",
		Server:    "",
	}
	log = logger.Logger()
)

func InitTraceConfig(traceConfig *config.TraceConfig) {
	_traceConfig = traceConfig
}

func Get(c context.Context, url string, querys map[string]string, headers map[string]string) (string, int, errors.AppError) {
	return do(c, "GET", url, querys, headers, "")
}

func Post(c context.Context, url string, querys map[string]string, headers map[string]string, body string) (string, int, errors.AppError) {
	return do(c, "POST", url, querys, headers, body)
}

func Put(c context.Context, url string, querys map[string]string, headers map[string]string, body string) (string, int, errors.AppError) {
	return do(c, "PUT", url, querys, headers, body)
}

func Delete(c context.Context, url string, querys map[string]string, headers map[string]string, body string) (string, int, errors.AppError) {
	return do(c, "DELETE", url, querys, headers, body)
}

func do(ctx context.Context, method string, url string, querys map[string]string, headers map[string]string, body string) (string, int, errors.AppError) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Timeout:   time.Second * timeOutSecond, //默认3秒超时时间
		Transport: tr,
	}

	// 拼接url
	reqUrl := url
	if querys != nil {
		reqUrl += "?" + ConvertToQueryParams(querys)
	}

	logmsg := method + "  url:" + reqUrl
	var req *http.Request
	var err error

	if body != "" {
		logmsg += "  body:" + body
		req, err = http.NewRequest(method, reqUrl, strings.NewReader(body))
	} else {
		req, err = http.NewRequest(method, reqUrl, nil)
	}
	if err != nil {
		log.ErrorCtx(ctx, logmsg+"  error:"+err.Error())
		return "", 0, errors.NewAppErrorByExistError(errors.HttpCreateRequestFail, err)
	}

	//设置header
	if headers != nil {
		header := "  header:"
		for k, v := range headers {
			header += k + "=" + v + " "
			req.Header.Set(k, v)
		}
		logmsg += header
	}

	var reqSpan go2sky.Span
	if tracing.GetTracer() != nil {
		/// 设置 skywalking span
		reqSpan, err = tracing.GetTracer().CreateExitSpan(ctx, url, url, func(headerKey, headerValue string) error {
			key := headerKey
			if _traceConfig.Namespace != "" {
				key = _traceConfig.Namespace + "-" + key
			}
			logmsg += key + "=" + headerValue + " "
			req.Header.Set(key, headerValue)
			return nil
		})
		if err != nil {
			log.ErrorCtx2(ctx, err, errors.SkyWalkingSpanNotInit.Error()+" url:"+url)
			return "", 0, errors.NewAppErrorByExistError(errors.SkyWalkingSpanNotInit, err)
		}
		reqSpan.SetComponent(tracing.HttpComponent)
		reqSpan.SetSpanLayer(v3.SpanLayer_Http)

		reqSpan.Tag(go2sky.TagHTTPMethod, method)
		reqSpan.Tag(go2sky.TagURL, url)
	}

	log.InfoCtx(ctx, logmsg)
	resp, err := client.Do(req)
	if err != nil {
		log.ErrorCtx(ctx, "error:"+err.Error())
		return "", 0, errors.NewAppErrorByExistError(errors.HttpRequestFail, err)
	}
	if reqSpan != nil {
		reqSpan.End()
	}

	b, err := ioutil.ReadAll(resp.Body)
	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()

	if err != nil {
		log.ErrorCtx(ctx, "error:"+err.Error())
		return "", resp.StatusCode, errors.NewAppErrorByExistError(errors.HttpRequestFail, err)
	}
	content := string(b)
	log.InfoCtx(ctx, "result:"+content)

	return content, resp.StatusCode, nil
}

func ConvertToQueryParams(queryParams map[string]string) string {
	rq := neturl.Values{}
	for k, v := range queryParams {
		rq.Add(k, v)
	}
	return rq.Encode()
}
