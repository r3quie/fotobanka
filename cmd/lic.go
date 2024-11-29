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

const ABST_LIC = `
            PRAVIDLA UŽÍVÁNÍ OBRÁZKŮ Z FOTOBANKY
                  
   ` + BBlue + "A" + Clear + `. Obrázek z fotobanky ` + BGreen + `smíte` + Clear + `:
         ` + BPurple + "1" + Clear + `. použít pro vlastní účely,
         ` + BPurple + "2" + Clear + `. upravit,
         ` + BPurple + "3" + Clear + `. spojit s jiným dílem,
         ` + BPurple + "4" + Clear + `. použít v jakékoli zemi.

   ` + BBlue + "B" + Clear + `. Obrázek z fotobanky ` + BRed + `nesmíte` + Clear + ` použít komerčně.

   ` + BBlue + "C" + Clear + `. U obrázku z fotobanky ` + BYellow + `nemusíte` + Clear + ` uvádět autora.
   `

const A1 = `
   Obrázek z fotobanky ` + BGreen + `smíte` + Clear + ` použít pro vlastní účely.
   
   Vlastními účely se rozumí použití obrázku pro soukromé, osobní nebo nevýdělečné účely.
   
   ` + BRed + `Výdělečné účely` + Clear + ` jsou například:
      - prodej výrobku s obrázkem,
      - vložení obrázku do reklamy,
      - vložení obrázku do knihy, která se prodává.

   ` + BGreen + `Osobní účely` + Clear + ` jsou například:
      - užití obrázku jako pozadí na monitoru,
      - použití na osobním webu nebo sociálních sítích,
      - tisk do fotoalba.
      `

const A2 = `
   Obrázek z fotobanky ` + BGreen + `smíte` + Clear + ` upravit.
   
   Úpravou se rozumí například:
      - oříznutí,
      - změna barev,
      - změna velikosti,
      - přidání textu nebo jiné grafiky
   `
const A3 = `
   Obrázek z fotobanky ` + BGreen + `smíte` + Clear + ` spojit s jiným dílem.

   Spojením s jiným dílem se rozumí například:
      - vložení obrázku do prezentace,
      - vložení obrázku do knihy,
      - vložení obrázku do webové stránky.
`

const A4 = `
   Obrázek z fotobanky ` + BGreen + `smíte` + Clear + ` použít v jakékoli zemi.

   Obrázek můžete použít v jakékoli zemi světa bez omezení. Licence není teritoriálně omezena.
`

const B = `
   Obrázek z fotobanky ` + BRed + `nesmíte` + Clear + ` použít komerčně.

   ` + BRed + `Komerčním použitím ` + Clear + ` se rozumí například:
      - prodej výrobku s obrázkem,
      - vložení obrázku do reklamy,
      - vložení obrázku do knihy, která je určena k prodeji.
`

const C = `
   U obrázku z fotobanky ` + BYellow + `nemusíte` + Clear + ` uvádět autora.

   Uvedení autora ` + BYellow + `není povinné` + Clear + `, ale je vítáno. Pokud si autora přejete podpořit,
   můžete uvést například v popisu obrázku nebo na webu, kde je obrázek zveřejněn.
   Stejně tak ` + BYellow + `není povinné` + Clear + ` uvádět jméno této fotobanky.
   ` + BRed + `Nesmíte` + Clear + ` však uvádět, že jste autorem obrázku, pokud tomu tak opravdu není.
`

var LICMAP = map[string]string{"A1": A1, "A2": A2, "A3": A3, "A4": A4, "B": B, "C": C}

const LIC = `
                  LICENČNÍ SMLOUVA K UŽITÍ DĚL VE FOTOBANCE

   1. Definice

      "Fotografií" se rozumí jednotka databáze. Fotografie je dílem ve 
      smyslu zákona č. 121/2000 Sb., autorský zákon, ve znění pozdějších 
      předpisů.
   
      "Licencí" se rozumí tato licenční smlouva ve smyslu 
      zákona č. 89/2012 Sb. (dále jen "OZ"). Licence je poskytována
      k jednotlivým Fotografiím.

      "Poskytovatelem" se rozumí osoba, která poskytuje oprávnění k 
      výkonu práva duševního vlastnictví k Fotografii a je oprávněna 
      k výkonu majetkových práv autorských k dílu nebo oprávněna k výkonu 
      práva duševního vlastnictví k Fotografii na základě nevýhradní licence.

      "Fotobankou" se rozumí databáze, která obsahuje Fotografie a je
      spravována Poskytovatelem.

      "Nabyvatelem" se rozumí osoba nabývající oprávnění k výkonu práva 
      duševního vlastnictví z Licence k Fotografii.
`
