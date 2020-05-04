package helpers


import (
"fmt"
"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Status int
	Data   interface{}
}

func RespondJSON(w *gin.Context, status int, payload interface{}) {
	fmt.Println("status ", status)
	var res ResponseData

	res.Status = status
	res.Data = payload

	w.JSON(200, res)
}