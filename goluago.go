package goluago

import (
	"github.com/Shopify/go-lua"
	"github.com/cardamaro/goluago/pkg/crypto/hmac"
	"github.com/cardamaro/goluago/pkg/encoding/base64"
	"github.com/cardamaro/goluago/pkg/encoding/json"
	"github.com/cardamaro/goluago/pkg/env"
	"github.com/cardamaro/goluago/pkg/fmt"
	"github.com/cardamaro/goluago/pkg/net/url"
	"github.com/cardamaro/goluago/pkg/regexp"
	"github.com/cardamaro/goluago/pkg/strings"
	"github.com/cardamaro/goluago/pkg/time"
	"github.com/cardamaro/goluago/pkg/uuid"
	"github.com/cardamaro/goluago/util"
)

func Open(l *lua.State) {
	regexp.Open(l)
	strings.Open(l)
	json.Open(l)
	time.Open(l)
	fmt.Open(l)
	url.Open(l)
	util.Open(l)
	hmac.Open(l)
	base64.Open(l)
	env.Open(l)
	uuid.Open(l)
}
