package config

import "os"

var (
	APIPort     = os.Getenv("APIPort")
	APIKey      = os.Getenv("APIKey")
	TokenSecret = "SuperbMIFTAH"
)
