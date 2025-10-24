package utils

import (
	"log"

	"github.com/goodylabs/awxhelper/pkg/config"
)

func OptionalLog(msg string) {
	if config.IsDebugMode() {
		log.Println(msg)
	}
}
