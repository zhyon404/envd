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

package main

import (
	"fmt"
	"os"

	"github.com/cockroachdb/errors"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
	"go.starlark.net/starlark"
	"go.starlark.net/syntax"

	"github.com/tensorchord/envd/pkg/app"
	"github.com/tensorchord/envd/pkg/flag"
	"github.com/tensorchord/envd/pkg/version"
)

func run(args []string) error {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println(c.App.Name, version.Package, c.App.Version, version.Revision)
	}

	app := app.New()
	return app.Run(args)
}

func handleErr(err error) {
	if err == nil {
		return
	}

	if viper.GetBool(flag.FlagDebug) {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}

	rootCause := errors.Cause(err)
	if evalErr, ok := rootCause.(*starlark.EvalError); ok {
		fmt.Fprintln(os.Stderr, evalErr.Backtrace())
	} else if syntaxErr, ok := rootCause.(*syntax.Error); ok {
		fmt.Fprintln(os.Stderr, syntaxErr)
	} else {
		fmt.Fprintf(os.Stderr, "error: %v\n", rootCause)
	}

	os.Exit(1)
}

func main() {
	err := run(os.Args)
	handleErr(err)
}
