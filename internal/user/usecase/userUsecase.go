package usecases

import (
	"context"
	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
	"net/http"

	"bytes"
	"crypto/rand"

	"golang.org/x/crypto/argon2"

	kmodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/keys/models"

	smodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/models"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
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
		Login:   authorInDb.Login,
		Name:    authorInDb.Name,
		Surname: authorInDb.Surname,
		Score:   authorInDb.Score,
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
		Login:   userInDb.Login,
		Name:    userInDb.Name,
		Surname: userInDb.Surname,
		Score:   userInDb.Score,
	}
	response := umodels.GetUserResponse{
		Status: http.StatusOK,
		Data:   responseData,
		Msg:    "ok",
	}

	return response, nil
}

func (uu *userUsecase) UpdateProfile(ctx context.Context, user *umodels.User, sessionID string) (umodels.UpdateProfileResponse, error) {
	var response umodels.UpdateProfileResponse

	login, err := uu.sessionRepo.GetSessionLogin(ctx, sessionID)
	if err != nil {
		return response, errors.Wrap(err, "userUsecase/UpdateProfile")
	}

	user.Login = login

	updatedUser, err := uu.userRepo.UpdateUser(ctx, user)
	if err != nil {
		return response, errors.Wrap(err, "userUsecase/UpdateProfile")
	}

	responseData := umodels.UpdateProfileData{
		Name:    updatedUser.Name,
		Surname: updatedUser.Surname,
	}
	response = umodels.UpdateProfileResponse{
		Status: http.StatusOK,
		Data:   responseData,
		Msg:    "OK",
	}

	return response, nil
}

func checkPass(saltedPass []byte, plainPass string, salt string) bool {
	saltedPlainPass := saltPass([]byte(salt), plainPass)
	return bytes.Equal(saltedPlainPass, saltedPass)
}

func (uu *userUsecase) LoginUser(ctx context.Context, user *umodels.User) (umodels.LoginResponse, string, error) {
	var response umodels.LoginResponse

	userInRepo, err := uu.userRepo.GetByLogin(ctx, user.Login)
	if err != nil {
		return response, "", errors.Wrap(err, "userUsecase/LoginUser")
	}

	salt, err := uu.keyRepo.GetSalt(ctx, user.Login)
	if err != nil {
		return response, "", errors.Wrap(err, "userUsecase/LoginUser")
	}

	if !checkPass([]byte(userInRepo.Password), user.Password, salt) {
		return response, "", sbErr.ErrWrongPassword{
			Reason:   "wrong password",
			Function: "userUsecase/LoginUser"}
	}

	sessionID, err := uu.sessionRepo.CreateSession(ctx, user.Login)
	if err != nil {
		return response, "", errors.Wrap(err, "userUsecase/LoginUser")
	}

	responseData := umodels.LoginData{
		Login:   userInRepo.Login,
		Name:    userInRepo.Name,
		Surname: userInRepo.Surname,
		Email:   userInRepo.Email,
		Score:   userInRepo.Score,
	}
	response = umodels.LoginResponse{
		Status: http.StatusOK,
		Data:   responseData,
		Msg:    "OK",
	}

	return response, sessionID, nil
}

func saltPass(salt []byte, plainPassword string) []byte {
	saltedPass := argon2.IDKey([]byte(plainPassword), salt, 1, 64*1024, 4, 32)
	return saltedPass
}

func makeSalt() ([]byte, error) {
	salt := make([]byte, 8)

	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}

	return salt, nil
}

func (uu *userUsecase) Signup(ctx context.Context, user *umodels.User) (umodels.SignupResponse, string, error) {
	var response umodels.SignupResponse

	salt, err := makeSalt()
	if err != nil {
		return response, "", sbErr.ErrInternal{
			Reason:  err.Error(),
			Function: "userUsecase/Signup",
		}
	}

	user.Password = string(saltPass(salt, user.Password))

	err = uu.keyRepo.StoreSalt(ctx, kmodels.Key{
		Salt: string(salt),
		Login: user.Login,
	})
	if err != nil {
		return response, "", errors.Wrap(err, "userUsecase/Signup")
	}

	signedupUser, err := uu.userRepo.Store(ctx, user)
	if err != nil {
		err := uu.keyRepo.DeleteSalt(ctx, user.Login)
		return response, "", errors.Wrap(err, "userUsecase/Signup")
	}

	sessionID, err := uu.sessionRepo.CreateSession(ctx, user.Login)
	if err != nil {
		return response, "", errors.Wrap(err, "userUsecase/Signup")
	}

	responseData := umodels.SignUpData{
		Login:   signedupUser.Login,
		Name:    signedupUser.Name,
		Surname: signedupUser.Surname,
		Email:   signedupUser.Email,
		Score:   signedupUser.Score,
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
