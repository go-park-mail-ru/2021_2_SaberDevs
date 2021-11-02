package syberValidation

import (
	"regexp"
	"strconv"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	umodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
	"github.com/pkg/errors"
	emoji "github.com/tmdvs/Go-Emoji-Utils"
)

func ValidateSignUp(user umodels.User) error {
	err := validation.ValidateStruct(&user,
		validation.Field(&user.Login, validation.Required, validation.Length(4, 20),
			validation.Match(regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_]{4,20}$"))),
		validation.Field(&user.Email, validation.Required, is.Email, validation.Length(4, 40)),
		validation.Field(&user.Password, validation.Required, validation.By(isPasswordValid)),
	)
	if err != nil {
		return err
	}

	return nil
}

func ValidateUpdate(user umodels.User) error {
	err := validation.ValidateStruct(&user,
		validation.Field(&user.Name, validation.When(user.Name != "", validation.Length(4, 20),
			validation.Match(regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_]{4,20}$"))).Else(validation.Nil)),
		validation.Field(&user.Surname, validation.When(user.Surname != "",validation.Length(4, 20),
			validation.Match(regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_]{4,20}$"))).Else(validation.Nil)),
		validation.Field(&user.Password, validation.When(user.Password != "", validation.By(isPasswordValid)).Else(validation.Nil)),
	)
	if err != nil {
		return err
	}
	return nil
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
