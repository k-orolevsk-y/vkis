package api

import (
	"log"
	"net/url"
	"strings"
	"vkIntership/vk/models"
)

func (a *Api) UsersGet(userIds []string, fields string) (models.ResultUsersGet, error) {
	params := make(url.Values)
	params.Add("user_ids", strings.Join(userIds, ","))
	params.Add("fields", fields)

	resp, err := a.req.MakeRequest("users.get", params)
	if err != nil {
		if *a.req.Debug {
			log.Printf("[Debug] %v - error users get.", err)
		}

		return models.ResultUsersGet{}, err
	}

	var result models.ResultUsersGet
	if err = a.decodeBody(resp, &result); err != nil {
		return result, err
	}

	if err = a.checkError(resp); err != nil {
		if *a.req.Debug {
			log.Printf("[Debug] %v - error users get.", err)
		}

		return result, err
	}

	if err != nil && *a.req.Debug {
		log.Printf("[Debug] %v - error users get.", err)
	}

	return result, nil
}
