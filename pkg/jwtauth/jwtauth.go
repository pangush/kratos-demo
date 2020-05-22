// Package jwtauth implements functions for
package jwtauth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"strings"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
)

//
var (
	ErrInvalidToken      = errors.New("Invalid token")
	ErrSigningMethodHMAC = errors.New("jwt SigningMethodHMAC")
	ErrInvalidateToken   = errors.New("Invalidate token")
	ErrRandRead          = errors.New("rand.Read Err")
)

// Config jwt settings.
type Config struct {
	Algo   string // 签名算法(支持：HS512/HS384/HS512)
	TTL    int64  // 有效时长(minutes)
	Secret string // 密钥
}

// Storer interface
type Storer interface {
	// Invalidate a token (add it to the blacklist).
	Invalidate(ctx context.Context, jti string, ttl int64) error
	// Check token is Invalidate
	IsInvalidate(ctx context.Context, jti string) (bool, error)
}

// Jwtauth struct
type Jwtauth struct {
	conf   *Config
	method jwtgo.SigningMethod
	store  Storer
}

// New Jwt returns a new Jwtauth
func New(conf *Config, store Storer) (*Jwtauth, error) {
	jwt := &Jwtauth{conf: conf, store: store}
	jwt.setSigningMethod()

	return jwt, nil
}

func (jwt *Jwtauth) setSigningMethod() {
	switch jwt.conf.Algo {
	case "HS256":
		jwt.method = jwtgo.SigningMethodHS256
	case "HS384":
		jwt.method = jwtgo.SigningMethodHS384
	case "HS512":
		jwt.method = jwtgo.SigningMethodHS512
	default:
		log.Fatalf("unknow jwt Algo: %v", jwt.conf.Algo)
	}
}

// GenerateToken 生成token
func (jwt *Jwtauth) GenerateToken(ctx context.Context, userID string) (string, error) {
	now := time.Now()
	expiresAt := now.Add(time.Duration(jwt.conf.TTL) * time.Second).Unix()

	jtiBytes := make([]byte, 16)
	_, err := rand.Read(jtiBytes)
	if err != nil {
		return "", ErrRandRead
	}
	jti := base64.StdEncoding.EncodeToString(jtiBytes)
	jti = strings.Replace(jti, "/", "", -1)
	jti = strings.Replace(jti, "+", "", -1)
	jti = strings.Replace(jti, "=", "", -1)

	token := jwtgo.NewWithClaims(jwt.method, &jwtgo.StandardClaims{
		Id:        jti,
		IssuedAt:  now.Unix(),
		ExpiresAt: expiresAt,
		NotBefore: now.Unix(),
		Subject:   userID,
	})

	tokenString, err := token.SignedString([]byte(jwt.conf.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GetUserID returns a single string of userID
func (jwt *Jwtauth) GetUserID(ctx context.Context, tokenString string) (string, error) {
	claims, err := jwt.parseToken(ctx, tokenString)
	if err != nil {
		return "", err
	}

	return claims.Subject, nil
}

// parseToken 解析token
func (jwt *Jwtauth) parseToken(ctx context.Context, tokenString string) (*jwtgo.StandardClaims, error) {
	token, err := jwtgo.ParseWithClaims(tokenString, &jwtgo.StandardClaims{}, func(t *jwtgo.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwtgo.SigningMethodHMAC); !ok {
			return nil, ErrSigningMethodHMAC
		}
		return []byte(jwt.conf.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwtgo.StandardClaims)

	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	IsInvalidate, err := jwt.store.IsInvalidate(ctx, claims.Id)

	if IsInvalidate {
		return nil, ErrInvalidateToken
	}

	return claims, nil
}

// ParseUserID returns a single string of userID
func (jwt *Jwtauth) ParseUserID(ctx context.Context, tokenString string) (string, error) {
	jwtStandardClaims, err := jwt.parseToken(ctx, tokenString)
	if err != nil {
		return "", err
	}

	return jwtStandardClaims.Subject, nil
}

// DestroyToken 销毁token
func (jwt *Jwtauth) DestroyToken(ctx context.Context, tokenString string) error {

	claims, err := jwt.parseToken(ctx, tokenString)
	if err != nil {
		return err
	}

	// 如果设定了存储，则将未过期的令牌放入
	return jwt.callStore(func(store Storer) error {
		expired := int64(time.Unix(claims.ExpiresAt, 0).Sub(time.Now()).Seconds())
		return store.Invalidate(ctx, claims.Id, expired)
	})
}

func (jwt *Jwtauth) callStore(fn func(Storer) error) error {
	if store := jwt.store; store != nil {
		return fn(store)
	}
	return nil
}
