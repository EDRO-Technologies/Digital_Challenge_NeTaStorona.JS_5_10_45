package controllers

import (
	"github.com/gofiber/fiber/v2"
	S "other-side/api/services"
	H "other-side/handler"
	T "other-side/types"
)

// PostPrompt
// @Summary        PostPrompt
// @Description    Post prompt
// @Tags           Prompts
// @Accept         json
// @Produce        json
//
//	@Param          prompt  body    T.UserRequest   true    "prompt"  example({
//	 "model": "mistral-small-latest",
//	 "temperature": 1.5,
//	 "top_p": 1,
//	 "max_tokens": 0,
//	 "stream": false,
//	 "stop": "string",
//	 "random_seed": 0,
//	 "messages": [
//	   {
//	     "role": "user",
//	     "content": "Who is the best French painter? Answer in one short sentence."
//	   }
//	 ],
//	 "response_format": {
//	   "type": "text"
//	 },
//	 "tools": [
//	   {
//	     "type": "function",
//	     "function": {
//	       "name": "string",
//	       "description": "",
//	       "parameters": {}
//	     }
//	   }
//	 ],
//	 "tool_choice": "auto",
//	 "presence_penalty": 0,
//	 "frequency_penalty": 0,
//	 "n": 1,
//	 "safe_prompt": false
//	})
//
// @Success        200     {object}  map[string]interface{}
// @Router         /api/prompt [post]
func PostPrompt(ctx *fiber.Ctx) error {
	var prompt T.UserRequest

	if err := ctx.BodyParser(&prompt); err != nil {
		return H.BuildError(ctx, "Invalid JSON", fiber.StatusBadRequest, err)
	}

	response, serviceErr := S.UseModel(&prompt)

	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	return H.Success(ctx, fiber.Map{
		"response": response.Choices[0].Message.Content,
	})
}

// DocumentPrompt
// @Summary        DocumentPrompt
// @Description    Post prompt with image of document
// @Tags           Prompts
// @Accept         json
// @Produce        json
//
//	@Param          prompt  body    T.UserRequestImage   true    "prompt"
//
// @Success        200     {object}  map[string]interface{}
// @Router         /api/documentPrompt [post]
func DocumentPrompt(ctx *fiber.Ctx) error {
	var prompt T.UserRequestImage

	if err := ctx.BodyParser(&prompt); err != nil {
		return H.BuildError(ctx, "Invalid JSON", fiber.StatusBadRequest, err)
	}

	response, serviceErr := S.UseModelImage(&prompt)

	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	return H.Success(ctx, fiber.Map{
		"response": response.Choices[0].Message.Content,
	})
}
