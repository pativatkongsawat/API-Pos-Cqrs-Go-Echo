package utils

import "github.com/google/uuid"

func GennerateUUID() string {
	return uuid.New().String()
}
