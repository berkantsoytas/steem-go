package jsonrpc2_test

import (
	"testing"

	"github.com/berkantsoytas/steem-go/jsonrpc2"
)

func TestBuildSendData(t *testing.T) {
	client := jsonrpc2.NewClient("https://api.steemit.com")
	client.NewRequest("condenser_api.get_accounts", []interface{}{"berkantsoytas", "steemit"})

	getData := client.Data

	expected := "2.0"

	if getData.Version != expected {
		t.Errorf("Expected %s, got %s", expected, getData.Version)
	}
}

func TestSend(t *testing.T) {
	client := jsonrpc2.NewClient("https://api.steemit.com")

	accounts := []string{"steemit"}

	client.NewRequest("condenser_api.get_accounts", []interface{}{accounts})

	response, err := client.Send()

	if err != nil {
		t.Error(err)
	}

	if len(response.Result.([]interface{})) == 0 {
		t.Error("Expected more than 0, got 0")
	}
}
