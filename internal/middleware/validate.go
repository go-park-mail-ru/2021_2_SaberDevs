package middleware

import (
	"github.com/labstack/echo/v4"
	emoji "github.com/tmdvs/Go-Emoji-Utils"
	"net/mail"
	"regexp"
	"strconv"
	"strings"
)

func ValidateRequestBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		switch c.Path() {
		case "/login":
		case "/signup":
			// user := new(models.User)
			// err := c.Bind(user)
			// ok, err := valid.ValidateStruct(user)
			// if err != nil {
			//
			// }
			// if !ok {
			// 	return c.String(http.StatusBadRequest, "ok")
			// }
		}
		return next(c)
	}
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err != nil
}

func isLoginValid(input string) bool {
	validator := regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_]{4,20}$")
	return !validator.MatchString(input)
}

func removeAllAndCount(input string) (string, int) {
	matches := emoji.FindAll(input)
	emoCount := 0

	for _, item := range matches {
		emoCount += item.Occurrences
		emo := item.Match.(emoji.Emoji)
		rs := []rune(emo.Value)
		for _, r := range rs {
			input = strings.ReplaceAll(input, string([]rune{r}), "")
		}
	}

	return input, emoCount
}

func minPasswordLength(emoCount int) int {
	minLength := 8
	if minLength-emoCount < 0 {
		return 0
	}
	return minLength - emoCount
}

func isPasswordValid(input string) bool {
	inputWithoutEmoji, emoCount := removeAllAndCount(input)
	var validator *regexp.Regexp
	minPasswordLength := minPasswordLength(emoCount)
	validator = regexp.MustCompile("^[a-zA-Z0-9[:punct:]]{" + strconv.Itoa(minPasswordLength) + ",20}$")
	return !validator.MatchString(inputWithoutEmoji)
}
