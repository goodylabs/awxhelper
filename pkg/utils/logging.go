package utils

import (
	"log"

	"github.com/goodylabs/awxhelper/pkg/config"
)

func OptionalLog(msg string) {
	if config.IsVerboseMode() {
		log.Println(msg)
	}
}
