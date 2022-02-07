package commands

import (
	"bytes"
	"fmt"
	"github.com/Tnze/go-mc/chat"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type CommandClient struct {
	ApiKey string
}

type Command struct {
	CommandName string
	Run func() (result string, err error)
}

func NewCommandClient(apiKey string) *CommandClient {
	return &CommandClient{apiKey}
}

func (c *CommandClient) getResultFromCommand(serverName, commandName string) (result string, err error) {
	var data bytes.Buffer
	data.Write([]byte(commandName))
	req, err := http.NewRequest("POST", "https://infra.laybraid.fr/server/send/" + serverName, &data)
	if err != nil {
		log.Errorln(err)
		return "", err
	}
	req.Header.Set("apiKey", c.ApiKey)
	post, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorln(err)
		return "", err
	}

	all, err := io.ReadAll(post.Body)
	if err != nil {
		log.Errorln(err)
		return "", err
	}

	if post.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("Error %d: %s", post.StatusCode, string(all)))
	}

	seq, _ := chat.TransCtrlSeq(string(all), false)

	return seq, nil
}
