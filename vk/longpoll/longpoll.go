package longpoll

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
	"vkIntership/vk/api"
	"vkIntership/vk/models"
	"vkIntership/vk/requests"
)

func New(api *api.Api, req *requests.Requester, group *models.Group) LongPoll {
	return LongPoll{
		api:   api,
		req:   req,
		group: group,
	}
}

func (lp *LongPoll) NewUpdateConfig() (UpdateConfig, error) {
	server, err := lp.api.GroupsGetLongPollServer(lp.group.Id)
	if err != nil {
		return UpdateConfig{}, err
	} else if server.Error.ErrorCode != 0 {
		return UpdateConfig{}, errors.New(server.Error.ErrorMsg)
	}

	return UpdateConfig{
		Key:    server.Response.Key,
		Server: server.Response.Server,
		Ts:     server.Response.Ts,
		Wait:   25,
	}, nil
}

func (lp *LongPoll) GetUpdatesChan(config UpdateConfig) UpdatesChannel {
	ch := make(chan Update, 0)

	go func() {
		for {
			select {
			case <-lp.shutdownChannel:
				close(ch)
				return
			default:
			}

			result, err := lp.GetUpdates(config)
			if err != nil && *lp.req.Debug {
				log.Printf("[Debug] %v - error get updates.", err)
				time.Sleep(time.Second)

				continue
			}

			err = lp.checkFailed(result, &config)
			if err != nil {
				log.Printf("[Debug] %v - error get updates.", err)
				time.Sleep(time.Second)

				continue
			}

			config.Ts = result.Ts
			for _, update := range result.Updates {
				ch <- update
			}
		}
	}()

	return ch
}

func (lp *LongPoll) GetUpdates(config UpdateConfig) (Updates, error) {
	if *lp.req.Debug {
		log.Printf("[Debug] LongPoll; Config: %#v\n", config)
	}

	params := make(url.Values)
	params.Add("act", "a_check")
	params.Add("key", config.Key)
	params.Add("ts", config.Ts)
	params.Add("wait", fmt.Sprint(config.Wait))

	req, err := http.NewRequest("POST", config.Server, strings.NewReader(params.Encode()))
	if err != nil {
		return Updates{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if *lp.req.Debug {
			log.Printf("[Debug] LongPoll; Config: %#v, error request: %v\n", config, err)
		}

		return Updates{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil && *lp.req.Debug {
			log.Printf("[Debug] LongPoll; Config: %#v, error body close: %v\n", config, err)
		}
	}(resp.Body)

	decodeBody, err := io.ReadAll(resp.Body)
	if err != nil {
		if *lp.req.Debug {
			log.Printf("[Debug] LongPoll; Config: %#v, error body decode: %v\n", config, err)
		}

		return Updates{}, err
	}

	var result Updates
	err = json.Unmarshal(decodeBody, &result)

	return result, err
}

func (lp *LongPoll) checkFailed(result Updates, cfg *UpdateConfig) (err error) {
	switch result.Failed {
	case 0:
		cfg.Ts = result.Ts
	case 1:
		cfg.Ts = result.Ts
	case 2:
	case 3:
		newCfg, e := lp.NewUpdateConfig()
		err = e

		newCfg.Wait = cfg.Wait
		if result.Failed == 2 {
			newCfg.Ts = cfg.Ts
		}
		cfg = &newCfg
	default:
		err = errors.New(fmt.Sprintf("error #%v get updates", result.Failed))
	}

	return
}

func (lp *LongPoll) StopReceivingUpdates() {
	if *lp.req.Debug {
		log.Println("[Debug] Stop updates chan.")
	}
	close(lp.shutdownChannel)
}
