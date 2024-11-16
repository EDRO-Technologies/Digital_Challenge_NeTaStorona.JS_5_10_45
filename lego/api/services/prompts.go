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
		TopP:        1,
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

//func GetAnswered(prompt T.Prompt) (string, *T.ServiceError) {
//	if len(prompt.Message) == 0 {
//		return "", nil
//	}
//
//	msgs := Map(prompt.Message, func(item T.Message) DimaRequestMessage {
//		return DimaRequestMessage{
//			Role:    BoolRoleToString(item.IsAI),
//			Content: item.Content,
//		}
//	})
//
//	req := MakeDimaRequest(msgs)
//
//	reqJSON, err := json.Marshal(req)
//	if err != nil {
//		return "", &T.ServiceError{
//			Message: "Косяк перепаковки реквеста",
//			Error:   err,
//			Code:    fiber.StatusInternalServerError,
//		}
//	}
//
//	reqJSONBytes := bytes.NewBuffer(reqJSON)
//
//	apiURL := fmt.Sprintf("%s", C.Conf.ModelHost)
//
//	promptReq, err := http.NewRequest("POST", apiURL, reqJSONBytes)
//	if err != nil {
//		return "", &T.ServiceError{
//			Message: "Косяк в реквесте",
//			Error:   err,
//			Code:    fiber.StatusInternalServerError,
//		}
//	}
//
//	promptReq.Header.Set("Content-Type", "application/json")
//
//	client := http.Client{Timeout: 80 * time.Second}
//
//	res, err := client.Do(promptReq)
//
//	if err != nil {
//		return "", &T.ServiceError{
//			Message: "Ploho",
//			Error:   err,
//			Code:    fiber.StatusInternalServerError,
//		}
//	}
//
//	if res.Body == nil {
//		return "", &T.ServiceError{
//			Message: "Flawed response from LLM",
//		}
//	}
//
//	body, err := io.ReadAll(res.Body)
//
//	responseBody := string(body)
//
//	if err != nil {
//		return "", &T.ServiceError{
//			Message: "Unable to get federal act",
//			Error:   err,
//			Code:    fiber.StatusInternalServerError,
//		}
//	}
//
//	return responseBody, nil
//}
//
//func GetFinalAnswer(ctx context.Context) (*T.Prompt, *T.ServiceError) {
//	//federalAct, err := http.//TODO: first API call
//
//	if err != nil {
//		return nil, &T.ServiceError{
//			Message: "Unable to get federal act",
//			Error:   err,
//			Code:    fiber.StatusInternalServerError,
//		}
//	}
//
//	return federalAct, nil
//}
