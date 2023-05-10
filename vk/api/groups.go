package api

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"vkIntership/vk/models"
)

func (a *Api) GroupsGetById(groupIds ...string) (models.ResultGroupsGetById, error) {
	params := make(url.Values)
	params.Add("group_ids", strings.Join(groupIds, ","))

	resp, err := a.req.MakeRequest("groups.getById", params)
	if err != nil {
		if *a.req.Debug {
			log.Printf("[Debug] %v - error groups get by id.", err)
		}

		return models.ResultGroupsGetById{}, err
	}

	var result models.ResultGroupsGetById
	if err = a.decodeBody(resp, &result); err != nil {
		return result, err
	}

	if err = a.checkError(resp); err != nil {
		if *a.req.Debug {
			log.Printf("[Debug] %v - error groups get by id.", err)
		}

		return result, err
	}

	if err != nil && *a.req.Debug {
		log.Printf("[Debug] %v - error groups get by id.", err)
	}

	return result, err
}

func (a *Api) GroupsGetLongPollServer(groupId int) (models.ResultGetLongPollServer, error) {
	params := make(url.Values)
	params.Add("group_id", fmt.Sprint(groupId))

	resp, err := a.req.MakeRequest("groups.getLongPollServer", params)
	if err != nil {
		if *a.req.Debug {
			log.Printf("[Debug] %v - error groups get long poll server.", err)
		}

		return models.ResultGetLongPollServer{}, err
	}

	var result models.ResultGetLongPollServer
	if err = a.decodeBody(resp, &result); err != nil {
		return result, err
	}

	if err = a.checkError(resp); err != nil {
		if *a.req.Debug {
			log.Printf("[Debug] %v - error groups get long poll server.", err)
		}

		return result, err
	}

	if err != nil && *a.req.Debug {
		log.Printf("[Debug] %v - error groups get long poll server.", err)
	}

	return result, err
}
