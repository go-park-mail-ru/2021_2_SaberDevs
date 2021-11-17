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

const passwordInvalidMsg = `пароль должен быть от 4 до 40 символов английского алфавита, символы !"#$&'()*+,\-./:;<=>?@[\]^_{|}~[] и цифры`
const nameInvalidMsg = "имя должно быть от 4 до 20 символов русского, английского алфавита и цифры"
const surnameInvalidMsg = "фамилия должна быть от 4 до 20 символов русского, английского алфавита и цифры"
const loginInvalidMsg = "логин должен быть от 4 до 20 символов английского алфавита и цифры"

func ValidateSignUp(user umodels.User) error {
	err := validation.ValidateStruct(&user,
		validation.Field(&user.Login, validation.Required.Error("Логин это обязательное поле"),
			validation.Match(regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_]{4,20}$")).Error(loginInvalidMsg)),
		validation.Field(&user.Email, validation.Required.Error("Email это обязательное поле"), is.EmailFormat.Error("Неверный email")),
		validation.Field(&user.Password, validation.Required.Error("Пароль это обязательное поле"), validation.By(isPasswordValid)),
	)
	if err != nil {
		return err
	}

	return nil
}

func ValidateUpdate(user umodels.User) error {
	err := validation.ValidateStruct(&user,
		validation.Field(&user.Name, validation.When(user.Name != "",
			validation.Match(regexp.MustCompile("[a-zA-Zа-яА-ЯЁё][a-zA-Z_а-яА-ЯЁё'-]{4,20}$")).Error(nameInvalidMsg))),
		validation.Field(&user.Surname, validation.When(user.Surname != "",
			validation.Match(regexp.MustCompile("[a-zA-Zа-яА-ЯЁё][a-zA-Z_а-яА-ЯЁё'-]{4,20}$")).Error(surnameInvalidMsg))),
		validation.Field(&user.Password, validation.When(user.Password != "", validation.By(isPasswordValid))),
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
	// password length: min 8, max 40
	validator = regexp.MustCompile("^[a-zA-Z0-9[:punct:]]{" + strconv.Itoa(minPasswordLength) + ",40}$")
	if !validator.MatchString(inputWithoutEmoji) {
		return errors.New(passwordInvalidMsg)
	}
	return nil
}
