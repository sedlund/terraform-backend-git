package sops

import (
	"os"

	sops "go.mozilla.org/sops/v3"
	"go.mozilla.org/sops/v3/age"
)

func init() {
	Configs["age"] = &AgeConfig{}
}

type AgeConfig struct{}

func (c *AgeConfig) IsActivated() bool {
  _, ok := os.LookupEnv("TF_BACKEND_HTTP_SOPS_AGE_RECIPIENTS")
	return ok
}

func (c *AgeConfig) KeyGroup() (sops.KeyGroup, error) {
	recipients := os.Getenv("TF_BACKEND_HTTP_SOPS_AGE_RECIPIENTS")

  ageKeys, err := age.MasterKeysFromRecipients(recipients)
  if err != nil {
		return nil, err
	}

	var keyGroup sops.KeyGroup

  for _, k := range ageKeys {
		keyGroup = append(keyGroup, k)
	}

	return keyGroup, nil
}
