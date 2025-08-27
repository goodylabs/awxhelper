package awxconnector

import "github.com/goodylabs/awxhelper/internal/services/dto"

func main() {
	// testing purposes
	instance := NewAwxConnector()
	instance.ConfigureConnection(&dto.AwxConfig{})
}
