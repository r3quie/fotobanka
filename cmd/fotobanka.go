/*
Copyright © 2024 Pavel Kadera info@pavelkadera.cz

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
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/r3quie/fotobanka/install"
	"github.com/spf13/cobra"
)

// bankaCmd represents the fotobanka command
var bankaCmd = &cobra.Command{
	Use:   "banka",
	Short: "Zobrazí obrázky z foto banky.",
	Long: `Zobrazí obrázky z foto banky. Obrázky jsou vykresleny pomocí ASCII art.
	Pro zobrazení celého obrázku, držte klávesu CTRL a zmáčkněte klávesu +/- nebo použijte kolečko myši.
	Obrázky jsou uloženy v adresáři imgs.`,
	Run: func(cmd *cobra.Command, args []string) {
		var localData string
		if runtime.GOOS == "windows" {
			localData = filepath.Join(os.Getenv("APPDATA"), "fotobanka")
		} else {
			localData = filepath.Join(os.Getenv("HOME"), ".fotobanka")
		}

		if !install.AreImgsInPwd("") && !install.AreImgsInPwd(localData) {
			install.DownloadImgs()
		}
		dir, err := os.ReadDir("imgs")
		if err != nil {
			dir, err = os.ReadDir(filepath.Join(localData, "imgs"))
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		fmt.Println("Zadejte číslo obrázku, který chcete zobrazit:")
		for i, f := range dir {
			fmt.Println(i+1, f.Name())
		}
		var i int
		fmt.Scanf("%d", &i)
		if i < 1 || i > len(dir) {
			fmt.Println("Neplatný vstup.")
			return
		}
		PrintImg(dir[i-1].Name())
	},
}

func init() {
	rootCmd.AddCommand(bankaCmd)
}
