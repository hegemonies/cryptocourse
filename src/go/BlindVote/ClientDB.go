package BlindVote

import (
	"crypto/sha1"
	"encoding/hex"
)

type ClientDB struct {
	clients map[string]*VoteClient
}

func InitClientDB() *ClientDB {
	cdb := &ClientDB{}
	cdb.clients = make(map[string]*VoteClient)
	return cdb
}

func (c *ClientDB) RegisterClient(name string) (*VoteClient, bool) {
	id := ComputeHashToString(name)
	if c.CheckExistsUser(id) {
		return nil, false
	}
	client := &VoteClient{}
	client.Name = id
	c.clients[id] = client
	return client, true
}

func (c *ClientDB) CheckExistsUser(name string) bool {
	if _, ok := c.clients[name]; ok {
		return true
	}
	return false
}

func (c *ClientDB) AddClient(client *VoteClient) bool {
	if c.CheckExistsUser(client.Name) {
		c.clients[client.Name] = client
	}
	return false
}

func ComputeHashToString(str string) string {
	hasher := sha1.New()
	hasher.Write([]byte(str))
	checksum := hasher.Sum(nil)
	return hex.EncodeToString(checksum)
}
