package main

import (
	"Proxy/api2Ban"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

var extIP string
var user = flag.String("user", "user", "Auth user name")
var pass = flag.String("pass", "pass", "Auth password")
var addr = flag.String("addr", ":51461", "proxy address")
var go2ban = flag.String("go2ban", "1.1.1.1:2048", "go2ban gRPC address")

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	LogFile, err := os.OpenFile(time.Now().Format("06.01.02")+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalln("Dont create log file")
	}
	log.SetOutput(LogFile)

	flag.Parse()

	extIP = externalIP()

	proxy := &forwardProxy{}

	fmt.Println("Starting proxy server on", *addr)
	if err := http.ListenAndServe(*addr, proxy); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

type forwardProxy struct {
}

func (p *forwardProxy) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//Сплитит на хост, порт и ошибку
	host, _, _ := net.SplitHostPort(req.RemoteAddr)

	auth := req.Header.Get("Proxy-Authorization")
	if auth != `Basic `+base64.StdEncoding.EncodeToString([]byte(*user+":"+*pass)) {

		loginAttemptsIP.Increment(host)
		//chrome сперва заходит 3 раза без авторизации
		if loginAttemptsIP.Load(host) >= 10 {
			api2Ban.SendIp2BanGrpc(*go2ban, host)
		}

		w.Header().Add("Proxy-Authenticate", `Basic realm="Proxy Authorization"`)
		w.Header().Set("WWW-Authenticate", `Basic realm="Please enter your login and password"`)
		w.WriteHeader(http.StatusProxyAuthRequired)
		return
	}

	//Если один раз авторизован, то не блочить
	loginAttemptsIP.Unblock(host)

	hideIP(req.Header, extIP)

	//Работать только с https
	if req.Method == http.MethodConnect {
		proxyConnect(w, req)
	} else {
		http.Error(w, "This proxy only supports CONNECT", http.StatusMethodNotAllowed)
	}
}

func proxyConnect(w http.ResponseWriter, req *http.Request) {
	targetConn, err := net.Dial("tcp", req.Host)
	if err != nil {
		log.Println("failed to dial to target", req.Host)
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)

	hj, ok := w.(http.Hijacker)
	if !ok {
		log.Fatal("http server doesn't support hijacking connection")
	}
	clientConn, _, err := hj.Hijack()
	if err != nil {
		log.Fatal("http hijacking failed")
	}

	go tunnelConn(targetConn, clientConn)
	go tunnelConn(clientConn, targetConn)
}

func tunnelConn(dst io.WriteCloser, src io.ReadCloser) {
	io.Copy(dst, src)
	dst.Close()
	src.Close()
}

// Скрывать юзер IP
func hideIP(header http.Header, host string) {
	if prior, ok := header["X-Forwarded-For"]; ok {
		host = strings.Join(prior, ", ") + ", " + host
	}
	header.Set("X-Forwarded-For", host)
}

func externalIP() (ipS string) {
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Get("https://api.ipify.org?format=json")
	defer resp.Body.Close()
	if err != nil {
		log.Println("Can't GET to api.ipify.org ", err)
		return
	}
	//Мапа для JSON
	var ip map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&ip)
	if err != nil {
		log.Println("Wrong JSON api.ipify.org ", err)
		return
	}

	resIP, ok := ip["ip"].(string)
	if !ok {
		log.Fatal("Wrong JSON IP api.ipify.org ", err)
		return
	}

	return resIP
}
