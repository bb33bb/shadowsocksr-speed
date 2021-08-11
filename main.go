package main

import (
	"encoding/json"
	"fmt"
	"github.com/ismdeep/args"
	"io/ioutil"
	"os/exec"
	"strconv"
	"time"
)

type ConfigFile struct {
	Configs []struct {
		Enable        bool   `json:"enable"`
		ID            string `json:"id"`
		Method        string `json:"method"`
		OBFS          string `json:"obfs"`
		OBFSParam     string `json:"obfsparam"`
		Password      string `json:"password"`
		Protocol      string `json:"protocol"`
		ProtocolParam string `json:"protocolparam"`
		Remarks       string `json:"remarks"`
		RemarksBase64 string `json:"remarks_base64"`
		Server        string `json:"server"`
		ServerPort    int    `json:"server_port"`
	} `json:"configs"`
}

func helpMsg() string {
	return `Usage: shadowsocksr-speed -c <config.json>`
}

func main() {
	if !args.Exists("-c") || args.Exists("--help") {
		fmt.Println(helpMsg())
		return
	}

	contentByte, err := ioutil.ReadFile(args.GetValue("-c"))
	if err != nil {
		panic(err)
	}

	config := &ConfigFile{}
	if err := json.Unmarshal(contentByte, config); err != nil {
		panic(err)
	}

	for _, conf := range config.Configs {
		fmt.Println(conf)
		cmd := exec.Command("python", "/home/ismdeep/Data/Projects/shadowsocksr/shadowsocks/local.py",
			"-s", conf.Server,
			"-p", strconv.Itoa(conf.ServerPort),
			"-k", conf.Password,
			"-m", conf.Method,
			"-O", conf.Protocol,
			"-G", conf.ProtocolParam,
			"-o", conf.OBFS,
			"-b", "0.0.0.0",
			"-l", "1080")
		if err := cmd.Start(); err != nil {
			panic(err)
		}


		time.Sleep(500 * time.Second)
		if err := cmd.Process.Kill(); err != nil {
			panic(err)
		}
	}

}
