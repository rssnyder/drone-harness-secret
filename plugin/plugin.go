// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"log"

	"github.com/rssnyder/harness-go-utils/config"
	"github.com/rssnyder/harness-go-utils/secrets"
)

// Args provides plugin execution arguments.
type Args struct {
	Pipeline

	// Level defines the plugin log level.
	Level string `envconfig:"PLUGIN_LOG_LEVEL"`

	Name          string `envconfig:"PLUGIN_NAME"`
	Value         string `envconfig:"PLUGIN_VALUE"`
	SecretManager string `envconfig:"PLUGIN_SECRET_MANAGER"`
}

// Exec executes the plugin.
func Exec(ctx context.Context, args Args) (err error) {

	client, hCtx := config.GetNextgenClient()

	err = secrets.SetSecretText(hCtx, client, args.Name, args.Value, args.SecretManager)
	if err != nil {
		return err
	}

	log.Printf("saved: %s\n", args.Name)

	return
}
