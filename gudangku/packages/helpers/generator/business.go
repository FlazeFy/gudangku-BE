package generator

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
	"time"
)

func GetSlug(val string) string {
	res := strings.ReplaceAll(val, " ", "-")

	res = strings.ReplaceAll(res, "_", "-")

	regExp := regexp.MustCompile(`[!:\\\[/"\;\.\'^£$%&*()}{@#~?><>,|=+¬\]]`)
	res = regExp.ReplaceAllString(res, "")

	return res
}

func GenerateUUID(len int) (string, error) {
	var id string
	uuid := make([]byte, 16)

	_, err := rand.Read(uuid)
	if err != nil {
		return "", err
	}

	uuid[6] = (uuid[6] & 0x0F) | 0x40
	uuid[8] = (uuid[8] & 0x3F) | 0x80

	if len == 16 {
		id = fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
	} else if len == 32 {
		id = hex.EncodeToString(uuid)
	} else {
		id = ""
	}

	return id, nil
}

func GenerateTimeNow(name string) string {
	if name == "timestamp" {
		now := time.Now()
		res := now.Format("2006-01-02 15:04:05")
		return res
	}
	return ""
}
