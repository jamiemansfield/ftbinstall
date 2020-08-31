// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/jamiemansfield/go-modpacksch/modpacksch"
	"github.com/jamiemansfield/mcinstall/ftb"
	"github.com/jamiemansfield/mcinstall/minecraft"
	"github.com/jamiemansfield/mcinstall/util"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "ftbinstall",
		Usage:   "install packs from the modpacks.ch service",
		Version: "0.1.0-indev",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "target",
				Aliases: []string{"t"},
				Usage:   "sets the install target",
				Value:   "client",
			},
		},
		Action: func(ctx *cli.Context) error {
			if ctx.Args().Len() < 2 {
				return errors.New("usage: ftbinstall pack version")
			}

			packId, err := strconv.Atoi(ctx.Args().Get(0))
			if err != nil {
				return errors.New("usage: pack must be an integer")
			}
			versionId, err := strconv.Atoi(ctx.Args().Get(1))
			if err != nil {
				return errors.New("usage: version must be an integer")
			}
			installTargetRaw := ctx.Value("target").(string)

			var installTarget minecraft.InstallTarget
			if installTargetRaw == "client" || installTargetRaw == "c" {
				installTarget = minecraft.Client
			} else
			if installTargetRaw == "server" || installTargetRaw == "s" {
				installTarget = minecraft.Server
			} else {
				return errors.New("unknown install target " + installTargetRaw)
			}

			client := modpacksch.NewClient(nil)
			client.UserAgent = util.UserAgent

			pack, err := client.Packs.GetPack(packId)
			if err != nil {
				return err
			}

			version, err := client.Packs.GetVersion(packId, versionId)
			if err != nil {
				return err
			}

			return ftb.InstallPackVersion(installTarget, "", pack, version)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
