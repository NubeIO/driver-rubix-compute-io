package common

import "github.com/gin-gonic/gin"

type Message struct {
	Body interface{} `json:"body"`
	Err  string      `json:"error"`
}

func ReposeHandler(body interface{}, err error, ctx *gin.Context) {
	if err != nil {
		ctx.JSON(404, Message{Err: err.Error()})
	} else {
		ctx.JSON(202, body)
	}
}
