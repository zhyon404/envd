// Copyright 2022 The envd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	"github.com/cockroachdb/errors"
	"github.com/sirupsen/logrus"
	"github.com/tensorchord/envd/pkg/home"
	"github.com/urfave/cli/v2"
)

var CommandContextRemove = &cli.Command{
	Name:  "rm",
	Usage: "Remove envd context",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "name",
			Usage: "Name of the context",
			Value: "",
		},
	},
	Action: contextRemove,
}

func contextRemove(clicontext *cli.Context) error {
	name := clicontext.String("name")

	err := home.GetManager().ContextRemove(name)
	if err != nil {
		return errors.Wrap(err, "failed to remove context")
	}
	logrus.Infof("Context %s is removed", name)
	return nil
}
