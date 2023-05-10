package main

import (
	"math/rand"
	"time"
	"vkIntership/config"
	"vkIntership/core"
)

func main() {
	rand.Seed(time.Now().Unix())

	c, err := core.New(config.AccessToken, config.APIVersion)
	if err != nil {
		panic(err)
	}

	if err = c.Run(); err != nil {
		panic(err)
	}
}
