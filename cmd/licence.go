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
	"time"

	"github.com/spf13/cobra"
)

// licenceCmd represents the licence command
var licenceCmd = &cobra.Command{
	Use:   "licence",
	Short: "Zobrazí licenci foto banky.",
	Long: `Zobrazí licenci foto banky. Licence je zobrazena postupně 
	po řádcích s časovým odstupem 100 ms.`,
	Run: func(cmd *cobra.Command, args []string) {
		PrintLicence(ABST_LIC)
		var i string
		for {
			i = ""
			fmt.Println("Zadejte " + Col("písmeno", BBlue) + " a " + Col("číslo", BPurple) + " ujednání licence, jehož podrobnosti chcete zobrazit: (např. " + Col("A", BBlue) + Col("2", BPurple) + ", " + Col("C", BBlue) + ")\nZadejte 'q' pro ukončení programu.")
			fmt.Scanln(&i)

			if i == "q" || i == "Q" {
				fmt.Println("Pro zobrazní fotobanky zadejte 'fotobanka banka'. Pro zobrazení licence zadejte 'fotobanka licence'.")
				break
			} else if LICMAP[i] == "" {
				fmt.Println("Nesprávný vstup")
				continue
			}
			for i := 0; i < 33; i++ {
				time.Sleep(10 * time.Nanosecond)
				fmt.Print("###")
			}
			PrintLicence(LICMAP[i])
			time.Sleep(500 * time.Millisecond)
		}
	},
}

func init() {
	rootCmd.AddCommand(licenceCmd)
}
