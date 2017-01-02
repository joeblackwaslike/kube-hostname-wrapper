package main

/*
Wrapper that fixes container hostname resolution under kubernetes
*/

import (
    "net"
    "os"
    "fmt"
    "strings"
    "io/ioutil"
)

func readNsFile(path string) string {
    b, err := ioutil.ReadFile(path)
    if err != nil {
        ns := "default"
        return ns
    }
    ns := strings.TrimSpace(string(b))
    return ns
}

func getLocalIp() string {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        os.Exit(1)
    }

    var faddrs []string

    for _, a := range addrs {
        if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                faddrs = append(faddrs, ipnet.IP.String())
            }
        }
    }
    iaddr := faddrs[0]
    return iaddr
}

func formatIp(ip string) string {
    return strings.Replace(ip, ".", "-", -1)
}

func formatShort(ip string) string {
    return formatIp(ip)
}

func formatLong(ip string, ns string, suffix string) string {
    fip := formatIp(ip)

    return fmt.Sprintf("%s.%s.%s", fip, ns, suffix)
}

func main() {

    ns_path := "/var/run/secrets/kubernetes.io/serviceaccount/namespace"
    suffix := "pod.cluster.local"
    
    ns := readNsFile(ns_path)
    use_shortname := os.Getenv("KUBE_HOSTNAME_SHORT")
    
    os.Args = append(os.Args, "*")
    fmtArg := os.Args[1]

    ip := getLocalIp()

    switch fmtArg {
    case "-f":
        fmt.Println(formatLong(ip, ns, suffix))
    
    case "-s":
        fmt.Println(formatShort(ip))
    
    case "-i":
        fmt.Println(ip)

    default:
        if use_shortname == "true" {
            fmt.Println(formatShort(ip))
        } else {
            fmt.Println(formatLong(ip, ns, suffix))
        }
    }
}
