package main

import (
	cfg "github.com/conductorone/baton-percipio/pkg/config"
	"github.com/conductorone/baton-sdk/pkg/config"
)

func main() {
	config.Generate("percipio", cfg.Config)
}
