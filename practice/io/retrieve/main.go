package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	f, err := os.OpenFile("/data/app/go/src/gosrc/practice/io/retrieve/a.txt", os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	mf, err := os.OpenFile("/data/app/go/src/gosrc/practice/io/retrieve/mid.txt", os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer mf.Close()

	write := bufio.NewWriter(mf)

	var buf *bufio.Reader
	buf = bufio.NewReader(f)

	var i int64
	midMap := make(map[string]struct{})
	for {
		str, err := buf.ReadString('\n') //log agent
		if err == io.EOF {
			fmt.Printf("buf.ReadString error(%v)\n", err)
			break
		}

		//fmt.Println(str)

		var res Result
		if err = json.Unmarshal([]byte(str), &res); err != nil {
			fmt.Printf("json.Unmarshal str(%s) error(%v)\n", str, err)
			break
		}

		//fmt.Println(res.Source.Log)

		exp := regexp.MustCompile(`^RewardActivate s.newc.Lock mid\(([\d]+)\)|$`)
		params := exp.FindStringSubmatch(res.Source.Log)
		//fmt.Println(22,  params[1])

		fmt.Println(params[1])

		midMap[params[1]] = struct{}{}


		//comma := strings.Index(res.Source.Log, "(")
		//pos := strings.LastIndex(res.Source.Log[comma:], ")|")
		//fmt.Println(res.Source.Log[comma+1:], comma, pos)

		i++
	}


	fmt.Println("exit", i, len(midMap))

	for mid:= range midMap {
		write.WriteString(mid + "\n")
	}

	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}

//^$分别表示匹配的开始和结束，界定我们正则表达式的范围。
//
//[\d]{4}表示我们要正好匹配4位数字，因为年份是4位，所以我们定义为匹配4位。后面的月份和天是2位，所以定义为2位。
//
//[\w-]匹配字符串和中杠，加号(+)表示匹配1个或者多个。
//
//然后他们都加了括号()，意味着我们要提取这些字符串。

//{
//	"_index":"billions-main.archive.creative-@2020.01.24-jssz02-2",
//	"_type":"logs",
//	"_id":"AW_WjjqANiPc2RK9gNZB",
//	"_score":0,
//	"_source":{
//		"@timestamp":1579852636675,
//		"app_id":"main.archive.creative",
//		"env":"prod",
//		"instance_id":"creative-153048-6b6687c4f8-lxbzf",
//		"level":"INFO",
//		"level_value":1,
//		"log":"RewardActivate s.newc.Lock mid(13890986)|id(4700)",
//		"source":"/go/src/go-main/app/archive/creative/interface/service/newcomer/task.go:506",
//		"zone":"sh001"
//	}
//}

type Result struct {
	Index  string `json:"_index"`
	Source struct {
		Log string `json:"log"`
	} `json:"_source"`
}
