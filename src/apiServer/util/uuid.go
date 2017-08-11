package util

import uuid "github.com/satori/go.uuid"

/**
获取生成的 UUID
*/
func GetUUID() string {
	u := uuid.NewV4()
	return u.String()
}
