// nsq部分Http和Tcp请求封装
package nsq

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"net/http"

	"log"

	"github.com/Terry-Mao/goim/internal/logic/dao/mmsc/client/dsd"
)

type Producer struct {
	RemoteAddress    string   `json:"remote_address"`
	Hostname         string   `json:"hostname"`
	BroadcastAddress string   `json:"broadcast_address"`
	TcpPort          int      `json:"tcp_port"`
	HttpPort         int      `json:"http_port"`
	Version          string   `json:"version"`
	Tombstones       []bool   `json:"tombstones"`
	Topics           []string `json:"topics"`
}

// nsqLookUpd返回节点数据结构
type nsqNode struct {
	Producers []Producer `json:"producers"`
}

func GetNsqNodes() (*nsqNode, error) {
	url := fmt.Sprintf("http://%s/nodes", dsd.NsqLookupdHttp)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("get nsq nodes failed: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if len(body) == 0 {
		return nil, errors.New("get nsq nodes response is null")
	}

	node := nsqNode{}
	err = json.Unmarshal(body, &node)
	return &node, err
}

func DeleteChannel(prod Producer, topic, channel string) error {
	delurl := fmt.Sprintf("http://%s:%d/channel/delete?topic=%s&channel=%s",
		prod.BroadcastAddress, prod.HttpPort, topic, channel)
	_, err := http.Post(delurl, "application/json", nil)
	return err
}
