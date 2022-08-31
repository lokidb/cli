package cli

import (
	"fmt"

	client "github.com/lokidb/server/client"
)

func errorToMessage(err error) string {
	return fmt.Sprintf("ERROR: %s\n", err.Error())
}

func get(c *client.Client, key string) string {
	saveKey(key)
	value, err := c.Get(key)
	saveValue(value)

	if err != nil {
		return errorToMessage(err)
	}

	return value
}

func set(c *client.Client, key string, value string) string {
	saveKey(key)
	saveValue(value)

	err := c.Set(key, value)
	if err != nil {
		return errorToMessage(err)
	}

	return "done!\n"
}

func del(c *client.Client, key string) string {
	saveKey(key)
	deleted, err := c.Del(key)

	if err != nil {
		return errorToMessage(err)
	}

	if deleted {
		return "deleted\n"
	} else {
		return "not found\n"
	}
}

func keys(c *client.Client) string {
	ks, err := c.Keys()
	if err != nil {
		return errorToMessage(err)
	}

	message := ""
	for i, key := range ks {
		saveKey(key)
		message += fmt.Sprintf("%d) %s\n", i, key)
	}

	return message
}

func flush(c *client.Client) string {
	err := c.Flush()
	if err != nil {
		return errorToMessage(err)
	}

	return "flushed all keys\n"
}
