package zkosh

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

func SignatureZoksh(SecretKey string) (string, error) {
	// POST request body

	postBody := map[string]interface{}{
		"transaction": "0x01c483d2deb658e7cd6beea753aad0e176ea508b517b01eab9b45bf8e03b3a15",
	}

	// stringify the POST request body
	postBodyJSON, _ := json.Marshal(postBody)

	// Unix timestamp of when the request is being made. Same as ZOKSH-TS.
	ts := time.Now().UnixMilli()

	// Zoksh API Server URL you are sending your POST request to.
	requestPath := "/v1/validate-payment"

	// creating hmac object using your API Secret Key.
	hmac := hmac.New(sha256.New, []byte(SecretKey))

	// combined string to be signed
	toSign := fmt.Sprintf("%d%s%s", ts, requestPath, postBodyJSON)

	// ZOKSH-SIGN
	hmac.Write([]byte(toSign))
	signature := hex.EncodeToString(hmac.Sum(nil))

	fmt.Printf("ZOKSH-TS: %d\nZOKSH-SIGN: %s\n", ts, signature)

	return signature, nil
}
