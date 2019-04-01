// Copyright © 2017 Aeneas Rekkas <aeneas+oss@aeneas.io>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package oauth2

import (
	"time"

	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/spotxchange/fosite"
	enigma "github.com/spotxchange/fosite/token/hmac"
)

type HMACSHAStrategy struct {
	Enigma                *enigma.HMACStrategy
	AccessTokenLifespan   time.Duration
	AuthorizeCodeLifespan time.Duration
}

func (h HMACSHAStrategy) AccessTokenSignature(token string) string {
	return h.Enigma.Signature(token)
}
func (h HMACSHAStrategy) RefreshTokenSignature(token string) string {
	return h.Enigma.Signature(token)
}
func (h HMACSHAStrategy) AuthorizeCodeSignature(token string) string {
	return h.Enigma.Signature(token)
}

func (h HMACSHAStrategy) GenerateAccessToken(ctx context.Context, r fosite.Requester) (token string, signature string, err error) {
<<<<<<< HEAD
	return h.Enigma.Generate()
=======
	token, signature, err = h.Enigma.Generate()
	c := context.WithValue(ctx, "refresh_token", token)
	&ctx = &c
	return
>>>>>>> 7fe18edc6513491676d2d7e4279bc4c0da988512
}

func (h HMACSHAStrategy) ValidateAccessToken(_ context.Context, r fosite.Requester, token string) (err error) {
	var exp = r.GetSession().GetExpiresAt(fosite.AccessToken)
	if exp.IsZero() && r.GetRequestedAt().Add(h.AccessTokenLifespan).Before(time.Now().UTC()) {
		return errors.WithStack(fosite.ErrTokenExpired.WithDebug(fmt.Sprintf("Access token expired at %s", r.GetRequestedAt().Add(h.AccessTokenLifespan))))
	}
	if !exp.IsZero() && exp.Before(time.Now().UTC()) {
		return errors.WithStack(fosite.ErrTokenExpired.WithDebug(fmt.Sprintf("Access token expired at %s", exp)))
	}

	if err = h.Enigma.Validate(token); err != nil {
		// So... We know this isn't technically a valid token, but it's also in our DB...
		// Meaning it was migrated in via the api. To make sure we're still being safe(ish),
		// migrated tokens have to be passed in as `[token].[token]`.
		if split := strings.Split(token, "."); len(split) == 2 && split[0] == split[1] {
			return nil
		}
	}
	return
}

func (h HMACSHAStrategy) GenerateRefreshToken(ctx context.Context, r fosite.Requester) (token string, signature string, err error) {
<<<<<<< HEAD
	return h.Enigma.Generate()
=======
	token, signature, err = h.Enigma.Generate()
	c := context.WithValue(ctx, "refresh_token", token)
	&ctx = &c
	return
>>>>>>> 7fe18edc6513491676d2d7e4279bc4c0da988512
}

func (h HMACSHAStrategy) ValidateRefreshToken(_ context.Context, _ fosite.Requester, token string) (err error) {
	return h.Enigma.Validate(token)
}

func (h HMACSHAStrategy) GenerateAuthorizeCode(_ context.Context, _ fosite.Requester) (token string, signature string, err error) {
	return h.Enigma.Generate()
}

func (h HMACSHAStrategy) ValidateAuthorizeCode(_ context.Context, r fosite.Requester, token string) (err error) {
	var exp = r.GetSession().GetExpiresAt(fosite.AuthorizeCode)
	if exp.IsZero() && r.GetRequestedAt().Add(h.AuthorizeCodeLifespan).Before(time.Now().UTC()) {
		return errors.WithStack(fosite.ErrTokenExpired.WithDebug(fmt.Sprintf("Authorize code expired at %s", r.GetRequestedAt().Add(h.AuthorizeCodeLifespan))))
	}
	if !exp.IsZero() && exp.Before(time.Now().UTC()) {
		return errors.WithStack(fosite.ErrTokenExpired.WithDebug(fmt.Sprintf("Authorize code expired at %s", exp)))
	}

	return h.Enigma.Validate(token)
}
