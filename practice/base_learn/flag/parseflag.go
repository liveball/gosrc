package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var maxSrcPorts = flag.Int("maxSrcPorts", 256, "The maximum number of source ports to use")

func main() {
	married := flag.Bool("married", false, "Are you married?")
	age := flag.Int("age", 22, "How old are you?")
	name := flag.String("name", "", "What your name?")
	// fmt.Printf("%+v\n", married)

	var address string
	defaultAddress := "北京"
	//flag.StringVar这样的函数第一个参数换成了变量地址，后面的参数和flag.String是一样的。
	flag.StringVar(&address, "address", defaultAddress, "Where is your address?")
	flag.StringVar(&address, "addr", defaultAddress, "Where is your address?")

	var intervalFlag interval
	flag.Var(&intervalFlag, "deltaT", "逗号分割的时间间隔")
	flag.Parse() //解析输入的参数

	if flag.Arg(0) == "" {
		fmt.Fprintf(os.Stderr, "Must specify a target\n")
		return
	}
	for i := 0; i < len(flag.Args()); i++ { //flag.Arg(i)来获取非flag命令行参数
		fmt.Println("Arg", i, "=", flag.Arg(i))
	}
	fmt.Println("输出的参数married的值是:", *married) //不加*号的话,输出的是内存地址
	fmt.Println("输出的参数age的值是:", *age)
	fmt.Println("输出的参数name的值是:", *name)
	fmt.Println("输出的参数address的值是:", address)
	fmt.Println("输出的参数maxSrcPorts的值是:", *maxSrcPorts)
}

type interval []time.Duration

//实现String接口
func (i *interval) String() string {
	return fmt.Sprint(*i)
}

//实现Set接口,Set接口决定了如何解析flag的值
func (i *interval) Set(value string) error {
	// println(value)
	// spew.Dump(*i)
	// println(len(*i))
	if len(*i) > 0 { //此处决定命令行是否可以设置多次-deltaT
		return errors.New("over 2 interval flag set")
	}
	for _, dt := range strings.Split(value, ",") {
		duration, err := time.ParseDuration(dt)
		if err != nil {
			return err
		}
		*i = append(*i, duration)
	}
	return nil
}

func newflag() {
	var myFlagSet = flag.NewFlagSet("myflagset", flag.ExitOnError)
	abc := myFlagSet.String("abc", "default value", "help mesage")
	ghi := myFlagSet.Bool("def", true, "help mesage")
	// myFlagSet.Bool("fdf", true, "help mesage")

	myFlagSet.Parse([]string{"-abc", "abc-value", "-def", "sss"})
	fmt.Println("输出的参数abc的值是:", *abc)
	fmt.Println("输出的参数ghi的值是:", *ghi)
	args := myFlagSet.Args()
	for i := range args {
		fmt.Println(i, myFlagSet.Arg(i))
	}
}
