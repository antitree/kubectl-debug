/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/pflag"

	"github.com/verb/kubectl-debug/pkg/cmd"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func main() {
	flags := pflag.NewFlagSet("kubectl-debug", pflag.ExitOnError)
	pflag.CommandLine = flags

	_, calledAs := filepath.Split(os.Args[0])
	if strings.HasPrefix(calledAs, "kubectl-") {
		calledAs = strings.Replace(calledAs, "kubectl-", "kubectl ", 1)
		calledAs = strings.ReplaceAll(calledAs, "_", "-")
	}

	root := cmd.NewCmdDebug(calledAs, genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
