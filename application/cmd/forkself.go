/*
   Nging is a toolbox for webmasters
   Copyright (C) 2018-present Wenhui Shen <swh@admpub.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/admpub/log"
	"github.com/admpub/nging/v5/application/library/selfupdate"
	"github.com/kardianos/osext"
	"github.com/spf13/cobra"
)

var forkSelfCmd = &cobra.Command{
	Use:     "forkself",
	Short:   "Fork self",
	Example: "forkself",
	RunE:    forkSelfRunE,
}

func forkSelfRunE(cmd *cobra.Command, args []string) error {
	executable, err := osext.Executable()
	if err != nil {
		log.Errorf(`[forkself]osext.Executable(): %s`, err.Error())
		return err
	}
	procArgs := []string{executable}
	if len(os.Args) > 2 { // <workDir>/nging forkself [args...]
		procArgs = append(procArgs, os.Args[2:]...)
	}
	workDir := filepath.Dir(executable)
	log.Infof(`[forkself]command: %s`, strings.Join(procArgs, ` `))
	log.Infof(`[forkself]workDir: %s`, workDir)
	_, err = os.StartProcess(executable, procArgs, &os.ProcAttr{
		Dir:   workDir,
		Env:   os.Environ(),
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Sys:   selfupdate.NewSysProcAttr(),
	})
	if err != nil {
		log.Errorf(`[forkself]error: %s`, err.Error())
	}
	return err
}

func init() {
	rootCmd.AddCommand(forkSelfCmd)
}
