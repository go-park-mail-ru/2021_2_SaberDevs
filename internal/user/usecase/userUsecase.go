package usecases

import (
	"context"
	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
	"google.golang.org/grpc/status"
	"net/http"

	kmodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/keys/models"

	smodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/models"
	umodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
	"github.com/pkg/errors"
)

type userUsecase struct {
	userRepo    umodels.UserRepository
	sessionRepo smodels.SessionRepository
	keyRepo     kmodels.KeyRepository
	articleRepo amodels.ArticleRepository
}

func NewUserUsecase(ur umodels.UserRepository, sr smodels.SessionRepository, kr kmodels.KeyRepository, ar amodels.ArticleRepository) umodels.UserUsecase {
	return &userUsecase{
		userRepo:    ur,
		sessionRepo: sr,
		keyRepo:     kr,
		articleRepo: ar,
	}
}

func (uu *userUsecase) GetAuthorProfile(ctx context.Context, author string) (umodels.GetUserResponse, error) {
	authorInDb, err := uu.userRepo.GetByLogin(ctx, author)
	if err != nil {
		return umodels.GetUserResponse{}, errors.Wrap(err, "userUsecase/GetAuthorProfile")
	}

	responseData := umodels.GetUserData{
		Login:       authorInDb.Login,
		Name:        authorInDb.Name,
		Surname:     authorInDb.Surname,
		Score:       authorInDb.Score,
		AvatarURL:   authorInDb.AvatarURL,
		Description: authorInDb.Description,
	}
	response := umodels.GetUserResponse{
		Status: http.StatusOK,
		Data:   responseData,
		Msg:    "ok",
	}

	return response, nil
}

func (uu *userUsecase) GetUserProfile(ctx context.Context, sessionID string) (umodels.GetUserResponse, error) {
	userLogin, err := uu.sessionRepo.GetSessionLogin(ctx, sessionID)
	if err != nil {
		return umodels.GetUserResponse{}, errors.Wrap(err, "userUsecase/UpdateProfile")
	}

	userInDb, err := uu.userRepo.GetByLogin(ctx, userLogin)
	if err != nil {
		return umodels.GetUserResponse{}, errors.Wrap(err, "userUsecase/UpdateProfile")
	}

	responseData := umodels.GetUserData{
		Login:       userInDb.Login,
		Name:        userInDb.Name,
		Surname:     userInDb.Surname,
		Score:       userInDb.Score,
		AvatarURL:   userInDb.AvatarURL,
		Description: userInDb.Description,
	}
	response := umodels.GetUserResponse{
		Status: http.StatusOK,
		Data:   responseData,
		Msg:    "ok",
	}

	return response, nil
}

func (uu *userUsecase) UpdateProfile(ctx context.Context, user *umodels.User, sessionID string) (umodels.LoginResponse, error) {
	var response umodels.LoginResponse

	login, err := uu.sessionRepo.GetSessionLogin(ctx, sessionID)
	if err != nil {
		return response, errors.Wrap(err, "userUsecase/UpdateProfile")
	}

	user.Login = login

	_, err = uu.userRepo.UpdateUser(ctx, user)
	if err != nil {
		return response, errors.Wrap(err, "userUsecase/UpdateProfile")
	}

	updatedUser, err := uu.userRepo.GetByLogin(ctx, user.Login)
	if err != nil {
		return response, errors.Wrap(err, "userUsecase/UpdateProfile")
	}

	responseData := umodels.LoginData{
		Login:       updatedUser.Login,
		Name:        updatedUser.Name,
		Surname:     updatedUser.Surname,
		Email:       updatedUser.Email,
		Score:       updatedUser.Score,
		AvatarURL:   updatedUser.AvatarURL,
		Description: updatedUser.Description,
	}
	response = umodels.LoginResponse{
		Status: http.StatusOK,
		Data:   responseData,
		Msg:    "OK",
	}

	return response, nil
}

func (uu *userUsecase) LoginUser(ctx context.Context, user *umodels.User) (umodels.LoginResponse, string, error) {
	var response umodels.LoginResponse

	userInRepo, err := uu.userRepo.GetByLogin(ctx, user.Login)
	if err != nil {
		return response, "", errors.Wrap(err, "userUsecase/LoginUser")
	}

	if userInRepo.Password != user.Password {
		// return response, "", sbErr.ErrWrongPassword{
		// 	Reason:   "Неверный пароль",
		// 	Function: "userUsecase/LoginUser"}
		return response, "", status.Error(18, "Неверный пароль")
	}

	sessionID, err := uu.sessionRepo.CreateSession(ctx, user.Login)
	if err != nil {
		return response, "", errors.Wrap(err, "userUsecase/LoginUser")
	}

	responseData := umodels.LoginData{
		Login:       userInRepo.Login,
		Name:        userInRepo.Name,
		Surname:     userInRepo.Surname,
		Email:       userInRepo.Email,
		Score:       userInRepo.Score,
		AvatarURL:   userInRepo.AvatarURL,
		Description: userInRepo.Description,
	}
	response = umodels.LoginResponse{
		Status: http.StatusOK,
		Data:   responseData,
		Msg:    "OK",
	}

	return response, sessionID, nil
}

func (uu *userUsecase) Signup(ctx context.Context, user *umodels.User) (umodels.SignupResponse, string, error) {
	var response umodels.SignupResponse

	signedupUser, err := uu.userRepo.Store(ctx, user)
	if err != nil {
		// return response, "", errors.Wrap(err, "userUsecase/Signup")
		return response, "", err
	}

	sessionID, err := uu.sessionRepo.CreateSession(ctx, user.Login)
	if err != nil {
		return response, "", errors.Wrap(err, "userUsecase/Signup")
	}

	responseData := umodels.SignUpData{
		Login:       signedupUser.Login,
		Name:        signedupUser.Name,
		Surname:     signedupUser.Surname,
		Email:       signedupUser.Email,
		Score:       signedupUser.Score,
		AvatarURL:   signedupUser.AvatarURL,
		Description: signedupUser.Description,
	}
	response = umodels.SignupResponse{
		Status: http.StatusOK,
		Data:   responseData,
		Msg:    "OK",
	}

	return response, sessionID, nil
}

func (uu *userUsecase) Logout(ctx context.Context, cookieValue string) error {
	err := uu.sessionRepo.DeleteSession(ctx, cookieValue)
	return errors.Wrap(err, "userUsecase/Logout")
}
