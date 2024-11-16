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
		"ok":       1,
		"response": response,
	})
}

//func CallDima(prompt string) (string, *T.ServiceError) {
//	req := T.Prompt{
//		ActId: 0,
//		Message: []T.Message{{
//			IsAI:    false,
//			Content: prompt,
//		},
//		},
//	}
//
//	responseChapter, serviceErr := S.GetAnswered(req)
//	if serviceErr != nil {
//		return "", serviceErr
//	}
//
//	var content_1 federalChapter
//
//	err := json.Unmarshal([]byte(responseChapter), &content_1)
//	if err != nil {
//		return "", &T.ServiceError{
//			Message: "O_o",
//			Error:   err,
//			Code:    0,
//		}
//	}
//
//	return content_1.Choices[0].Message.Content, nil
//}
//
//func MakeApiRequest(answer T.ApiResponse) *T.ApiResponse {
//	return &T.ApiResponse{
//		Answer:      answer.Choices[0].Message.Content, // first time
//		ChapterName: answer.Choices[0].Message.Content, // second time
//		ArticleName: answer.Choices[0].Message.Content, // third time
//	}
//}

//func PostPromt(ctx *fiber.Ctx) error {
//	var prompt T.Prompt
//
//	if err := ctx.BodyParser(&prompt); err != nil {
//		return H.BuildError(ctx, "Invalid JSON", fiber.StatusBadRequest, err)
//	}
//
//	response, serviceErr := S.GetAnswered(prompt)
//	if serviceErr != nil {
//		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
//	}
//
//	var content federalChapter
//
//	err := json.Unmarshal([]byte(response), &content)
//	if err != nil {
//		return err
//	}
//
//	var apiResponse = MakeApiRequest(content)
//
//	return H.Success(ctx, fiber.Map{
//		"ok":             1,
//		"federalChapter": apiResponse,
//	})
//}
//
//func CallDima(prompt string) (string, *T.ServiceError) {
//	req := T.Prompt{
//		ActId: 0,
//		Message: []T.Message{{
//			IsAI:    false,
//			Content: prompt,
//		},
//		},
//	}
//
//	responseChapter, serviceErr := S.GetAnswered(req)
//	if serviceErr != nil {
//		return "", serviceErr
//	}
//
//	var content_1 federalChapter
//
//	err := json.Unmarshal([]byte(responseChapter), &content_1)
//	if err != nil {
//		return "", &T.ServiceError{
//			Message: "O_o",
//			Error:   err,
//			Code:    0,
//		}
//	}
//
//	return content_1.Choices[0].Message.Content, nil
//}
//
//// FindFirstNumberInString находит первое число в строке и возвращает его в виде int
//func FindFirstNumberInString(s string) (int, error) {
//	// Регулярное выражение для поиска чисел
//	re := regexp.MustCompile(`\d+`)
//	match := re.FindString(s)
//	if match == "" {
//		return 0, fmt.Errorf("no number found in string")
//	}
//	// Преобразование строки в int
//	number, err := strconv.Atoi(match)
//	if err != nil {
//		return 0, fmt.Errorf("failed to convert string to int: %v", err)
//	}
//	return number, nil
//}
//

//func PingPong(ctx *fiber.Ctx) error {
//	var prompt T.Prompt
//
//	if err := ctx.BodyParser(&prompt); err != nil {
//		return H.BuildError(ctx, "Invalid JSON", fiber.StatusBadRequest, err)
//	}
//
//	var mainFile = db.Init_db()
//
//	var fedText = ""
//
//	for chapterIdx := 0; chapterIdx < len(mainFile); chapterIdx++ {
//		fedText += "ID " + strconv.Itoa(chapterIdx+1) + " " + mainFile[chapterIdx].Title + " " + mainFile[chapterIdx].Description + "\n"
//	}
//
//	origPrompt := prompt.Message[len(prompt.Message)-1].Content
//
//	//var second T.Prompt
//	//var third T.Prompt
//	promptMsg := fedText + "\n\n" + "Выше представлены главы законодательного акта РФ о связи. " +
//		"Напиши только ID к которому относится следующий вопрос: \"" + origPrompt + "\"" + "\n" +
//		"Ответ должен быть представлен только числом"
//
//	aa, err := CallDima(promptMsg)
//	if err != nil {
//		return H.BuildError(ctx, err.Message, err.Code, err.Error)
//	}
//
//	chapterNumber, _ := FindFirstNumberInString(aa)
//	chapterIndex := chapterNumber - 1
//	if chapterIndex < 0 {
//		return H.BuildError(ctx, "Отрицательный ID главы", 0, nil)
//	}
//
//	chapterName := mainFile[chapterIndex].Title
//
//	if chapterIndex < 0 || chapterIndex >= 13 {
//		return H.BuildError(ctx, "Попытка выйти за пределы глав "+strconv.Itoa(chapterIndex), 0, nil)
//	}
//
//	fedText = ""
//
//	for articleIdx := 0; articleIdx < len(mainFile[chapterIndex].Articles); articleIdx++ {
//		fedText += "ID " + strconv.Itoa(articleIdx) + ". " + mainFile[chapterIndex].Articles[articleIdx].Title +
//			" " + mainFile[chapterIndex].Articles[articleIdx].ShortDescription + "\n"
//	}
//
//	promptMsg = fedText + "\n\n" + "Выше представлены статьи законодательного акта РФ о связи. " +
//		"Напиши только ID к которому относится следующий вопрос: \"" + origPrompt + "\"" + "\n" +
//		"Ответ должен быть представлен только числом"
//
//	bb, _ := CallDima(promptMsg)
//
//	articleIndex, _ := FindFirstNumberInString(bb)
//
//	if articleIndex < 0 || articleIndex >= len(mainFile[chapterIndex].Articles) {
//		return H.BuildError(ctx, "Попытка выйти за пределы статей "+strconv.Itoa(articleIndex), 0, nil)
//	}
//
//	articleText := mainFile[chapterIndex].Articles[articleIndex].OriginalText
//	articleName := mainFile[chapterIndex].Articles[articleIndex].Title
//	articlePage := mainFile[chapterIndex].Articles[articleIndex].StartPageNumber
//
//	fedText = articleText + "\n\n" +
//		"Выше представлена статья законодательного акта РФ о связи." +
//		"Ответь используя эту статью на вопрос: \"" + origPrompt + "\"" + "\n" +
//		"Ответ также должен содержать цитату с отрывком отвечающим на вопрос\n" +
//		"Ответ не должен содержать лишних домыслов, только то что написано в представленном тексте\n" +
//		"Сначала кратко ответь на вопрос в 2 предложения, потом приведи цитату из текста"
//
//	final, _ := CallDima(fedText)
//
//	_ = final
//
//	resp := ApiResponse{
//		Answer:      final,       // first time
//		ChapterName: chapterName, // second time
//		ArticleName: articleName, // third time
//		PageNumber:  articlePage,
//	}
//
//	return H.Success(ctx, fiber.Map{
//		"ok":             1,
//		"federalChapter": resp,
//	})
//}
