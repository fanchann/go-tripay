package gotripay

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type TRIPAY_SIGNATURE string

type Signature struct {
	Amount       int64
	PrivateKey   string
	MerchantCode string
	MerchantRef  string
	Channel      TRIPAY_CHANNEL
}

func (s *Signature) CreateSignature() TRIPAY_SIGNATURE {
	var signStr string
	if s.Amount != 0 {
		signStr = s.MerchantCode + s.MerchantRef + fmt.Sprint(s.Amount)
	} else {
		signStr = s.MerchantCode + string(s.Channel) + s.MerchantRef
	}

	key := []byte(s.PrivateKey)
	message := []byte(signStr)

	hash := hmac.New(sha256.New, key)
	hash.Write(message)
	signature := hex.EncodeToString(hash.Sum(nil))

	return TRIPAY_SIGNATURE(signature)
}
