// Copyright (c) TFG Co. All Rights Reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/topfreegames/pitaya-cli/log"

	"github.com/abiosoft/ishell/v2"
	"github.com/mitchellh/go-homedir"
	"github.com/topfreegames/pitaya/v2/client"
)

func protoClient(preLog Log, addr string) error {
	preLog.Println("Using protobuf client")
	protoclient := client.NewProto(docsString, log.Log)
	pClient = protoclient

	for k, v := range pushInfo {
		protoclient.AddPushResponse(k, v)
	}

	if err := protoclient.LoadServerInfo(addr); err != nil {
		preLog.Println("Failed to load server info")
		return err
	}

	return nil
}

func tryConnect(addr string) error {
	return pClient.ConnectTo(addr)
}

func readServerMessages(callback func(data []byte)) {
	channel := pClient.MsgChannel()
	for {
		select {
		case <-disconnectedCh:
			close(disconnectedCh)
			return
		case m := <-channel:
			callback(parseData(m.Data))
		}
	}
}

func configure(c *ishell.Shell) {
	historyPath := os.Getenv("PITAYACLI_HISTORY_PATH")
	if historyPath == "" {
		home, _ := homedir.Dir()
		historyPath = fmt.Sprintf("%s/.pitayacli_history", home)
	}

	c.SetHistoryPath(historyPath)
}

func parseData(data []byte) []byte {
	if prettyJSON {
		var m interface{}
		_ = json.Unmarshal(data, &m)
		data, _ = json.MarshalIndent(m, "", "\t")
	}

	return data
}
