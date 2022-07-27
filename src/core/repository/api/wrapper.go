package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func ApiRequest(c *gin.Context) (map[string]interface{}, error) {
	var data map[string]interface{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return data, err
	}
	json.Unmarshal(body, &data)
	return data, err
}

func Success(c *gin.Context, data interface{}) {
	var response = gin.H{
		"data": data,
	}
	c.JSON(http.StatusOK, response)
}

func SuccessRows(c *gin.Context, data interface{}, total int) {
	var response = gin.H{
		"data":         data,
		"totalRecords": total,
	}
	c.JSON(http.StatusOK, response)
}

func Error(c *gin.Context, error string) {
	var response = gin.H{
		"error": error,
	}
	c.JSON(http.StatusBadRequest, response)
}

func ErrorAuth(c *gin.Context, err error) {
	var response = gin.H{
		"error": err.Error(),
	}
	c.JSON(http.StatusInternalServerError, response)
	c.AbortWithError(500, err)
}
