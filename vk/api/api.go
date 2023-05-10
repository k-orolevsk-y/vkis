package api

import (
	"encoding/json"
	"errors"
	"log"
	"reflect"
	"vkIntership/vk/helpers"
	"vkIntership/vk/requests"
)

func New(req *requests.Requester) Api {
	return Api{
		req,
	}
}

func (a *Api) checkError(jsonBytes []byte) error {
	var jsonError struct{ Error helpers.Error }
	_ = json.Unmarshal(jsonBytes, &jsonError)

	if jsonError.Error.ErrorMsg != "" {
		return errors.New(jsonError.Error.ErrorMsg)
	}

	return nil
}

func (a *Api) decodeBody(jsonBytes []byte, object any) (err error) {
	if err = json.Unmarshal(jsonBytes, &object); err != nil {
		if *a.req.Debug {
			log.Printf("[Debug] %v - error decode body [%v].", err, reflect.TypeOf(object))
		}

		return
	}

	if *a.req.Debug {
		log.Printf("[Debug] %v - success decode body [%v].", object, reflect.TypeOf(object))
	}

	return
}
