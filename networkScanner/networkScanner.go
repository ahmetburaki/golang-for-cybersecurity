package main

import (
    "fmt"
    "net"
    "sync"
    "time"
)


func scanPort(host string, port int, results chan int, wg *sync.WaitGroup) {
    defer wg.Done()

    address := fmt.Sprintf("%s:%d", host, port)
    conn, err := net.DialTimeout("tcp", address, 1*time.Second)

    if err != nil {
        results <- 0 // Port is closed or unreachable
        return
    }

    conn.Close()
    results <- port // Port is open
}

func main() {
    var target string
    var startPort, endPort, timeout, maxConcurrent int

    fmt.Print("Enter target host: ")
    fmt.Scanln(&target)

    fmt.Print("Enter start port: ")
    fmt.Scanln(&startPort)

    fmt.Print("Enter end port: ")
    fmt.Scanln(&endPort)

    fmt.Print("Enter connection timeout (in seconds): ")
    fmt.Scanln(&timeout)

    fmt.Print("Enter maximum concurrent connections: ")
    fmt.Scanln(&maxConcurrent)

    fmt.Printf("Scanning %s...\n", target)

    var wg sync.WaitGroup
    results := make(chan int, maxConcurrent) // Use a buffered channel to control concurrency

    for port := startPort; port <= endPort; port++ {
        wg.Add(1)
        go scanPort(target, port, results, &wg)

        // Control the number of concurrent scans
        if len(results) >= maxConcurrent {
            wg.Wait()
        }
    }

    go func() {
        wg.Wait()
        close(results)
    }()

    openPorts := []int{}
    for port := range results {
        if port != 0 {
            openPorts = append(openPorts, port)
        }
    }

    if len(openPorts) == 0 {
        fmt.Println("No open ports found.")
    } else {
        fmt.Println("Open ports:")
        for _, port := range openPorts {
            fmt.Printf("Port %d is open\n", port)
        }
    }
}
