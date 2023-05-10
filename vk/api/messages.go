package api

import (
	"log"
	"net/url"
	"vkIntership/vk/models"
)

func (a *Api) MessagesSend(params url.Values) (models.ResultMessagesSend, error) {
	if params.Get("random_id") == "" {
		params.Set("random_id", "0")
	}

	resp, err := a.req.MakeRequest("messages.send", params)
	if err != nil {
		if *a.req.Debug {
			log.Printf("[Debug] %v - error messages send.", err)
		}

		return models.ResultMessagesSend{}, err
	}

	var result models.ResultMessagesSend
	if err = a.decodeBody(resp, &result); err != nil {
		return result, err
	}

	if err = a.checkError(resp); err != nil {
		if *a.req.Debug {
			log.Printf("[Debug] %v - error messages send.", err)
		}

		return result, err
	}

	if err != nil && *a.req.Debug {
		log.Printf("[Debug] %v - error messages send.", err)
	}

	return result, err
}

func (a *Api) MessagesSendWithPeerIds(params url.Values) (models.ResultMessagesSendPeerIds, error) {
	if params.Get("random_id") == "" {
		params.Set("random_id", "0")
	}

	resp, err := a.req.MakeRequest("messages.send", params)
	if err != nil {
		if *a.req.Debug {
			log.Printf("[Debug] %v - error messages send with peer ids.", err)
		}

		return models.ResultMessagesSendPeerIds{}, err
	}

	var result models.ResultMessagesSendPeerIds
	if err = a.decodeBody(resp, &result); err != nil {
		return result, err
	}

	if err = a.checkError(resp); err != nil {
		if *a.req.Debug {
			log.Printf("[Debug] %v - error messages send with peer ids.", err)
		}

		return result, err
	}

	if err != nil && *a.req.Debug {
		log.Printf("[Debug] %v - error messages send with peer ids.", err)
	}

	return result, nil
}

func (a *Api) MessagesSendMessageEventAnswer(params url.Values) (models.ResultMessagesSendMessageEventAnswer, error) {
	resp, err := a.req.MakeRequest("messages.sendMessageEventAnswer", params)
	if err != nil {
		if *a.req.Debug {
			log.Printf("[Debug] %v - error messages send message event answer.", err)
		}

		return models.ResultMessagesSendMessageEventAnswer{}, err
	}

	var result models.ResultMessagesSendMessageEventAnswer
	if err = a.decodeBody(resp, &result); err != nil {
		return result, err
	}

	if err = a.checkError(resp); err != nil {
		if *a.req.Debug {
			log.Printf("[Debug] %v - error messages send message event answer.", err)
		}

		return result, err
	}

	if err != nil && *a.req.Debug {
		log.Printf("[Debug] %v - error messages send message event answer.", err)
	}

	return result, nil
}

func (a *Api) MessagesEdit(params url.Values) (models.ResultMessageEdit, error) {
	resp, err := a.req.MakeRequest("messages.edit", params)
	if err != nil {
		if *a.req.Debug {
			log.Printf("[Debug] %v - error messages edit.", err)
		}

		return models.ResultMessageEdit{}, err
	}

	var result models.ResultMessageEdit
	if err = a.decodeBody(resp, &result); err != nil {
		return result, err
	}

	if err = a.checkError(resp); err != nil {
		if *a.req.Debug {
			log.Printf("[Debug] %v - error messages edit.", err)
		}

		return result, err
	}

	if err != nil && *a.req.Debug {
		log.Printf("[Debug] %v - error messages edit.", err)
	}

	return result, nil
}

func (a *Api) MessagesDelete(params url.Values) (models.ResultMessageDelete, error) {
	resp, err := a.req.MakeRequest("messages.delete", params)
	if err != nil {
		if *a.req.Debug {
			log.Printf("[Debug] %v - error messages delete.", err)
		}

		return models.ResultMessageDelete{}, err
	}

	var result models.ResultMessageDelete
	if err = a.decodeBody(resp, &result); err != nil {
		return result, err
	}

	if err = a.checkError(resp); err != nil {
		if *a.req.Debug {
			log.Printf("[Debug] %v - error messages delete.", err)
		}

		return result, err
	}

	if err != nil && *a.req.Debug {
		log.Printf("[Debug] %v - error messages delete.", err)
	}

	return result, nil
}
