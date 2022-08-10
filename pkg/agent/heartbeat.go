package agent

import (
	"encoding/json"
	"fmt"
	"github.com/williamfzc/sidebike/pkg/server"
	"io"
	"net/http"
	"net/url"
	"time"
)

func (agent *Agent) GetUrlPing() (*url.URL, error) {
	ret, err := url.Parse(fmt.Sprintf("http://%s%s", agent.GetRegistry(), server.Ping.Path))
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Add(server.FieldMachineLabel, agent.MachineLabel)
	ret.RawQuery = params.Encode()
	return ret, nil
}

// signal only
func (agent *Agent) heartBeat() {
	period := agent.GetPeriod()
	for range time.Tick(period) {
		agent.triggerHeartBeat()
	}
}

func (agent *Agent) triggerHeartBeat() {
	resp, err := agent.requestHeartBeat()
	if err != nil {
		// err has already been printed
		return
	}
	switch resp.Signal {
	case server.SignalNewTask:
		agent.taskRequestQueue <- &Event{}
		logger.Infof("found new task. trying to request")
	default:
		logger.Debugf("heartbeat: %v", resp)
	}
}

func (agent *Agent) requestHeartBeat() (*server.Response, error) {
	finalUrl, err := agent.GetUrlPing()
	if err != nil {
		logger.Errorf("failed to gen lifecycle url: %s", err)
		return nil, err
	}
	resp, err := http.Get(finalUrl.String())
	if err != nil {
		logger.Errorf("lifecycle error: %s", err)
		return nil, err
	}

	responseObj := &server.Response{}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("read body error: %s", err)
		return nil, err
	}

	err = json.Unmarshal(data, responseObj)
	if err != nil {
		logger.Errorf("json parse error: %s", err)
		return nil, err
	}
	return responseObj, nil
}
