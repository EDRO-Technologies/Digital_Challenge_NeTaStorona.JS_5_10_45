package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http"
	C "other-side/config"
	T "other-side/types"
	"time"
)

//func Map[T, V any](ts []T, fn func(T) V) []V {
//	result := make([]V, len(ts))
//	for i, t := range ts {
//		result[i] = fn(t)
//	}
//	return result
//}
//
//func BoolRoleToString(isAI bool) string {
//	if isAI {
//		return "assistant"
//	} else {
//		return "user"
//	}
//}

func MakeModelRequest(msg *T.UserRequest) *T.ModelRequest {
	return &T.ModelRequest{
		Model:       "mistral-large-latest",
		Temperature: 0.7,
		TopP:        1.0,
		MaxTokens:   1000,
		Stream:      false,
		Stop:        "geeez",
		RandomSeed:  0,
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{Role: "system", Content: msg.Context},
			{Role: "user", Content: msg.Message},
		},
		ResponseFormat: struct {
			Type string `json:"type"`
		}{
			Type: "text",
		},
		Tools:            nil,
		ToolChoice:       "auto",
		PresencePenalty:  0,
		FrequencyPenalty: 0,
		N:                1,
		SafePrompt:       true,
	}
}

func MakeModelRequestImage(msg *T.UserRequestImage) *T.ModelRequestImage {
	return &T.ModelRequestImage{
		Model:       "pixtral-12b-latest",
		Stop:        "geeez",
		Temperature: 1,
		TopP:        0.3,
		MaxTokens:   100,
		Stream:      false,
		Messages: []struct {
			Role    string `json:"role"`
			Content []struct {
				Type     string `json:"type"`
				Text     string `json:"text,omitempty"`
				ImageUrl string `json:"image_url,omitempty"`
			} `json:"content"`
		}{
			{
				Role: "system",
				Content: []struct {
					Type     string `json:"type"`
					Text     string `json:"text,omitempty"`
					ImageUrl string `json:"image_url,omitempty"`
				}{
					{
						Type: "text",
						Text: "Извлеки из изображения документа запроса user название документа (name), место или степень (grade) и укажи уровень участия (городской, региональный, всероссийский, межрегиональный, международный) (level).\n" +
							"Ответь исключительно по ровно одному слову в данном формате: 'name:string,grade:number,level:string' \n" +
							"Если документ не относится к IT, хакатонам, программированию, hack: score = 0.",
					},
				},
			},
			{
				Role: "user",
				Content: []struct {
					Type     string `json:"type"`
					Text     string `json:"text,omitempty"`
					ImageUrl string `json:"image_url,omitempty"`
				}{
					{
						Type:     "image_url",
						ImageUrl: msg.ImageUrl,
					},
				},
			},
		},
		ResponseFormat: struct {
			Type string `json:"type"`
		}{
			Type: "text",
		},
		ToolChoice:       "auto",
		PresencePenalty:  0,
		FrequencyPenalty: 0,
		N:                1,
		SafePrompt:       false,
	}
}

func UseModel(prompt *T.UserRequest) (*T.ModelResponse, *T.ServiceError) {
	req := MakeModelRequest(prompt)

	reqJSON, err := json.Marshal(req)
	if err != nil {
		return nil, &T.ServiceError{
			Message: "Косяк перепаковки реквеста",
			Error:   err,
			Code:    fiber.StatusInternalServerError,
		}
	}

	reqJSONBytes := bytes.NewBuffer(reqJSON)

	promptReq, err := http.NewRequest("POST", C.Conf.Model, reqJSONBytes)
	if err != nil {
		return nil, &T.ServiceError{
			Message: "Косяк в реквесте",
			Error:   err,
			Code:    fiber.StatusInternalServerError,
		}
	}

	promptReq.Header.Set("Content-Type", "application/json")
	promptReq.Header.Set("Accept", "application/json")
	promptReq.Header.Set("Authorization", C.Conf.ModelToken)

	client := http.Client{Timeout: 80 * time.Second}

	res, err := client.Do(promptReq)

	if err != nil {
		return nil, &T.ServiceError{
			Message: "Ploho",
			Error:   err,
			Code:    fiber.StatusInternalServerError,
		}
	}

	if res.Body == nil {
		return nil, &T.ServiceError{
			Message: "Empty body from LLM",
		}
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, &T.ServiceError{
			Message: "Unable to get response body",
			Error:   err,
			Code:    fiber.StatusInternalServerError,
		}
	}

	responseBody := string(body)

	_ = responseBody

	var modelResponse T.ModelResponse

	err = json.Unmarshal([]byte(responseBody), &modelResponse)
	if err != nil {
		return nil, &T.ServiceError{
			Message: fmt.Sprintf("Unmarshalling response error: %s", responseBody),
			Error:   err,
			Code:    0,
		}
	}

	return &modelResponse, nil
}

func UseModelImage(prompt *T.UserRequestImage) (*T.ModelResponse, *T.ServiceError) {
	req := MakeModelRequestImage(prompt)

	reqJSON, err := json.Marshal(req)
	if err != nil {
		return nil, &T.ServiceError{
			Message: "Косяк перепаковки реквеста",
			Error:   err,
			Code:    fiber.StatusInternalServerError,
		}
	}

	reqJSONBytes := bytes.NewBuffer(reqJSON)

	promptReq, err := http.NewRequest("POST", C.Conf.Model, reqJSONBytes)
	if err != nil {
		return nil, &T.ServiceError{
			Message: "Косяк в реквесте",
			Error:   err,
			Code:    fiber.StatusInternalServerError,
		}
	}

	promptReq.Header.Set("Content-Type", "application/json")
	promptReq.Header.Set("Accept", "application/json")
	promptReq.Header.Set("Authorization", C.Conf.ModelToken)

	client := http.Client{Timeout: 80 * time.Second}

	res, err := client.Do(promptReq)

	if err != nil {
		return nil, &T.ServiceError{
			Message: "Ploho",
			Error:   err,
			Code:    fiber.StatusInternalServerError,
		}
	}

	if res.Body == nil {
		return nil, &T.ServiceError{
			Message: "Empty body from LLM",
		}
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, &T.ServiceError{
			Message: "Unable to get response body",
			Error:   err,
			Code:    fiber.StatusInternalServerError,
		}
	}

	responseBody := string(body)

	_ = responseBody

	var modelResponse T.ModelResponse

	err = json.Unmarshal([]byte(responseBody), &modelResponse)
	if err != nil {
		return nil, &T.ServiceError{
			Message: fmt.Sprintf("Unmarshalling response error: %s", responseBody),
			Error:   err,
			Code:    0,
		}
	}

	return &modelResponse, nil
}
