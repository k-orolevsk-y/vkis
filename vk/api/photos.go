package api

import (
	"log"
	"net/url"
	"os"
	"vkIntership/vk/models"
)

func (a *Api) PhotosGetMessagesUploadServer(params url.Values) (models.ResultPhotosServer, error) {
	resp, err := a.req.MakeRequest("photos.getMessagesUploadServer", params)
	if err != nil {
		if *a.req.Debug {
			log.Printf("[Debug] %v - error photos get messages upload server.", err)
		}

		return models.ResultPhotosServer{}, err
	}

	var result models.ResultPhotosServer
	if err = a.decodeBody(resp, &result); err != nil {
		return result, err
	}

	if err = a.checkError(resp); err != nil {
		if *a.req.Debug {
			log.Printf("[Debug] %v - error photos get messages upload server.", err)
		}

		return result, err
	}

	if err != nil && *a.req.Debug {
		log.Printf("[Debug] %v - error photos get messages upload server.", err)
	}

	return result, nil
}

func (a *Api) PhotosUpload(server models.ResultPhotosServer, file *os.File) (models.ResultUploadServer, error) {
	resp, err := a.req.UploadFile(server.Response.UploadUrl, file)
	if err != nil {
		if *a.req.Debug {
			log.Printf("[Debug] %v - error photos upload server.", err)
		}

		return models.ResultUploadServer{}, err
	}

	var result models.ResultUploadServer
	if err = a.decodeBody(resp, &result); err != nil {
		return result, err
	}

	return result, nil
}

func (a *Api) PhotosSaveMessagesPhoto(params url.Values) (models.ResultSaveMessagesPhoto, error) {
	resp, err := a.req.MakeRequest("photos.saveMessagesPhoto", params)
	if err != nil {
		if *a.req.Debug {
			log.Printf("[Debug] %v - error photos save messages.", err)
		}

		return models.ResultSaveMessagesPhoto{}, err
	}

	var result models.ResultSaveMessagesPhoto
	if err = a.decodeBody(resp, &result); err != nil {
		return result, err
	}

	if err = a.checkError(resp); err != nil {
		if *a.req.Debug {
			log.Printf("[Debug] %v - error photos save messages.", err)
		}

		return result, err
	}

	if err != nil && *a.req.Debug {
		log.Printf("[Debug] %v - error photos save messages.", err)
	}

	return result, nil
}
