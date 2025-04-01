package ollama

import (
	"github.com/Meduzz/helper/http/client"
	"github.com/Meduzz/helper/http/herror"
)

func Chat(data *ChatRequest) (*ChatResponse, error) {
	req, err := client.POST("http://localhost:11434/api/chat", data)

	if err != nil {
		return nil, err
	}

	res, err := req.DoDefault()

	if err != nil {
		return nil, err
	}

	err = herror.IsError(res.Code())

	if err != nil {
		bs, _ := res.AsBytes()

		println(string(bs))

		return nil, err
	}

	resp := &ChatResponse{}
	err = res.AsJson(resp)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func List() (*ListModelResponse, error) {
	req, err := client.GET("http://localhost:11434/api/tags")

	if err != nil {
		return nil, err
	}

	res, err := req.DoDefault()

	if err != nil {
		return nil, err
	}

	err = herror.IsError(res.Code())

	if err != nil {
		return nil, err
	}

	resp := &ListModelResponse{}
	err = res.AsJson(resp)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
