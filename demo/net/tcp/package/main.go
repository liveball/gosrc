package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

// packet.LinkLayer() // 以太网
// packet.NetworkLayer() // 网络层，通常也就是 IP 层
// packet.TransportLayer() // 传输层，比如 TCP/UDP
// packet.ApplicationLayer() // 应用层，比如 HTTP 层。
// packet.ErrorLayer() // ……出错了

func main() {
	// test()
	getRedis()
}

func test() {
	//  获取 libpcap 的版本
	version := pcap.Version()
	fmt.Println(version)
	//  获取网卡列表
	var devices []pcap.Interface
	devices, _ = pcap.FindAllDevs()
	fmt.Println(devices)
}

func getRedis() {
	//  获取 libpcap 的版本
	version := pcap.Version()
	fmt.Println(version)

	var device string
	if device = findNetName("127."); device == "" {
		panic("not net is prefix 127.")
	}

	handle, e := pcap.OpenLive(
		device,         // device
		int32(65535),   //  snapshot length
		false,          //  promiscuous mode?
		-1*time.Second, // timeout 负数表示不缓存，直接输出
	)
	if e != nil {
		panic(e.Error())
	}
	defer handle.Close()

	handle.SetBPFFilter("dst port 6379")
	packetSource := gopacket.NewPacketSource(
		handle,
		handle.LinkType(),
	)

	for packet := range packetSource.Packets() {
		//  解析 IP 层
		ipLayer := packet.Layer(layers.LayerTypeIPv4)
		if ipLayer != nil {
			//  解析 TCP 层
			tcpLayer := packet.Layer(layers.LayerTypeTCP)
			if tcpLayer != nil {
				tcp, _ := tcpLayer.(*layers.TCP)
				if len(tcp.Payload) > 0 {
					ip, _ := ipLayer.(*layers.IPv4)
					fmt.Printf("%s:%s->%s:%s\n%s\n",
						ip.SrcIP, tcp.SrcPort,
						ip.DstIP, tcp.DstPort,
						string(tcp.Payload))
				}
			} else if errLayer := packet.ErrorLayer(); errLayer != nil {
				fmt.Printf("tcp.err: %v", errLayer)
			}
		} else if errLayer := packet.ErrorLayer(); errLayer != nil {
			fmt.Printf("ip.err: %v", errLayer)
		}
	}

	return
}

func findNetName(prefix string) string {
	//  获取网卡列表
	var devices []pcap.Interface
	devices, _ = pcap.FindAllDevs()
	for _, d := range devices {
		for _, addr := range d.Addresses {
			if ip4 := addr.IP.To4(); ip4 != nil {
				if strings.HasPrefix(ip4.String(), prefix) {
					data, _ := json.MarshalIndent(d, "", "  ")
					fmt.Println(string(data))
					return d.Name
				}
			}
		}
	}
	return ""
}
