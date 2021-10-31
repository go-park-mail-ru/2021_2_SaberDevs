package usecases

import (
	"context"
	"errors"
	smodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/models"
	umodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
	emoji "github.com/tmdvs/Go-Emoji-Utils"
	"net/http"
	"net/mail"
	"regexp"
	"strconv"
	"strings"
)

type userUsecase struct {
	userRepo umodels.UserRepository
	sessionRepo smodels.SessionRepository
}

func NewUserUsecase(ur umodels.UserRepository, sr smodels.SessionRepository) umodels.UserUsecase {
	return &userUsecase{
		userRepo: ur,
		sessionRepo: sr,
	}
}

// TODO error handling
func (uu *userUsecase) LoginUser(ctx context.Context, user *umodels.User) (umodels.LoginResponse, string, error) {
	var response umodels.LoginResponse
	userInRepo, err := uu.userRepo.GetByEmail(ctx, user.Email)
	if err != nil {
		// TODO user doesnt exist err
		return response, "", err
	}

	if userInRepo.Password != user.Password {
		var err = errors.New("wrong password")
		return response, "", err
	}

	cookieValue, err := uu.sessionRepo.CreateSession(ctx, user.Email)

	d := umodels.LoginData{
		Login:   userInRepo.Login,
		Name:    userInRepo.Name,
		Surname: userInRepo.Surname,
		Email:   userInRepo.Email,
		Score:   userInRepo.Score,
	}
	response = umodels.LoginResponse{
		Status: http.StatusOK,
		Data:   d,
		Msg:    "OK",
	}

	return response, cookieValue, nil
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err != nil
}

func isLoginValid(input string) bool {
	var validator *regexp.Regexp
	validator = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_]{4,20}$")
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

func (uu *userUsecase) Signup(ctx context.Context, user *umodels.User) (umodels.SignupResponse, string, error) {
	var response umodels.SignupResponse

	signedupUser, err := uu.userRepo.Store(ctx, user)
	if err != nil {
		// TODO error
		return response, "", err
	}

	d := umodels.SignUpData{
		Login:   signedupUser.Login,
		Name:    signedupUser.Name,
		Surname: signedupUser.Surname,
		Email:   signedupUser.Email,
		Score:   signedupUser.Score,
	}
	response = umodels.SignupResponse{
		Status: http.StatusOK,
		Data:   d,
		Msg:    "OK",
	}

	cookieValue, err := uu.sessionRepo.CreateSession(ctx, user.Email)
	if err != nil {
		// TODO error
		return response, "", err
	}

	return response, cookieValue, nil
}

func (uu *userUsecase) Logout(ctx context.Context, cookieValue string) error {
	err := uu.sessionRepo.DeleteSession(ctx, cookieValue)
	return err
}
