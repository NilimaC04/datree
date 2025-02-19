package cliClient

import (
	"github.com/datreeio/datree/bl/files"
	"net/http"
)

func (c *CliClient) PublishPolicies(policiesConfiguration files.UnknownStruct, cliId string) error {
	headers := map[string]string{"x-cli-id": cliId}
	_, err := c.httpClient.Request(http.MethodPut, "/cli/policy/publish", policiesConfiguration, headers)
	return err
}
