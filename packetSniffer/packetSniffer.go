package main

import (
    "fmt"
    "log"
    "os"

    "github.com/google/gopacket"
    "github.com/google/gopacket/layers"
    "github.com/google/gopacket/pcap"
)

func main() {
    iface := "en0" // Change this to your network interface name.

    // Açılmış bir dosya oluşturun veya var olanı açın
    file, err := os.Create("outputs.log")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Log çıktılarını dosyaya yönlendir
    log.SetOutput(file)

    handle, err := pcap.OpenLive(iface, 1600, true, pcap.BlockForever)
    if err != nil {
        log.Fatal(err)
    }
    defer handle.Close()

    fmt.Println("Packet Sniffer started. Listening for packets...")

    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

    for packet := range packetSource.Packets() {
        printPacketInfo(packet)
    }
}

func printPacketInfo(packet gopacket.Packet) {
    ethLayer := packet.Layer(layers.LayerTypeEthernet)
    if ethLayer != nil {
        ethPacket, _ := ethLayer.(*layers.Ethernet)
        fmt.Printf("Source MAC: %s\n", ethPacket.SrcMAC)
        fmt.Printf("Destination MAC: %s\n", ethPacket.DstMAC)

        // Log çıktılarını dosyaya yaz
        log.Printf("Source MAC: %s\n", ethPacket.SrcMAC)
        log.Printf("Destination MAC: %s\n", ethPacket.DstMAC)
    }

    ipLayer := packet.Layer(layers.LayerTypeIPv4)
    if ipLayer != nil {
        ip, _ := ipLayer.(*layers.IPv4)
        fmt.Printf("Source IP: %s\n", ip.SrcIP)
        fmt.Printf("Destination IP: %s\n", ip.DstIP)

        // Log çıktılarını dosyaya yaz
        log.Printf("Source IP: %s\n", ip.SrcIP)
        log.Printf("Destination IP: %s\n", ip.DstIP)
    }

    tcpLayer := packet.Layer(layers.LayerTypeTCP)
    if tcpLayer != nil {
        tcp, _ := tcpLayer.(*layers.TCP)
        fmt.Printf("Source Port: %d\n", tcp.SrcPort)
        fmt.Printf("Destination Port: %d\n", tcp.DstPort)

        // Log çıktılarını dosyaya yaz
        log.Printf("Source Port: %d\n", tcp.SrcPort)
        log.Printf("Destination Port: %d\n", tcp.DstPort)
    }

    udpLayer := packet.Layer(layers.LayerTypeUDP)
    if udpLayer != nil {
        udp, _ := udpLayer.(*layers.UDP)
        fmt.Printf("Source Port: %d\n", udp.SrcPort)
        fmt.Printf("Destination Port: %d\n", udp.DstPort)

        // Log çıktılarını dosyaya yaz
        log.Printf("Source Port: %d\n", udp.SrcPort)
        log.Printf("Destination Port: %d\n", udp.DstPort)
    }

    applicationLayer := packet.ApplicationLayer()
    if applicationLayer != nil {
        payload := applicationLayer.Payload()
        if len(payload) > 0 {
            fmt.Printf("Payload (first 64 bytes): %s\n", truncateString(string(payload), 64))

            // Log çıktılarını dosyaya yaz
            log.Printf("Payload (first 64 bytes): %s\n", truncateString(string(payload), 64))
        }
    }

    fmt.Println("--------------------------------------------------")
    log.Println("--------------------------------------------------")
}

func truncateString(s string, maxLength int) string {
    if len(s) > maxLength {
        return s[:maxLength] + "..."
    }
    return s
}
