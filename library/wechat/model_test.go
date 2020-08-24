package wechat

import (
	"encoding/xml"
	"github.com/smartystreets/goconvey/convey"
	"github.com/treeyh/soc-go-common/core/utils/json"
	"github.com/treeyh/soc-go-common/tests"
	"testing"
)

func TestModel(t *testing.T) {

	convey.Convey("TestModel test", t, tests.TestStartUp(func() {

		ctx := tests.GetNewContext()

		xmlStr := `<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>123456789</CreateTime>
  <MsgType><![CDATA[event]]></MsgType>
  <Event><![CDATA[LOCATION]]></Event>
  <Latitude>23.137466</Latitude>
  <Longitude>113.352425</Longitude>
  <Precision>119.385040</Precision>
</xml>`

		v := &WechatReqestMsg{}

		err := xml.Unmarshal([]byte(xmlStr), &v)
		convey.So(err, convey.ShouldBeNil)
		log.InfoCtx(ctx, json.ToJsonIgnoreError(v))

		xmlStr = `<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>12345678</CreateTime>
  <MsgType><![CDATA[news]]></MsgType>
  <ArticleCount>1</ArticleCount>
  <Articles>
    <item>
      <Title><![CDATA[title1]]></Title>
      <Description><![CDATA[description1]]></Description>
      <PicUrl><![CDATA[picurl]]></PicUrl>
      <Url><![CDATA[url<>]]></Url>
    </item>
  </Articles>
</xml>`

		vv := &WechatResponseArticlesMsg{}
		err = xml.Unmarshal([]byte(xmlStr), &vv)
		convey.So(err, convey.ShouldBeNil)
		log.InfoCtx(ctx, json.ToJsonIgnoreError(vv))

		bys, err1 := xml.Marshal(vv)
		log.InfoCtx(ctx, string(bys))
		convey.So(err1, convey.ShouldBeNil)

	}, initWechatTestConfig))

}
