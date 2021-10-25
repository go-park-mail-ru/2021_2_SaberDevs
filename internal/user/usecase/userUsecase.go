package usecases

import (
	"context"
	"net/http"
	"net/mail"
	"regexp"
	"strconv"
	"strings"

	smodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/models"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	umodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
	"github.com/pkg/errors"
	emoji "github.com/tmdvs/Go-Emoji-Utils"
)

type userUsecase struct {
	userRepo    umodels.UserRepository
	sessionRepo smodels.SessionRepository
}

func NewUserUsecase(ur umodels.UserRepository, sr smodels.SessionRepository) umodels.UserUsecase {
	return &userUsecase{
		userRepo:    ur,
		sessionRepo: sr,
	}
}

// TODO error handling
func (uu *userUsecase) LoginUser(ctx context.Context, user *umodels.User) (umodels.LoginResponse, string, error) {
	var response umodels.LoginResponse

	userInRepo, err := uu.userRepo.GetByEmail(ctx, user.Email)
	if err != nil {
		return response, "", errors.Wrap(err, "userUsecase/LoginUser")
	}

	if userInRepo.Password != user.Password {
		return response, "", sbErr.ErrWrongPassword{
			Reason:   "wrong password",
			Function: "userUsecase/LoginUser"}
	}

	sessionID, err := uu.sessionRepo.CreateSession(ctx, user.Email)
	if err != nil {
		return response, "", errors.Wrap(err, "userUsecase/LoginUser")
	}

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

	return response, sessionID, nil
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

func (uu *userUsecase) Signup(ctx context.Context, user *umodels.User) (umodels.SignupResponse, string, error) {
	var response umodels.SignupResponse

	signedupUser, err := uu.userRepo.Store(ctx, user)
	if err != nil {
		return response, "", errors.Wrap(err, "userUsecase/Signup")
	}

	sessionID, err := uu.sessionRepo.CreateSession(ctx, user.Email)
	if err != nil {
		return response, "", errors.Wrap(err, "userUsecase/Signup")
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

	return response, sessionID, nil
}

func (uu *userUsecase) Logout(ctx context.Context, cookieValue string) error {
	err := uu.sessionRepo.DeleteSession(ctx, cookieValue)
	return errors.Wrap(err, "userUsecase/Logout")
}
