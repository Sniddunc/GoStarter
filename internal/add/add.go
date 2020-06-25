package add

import (
	"GoStarter/internal/common"
	"GoStarter/pkg/app"
	"net/http"
	"net/url"
	"strconv"
)

// Handler is the route handler function for GET /api/v1/add
func Handler(c *app.Context) {
	num1str := c.Query("num1", "0")
	num2str := c.Query("num2", "0")

	num1, err := strconv.ParseInt(num1str, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.StandardResponse{
			Success: false,
			Message: "Could not effectuate addition",
			Errors: url.Values{
				"num1": []string{"num1 was not a valid int"},
			},
		})

		return
	}

	num2, err := strconv.ParseInt(num2str, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.StandardResponse{
			Success: false,
			Message: "Could not effectuate addition",
			Errors: url.Values{
				"num2": []string{"num2 was not a valid int"},
			},
		})

		return
	}

	c.JSON(http.StatusOK, common.StandardResponse{
		Success: true,
		Message: "Addition completed",
		Payload: num1 + num2,
	})
}
