package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

// EsPage for es page.
type EsPage struct {
	Num   int `json:"num"`
	Size  int `json:"size"`
	Total int `json:"total"`
}

// EsArc for search archive.
type EsArc struct {
	OID      int64    `json:"oid"`
	TID      []int64  `json:"tid"`
	Business int      `json:"business"`
	Title    []string `json:"title"` //highlight
}

// SearchResult archive list from search.
type SearchResult struct {
	Page   *EsPage  `json:"page"`
	Result []*EsArc `json:"result"`
}

//LinkTag for link tag.
type LinkTag struct {
	ID     int64 `json:"id"`
	TID    int64 `json:"tid"`
	LinkID int64 `json:"link_id"`
}

var (
	test = `{"code":0,"message":"0","ttl":1,"data":{"order":"hot","sort":"desc","result":[{"business":1,"oid":23499586,"tid":[4,30,54]},{"title":["哪种画风最吃香？抄袭画风算侵权吗？如何找到属于\u003cem class=\"keyword\"\u003e自己\u003c/em\u003e的画风？【抖抖村】"]},{"business":1,"oid":22868817,"tid":[4,30,54]},{"title":["你是什么等级的\u003cem class=\"keyword\"\u003e画\u003c/em\u003e手？自测目前绘\u003cem class=\"keyword\"\u003e画\u003c/em\u003e实力和段位水平【抖抖村】"]},{"business":1,"oid":28841222,"tid":[4,30,53]},{"title":["看画师 Misa是怎么在手机绘\u003cem class=\"keyword\"\u003e画\u003c/em\u003e软件上指绘动漫小人儿哒？"]},{"business":1,"oid":26596016,"tid":[4,30,54]},{"title":["绘\u003cem class=\"keyword\"\u003e画\u003c/em\u003e能自学成才吗？和美院正规军到底有多大差距？抖叽～【抖抖村】"]},{"business":1,"oid":26937697,"tid":[5,30,54]},{"title":["沙包的零基础二次元绘\u003cem class=\"keyword\"\u003e画\u003c/em\u003e系列教学 第一期 | 人头部结构概括及结构剖析"]},{"business":1,"oid":19219788,"tid":[5,30,54]},{"title":["前方高能！解密胸部画法，请保持镇静！淡定！坐\u003cem class=\"keyword\"\u003e好\u003c/em\u003e扶稳！"]},{"business":1,"oid":28808281,"tid":[4,30,54]},{"title":["靠“多画”真的就能画好？提高绘\u003cem class=\"keyword\"\u003e画\u003c/em\u003e水平的捷径【抖抖村】"]},{"business":1,"oid":14894164,"tid":[4,30,54]},{"title":["专业治疗人体废20年，看完后你就知道\u003cem class=\"keyword\"\u003e自己\u003c/em\u003e错在哪了！"]},{"business":1,"oid":22550673,"tid":[4,30,54]},{"title":["画画有前途吗？怎么练提高最快？一次性解答萌新\u003cem class=\"keyword\"\u003e画\u003c/em\u003e手所有问题！【抖抖村】"]},{"business":1,"oid":23185387,"tid":[4,30,54]},{"title":["能回答上来的都是高段\u003cem class=\"keyword\"\u003e画\u003c/em\u003e手！那些让人哭笑不得的绘\u003cem class=\"keyword\"\u003e画\u003c/em\u003e问题【抖抖村】"]},{"business":1,"oid":25668053,"tid":[4,30,52]},{"title":["测测\u003cem class=\"keyword\"\u003e自己\u003c/em\u003e是有创意的人吗？怎么掌握创意这个人类最初和最终的武器？【抖抖村】"]},{"business":1,"oid":19355906,"tid":[5,30,28]},{"title":["如何用手机把\u003cem class=\"keyword\"\u003e画\u003c/em\u003e拍出扫描仪的效果"]},{"business":1,"oid":5466443,"tid":[4,30,17]},{"title":["【Vegas教程】【剪辑基础篇】【\u003cem class=\"keyword\"\u003e自己\u003c/em\u003e看，别问我系列_(:з」∠)_·再缠就删】by风墨逐辰"]},{"business":1,"oid":21938902,"tid":[4,30,48]},{"title":["史上最全CG绘画工具大评测-软件篇！\u003cem class=\"keyword\"\u003e画\u003c/em\u003e手动画师必看！【抖抖村】"]},{"business":1,"oid":31564623,"tid":[5,30,60]},{"title":["\u003cem class=\"keyword\"\u003e自己\u003c/em\u003e\u003cem class=\"keyword\"\u003e画\u003c/em\u003e的好没人看，别人\u003cem class=\"keyword\"\u003e画\u003c/em\u003e的烂却被吹捧？如何推广\u003cem class=\"keyword\"\u003e自己\u003c/em\u003e的作品【抖抖村】"]},{"business":1,"oid":9338941,"tid":[5,30,53]},{"title":["［MediBang Colors］填色\u003cem class=\"keyword\"\u003e画\u003c/em\u003e现场制作"]},{"business":1,"oid":4308434,"tid":[4,30,28]},{"title":["一分钟教会你ps！不会做特效的摄影师不是\u003cem class=\"keyword\"\u003e好\u003c/em\u003e男票【LHG菌】"]},{"business":1,"oid":20748710,"tid":[5,30,52]},{"title":["怎么拥有神(经病)一般的编剧能力!  漫\u003cem class=\"keyword\"\u003e画\u003c/em\u003e手、文案必修秘技！安卓android和苹果apple的战争【抖抖村】"]},{"business":1,"oid":21321460,"tid":[4,30,48]},{"title":["史上最全的电脑CG绘\u003cem class=\"keyword\"\u003e画\u003c/em\u003e设备和软件大评测！\u003cem class=\"keyword\"\u003e画\u003c/em\u003e手必看攻略【抖抖村】"]},{"business":1,"oid":22242196,"tid":[4,30,48]},{"title":["松鼠都能看懂！ 绘\u003cem class=\"keyword\"\u003e画\u003c/em\u003e、做3D应该选什么配置的电脑【抖抖村】"]}],"debug":null,"page":{"num":1,"size":20,"total":126}}}`
)

func main() {
	// println(stringsToJSON("aaa"))

	var res struct {
		Code int           `json:"code"`
		Data *SearchResult `json:"data"`
	}

	err := json.Unmarshal([]byte(test), &res)

	if err != nil {
		fmt.Println(err)
		return
	}

	res2 := make([]*EsArc, 0, len(res.Data.Result)/2)

	// for k, v := range res.Data.Result {
	// 	fmt.Println(k, v)
	// }
	for i := 0; i < len(res.Data.Result)-1; i += 2 {
		res.Data.Result[i].Title = res.Data.Result[i+1].Title
		res2 = append(res2, res.Data.Result[i])
	}

	for k, v := range res2 {
		fmt.Println(k, v)
	}
	// spew.Dump(res)
}

func stringsToJSON(str string) string {
	var jsons bytes.Buffer
	for _, r := range str {
		rint := int(r)
		if rint < 128 {
			jsons.WriteRune(r)
		} else {
			jsons.WriteString("\\u")
			jsons.WriteString(strconv.FormatInt(int64(rint), 16))
		}
	}
	return jsons.String()
}
