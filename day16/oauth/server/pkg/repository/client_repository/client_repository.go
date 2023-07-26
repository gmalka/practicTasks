package clientrepository

import (
	"encoding/base64"
	"errors"
	"math/rand"
	"time"
)

type Clients struct {
	id      map[int]string
	names   map[string]int
	secrets map[string]string
}

func NewClient() Clients {
	return Clients{id: make(map[int]string), names: make(map[string]int), secrets: make(map[string]string)}
}

func (c Clients) RandomString(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:length]
}

func (c Clients) Create(name string) error {
	rand.Seed(time.Now().UnixNano())

	if _, ok := c.names[name]; ok {
		return errors.New("app clready exists")
	}

	for {
		randomNumber := rand.Intn(900001) + 100000
		if _, ok := c.id[randomNumber]; !ok {
			c.id[randomNumber] = name
			c.names[name] = randomNumber
			c.secrets[name] = c.RandomString(8)
			return nil
		}
	}
}

func (c Clients) GetByName(name string) (int, string, bool) {
	if v, ok := c.names[name]; ok {
		return v, c.secrets[name], true
	}

	return 0, "", false
}

func (c Clients) GetById(id int) (string, string, bool) {
	if v, ok := c.id[id]; ok {
		return v, c.secrets[v], true
	}

	return "", "", false
}
