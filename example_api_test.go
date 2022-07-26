package eos_test

import "os"

func getAPIURL() string {
	apiURL := os.Getenv("EOS_CHAIN_API_URL")
	if apiURL != "" {
		return apiURL
	}

	return "https://api.eosn.io/"
}
