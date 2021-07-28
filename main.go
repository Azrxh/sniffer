package main

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/layers"
	"log"
	"fmt"
	"time"
)

var (
	device string
	err 	error
	handle	*pcap.Handle
	timeout	= -1 * time.Second
)

func sniff() {
	fmt.Println("Sniffer running...")
	handle, err = pcap.OpenLive(device, 65535, true, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		ipLayer := packet.Layer(layers.LayerTypeIPv4)
		ipPacket := ipLayer.(*layers.IPv4)

		fmt.Println("src " + ipPacket.SrcIP.String() + " -> dst " + ipPacket.DstIP.String())
		fmt.Println("PROTOCOL: " + ipPacket.Protocol.String())
	}
}

func main() {
	fmt.Print("Catpure card~>: ")
	fmt.Scanln(&device)

	sniff()
}
