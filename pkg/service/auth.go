package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/daria40tim/packom"
	"github.com/daria40tim/packom/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt      = "ewirufhqp3i94785asdf"
	signinKey = "i4r394po8ngfw3984gnwi983e74bgf47tgasdojf"
	tokenTTL  = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	OId int `json:"o_id"`
}

type AuthService struct {
	repo repository.Autherization
}

func NewAuthService(repo repository.Autherization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateOrg(org packom.Org) (int, error) {
	org.Pwd = generatePasswordHash(org.Pwd)
	return s.repo.CreateOrg(org)
}

func (s *AuthService) GenerateToken(login, pwd string) (string, error, packom.Org) {
	org, err := s.repo.GetOrg(login, generatePasswordHash(pwd))
	if err != nil {
		return "", err, org
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		org.O_id})

	h, error := token.SignedString([]byte(signinKey))

	return h, error, org
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signinKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.OId, nil
}

func (s *AuthService) SelectAllCountries() (packom.Countries, error) {
	return s.repo.SelectAllCountries()
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) SelectLogin(input string) (packom.Countries, error) {
	return s.repo.SelectLogin(input)
}
