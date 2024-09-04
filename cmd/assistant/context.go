package main

type Context struct {
	Messages   []Message
	Questions  []string
	Responses  []string
	Screenshot string
	// NOTE: initially, im thinking to merge Questions and Responses
	// try to mimic hows ChatGPT messages are structured {"user", "content"}
	// chronologically, more straighforward.
	// but I think its easier to split them for now
}

func (c *Context) AddMessage(message Message) {
	c.Messages = append(c.Messages, message)
}

func (c *Context) AddQuestionAndResponse(question, response string) {
	c.Questions = append(c.Questions, question)
	c.Responses = append(c.Responses, response)
}

func (c *Context) SetScreenshot(screenshot string) {
	c.Screenshot = screenshot
}
