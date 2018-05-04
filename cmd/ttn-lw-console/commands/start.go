// Copyright © 2018 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"github.com/spf13/cobra"
	"go.thethings.network/lorawan-stack/pkg/assets"
	"go.thethings.network/lorawan-stack/pkg/component"
	"go.thethings.network/lorawan-stack/pkg/console"
	"go.thethings.network/lorawan-stack/pkg/errors"
)

var (
	startCommand = &cobra.Command{
		Use:   "start",
		Short: "Start the Console",
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := component.New(logger, &component.Config{
				ServiceBase: config.ServiceBase,
			})
			if err != nil {
				return errors.NewWithCause(err, "Failed to initialize base component")
			}

			_, err = console.New(c, assets.New(c, config.Assets), config.Console)
			if err != nil {
				return errors.NewWithCause(err, "Failed to initialize Console")
			}

			logger.Info("Starting Console...")
			return c.Run()
		},
	}
)

func init() {
	Root.AddCommand(startCommand)
}