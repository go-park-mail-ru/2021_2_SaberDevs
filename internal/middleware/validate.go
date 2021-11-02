package middleware

import (
	errResp "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/errResponses"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
	"github.com/labstack/echo/v4"
	emoji "github.com/tmdvs/Go-Emoji-Utils"
)

func ValidateRequestBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		switch c.Path() {
		case "/api/v1/user/signup":
			var user models.User

			err := c.Bind(&user)
			if err != nil {
				return sbErr.ErrUnpackingJSON{
					Reason:   err.Error(),
					Function: "validation/signup",
				}
			}

			err = validation.ValidateStruct(&user,
				validation.Field(&user.Login, validation.Required, validation.Length(4, 20),
					validation.Match(regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_]{4,20}$"))),
				validation.Field(&user.Email, validation.Required, is.Email, validation.Length(4, 40)),
				validation.Field(&user.Password, validation.Required, validation.By(isPasswordValid)),
			)
			if err != nil {
				return c.JSON(http.StatusFailedDependency, errResp.ErrorResponse{
					Status:   http.StatusFailedDependency,
					ErrorMsg: err.Error()})
			}

		case "/api/v1/user/profile/update":
			var user models.User

			err := c.Bind(&user)
			if err != nil {
				return sbErr.ErrUnpackingJSON{
					Reason:   err.Error(),
					Function: "validation/update",
				}
			}

			err = validation.ValidateStruct(&user,
				validation.Field(&user.Name, validation.Length(4, 20),
					validation.Match(regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_]{4,20}$"))),
				validation.Field(&user.Surname, validation.Length(4, 20),
					validation.Match(regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_]{4,20}$"))),
				validation.Field(&user.Password, validation.By(isPasswordValid)),
			)
			if err != nil {
				return c.JSON(http.StatusFailedDependency, errResp.ErrorResponse{
					Status:   http.StatusFailedDependency,
					ErrorMsg: err.Error()})
			}
		}

		return next(c)
	}
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

func isPasswordValid(input interface{}) error {
	inputWithoutEmoji, emoCount := removeAllAndCount(input.(string))
	var validator *regexp.Regexp
	minPasswordLength := minPasswordLength(emoCount)
	validator = regexp.MustCompile("^[a-zA-Z0-9[:punct:]]{" + strconv.Itoa(minPasswordLength) + ",20}$")
	if !validator.MatchString(inputWithoutEmoji) {
		return errors.New("invalid symbols in password")
	}
	return nil
}
