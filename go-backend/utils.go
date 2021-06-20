package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/buger/jsonparser"
)

type Response struct {
	StatusCode int               `json:"statusCode"`
	Message    string            `json:"message"`
	Data       map[string]string `json:"data"`
}

func getSystemIp(w http.ResponseWriter, r *http.Request) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		jData, err2 := json.Marshal(Response{
			StatusCode: 200,
			Message:    "fail",
			Data:       map[string]string{"ip": err.Error()},
		})
		if err2 != nil {
			fmt.Fprintf(w, string(err2.Error()))
		}
		fmt.Fprintf(w, string(jData))
	}

	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	jData, err := json.Marshal(Response{
		StatusCode: 200,
		Message:    "success",
		Data:       map[string]string{"ip": localAddr.IP.String() + port},
	})

	if err != nil {
		fmt.Print(w, string(err.Error()))
	}

	fmt.Fprintf(w, string(jData))
}

func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func resolveHostIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Print("Server Running At :  http://")
	fmt.Print(localAddr.IP.String())
	fmt.Println(":7007")
	return localAddr.IP.String()
}

func getStringValue(val []byte, key string) string {
	value, err := jsonparser.GetString(val, key)
	if err == nil {
		return value
	}
	print(err)
	return "error"
}

func getBoolean(val []byte, key string) bool {
	value, err := jsonparser.GetBoolean(val, key)
	if err == nil {
		return value
	}
	print(err)
	return false
}

func getValue(val []byte, key string) []byte {
	value, dataType, offset, err := jsonparser.Get(val, key)
	if err == nil {
		return value
	} else {
		print(err)
		println(dataType)
		println(offset)
	}
	return []byte("error")
}
