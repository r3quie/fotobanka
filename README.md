# Licence Fotobanky (CLI)

# Koncepce programu

Prvním krokem při tvorbě práce bylo sepsat koncepční plán pro aplikaci. Vzhledem k tomu, že se mělo jednat jednoduchý program, který bude imitovat funkce fotobanky a poskytne uživateli “licenční smlouvu”, bylo nutné znát, co přesně by měl program umět a jak by měl vypadat.

Nakonec byl zvolen tzv. CLI (command line interface) program, neboť toto je jediná forma dedikovaného programu, u které je ospravedlnitelné nepoužívat namísto něj jednoduchou webovou stránku. Jinými slovy je přirozenější pro provozovatele fotobanky vyvinout/poptat webovou stránku než vytvářet program běžící přímo na koncovém zařízení. Pro potřeby seminární práce se však webová stránka jevila jako příliš neoriginální.

CLI program je (přestože pouze vratce) ospravedlnitelná alternativa webové stránce, neboť mnoho vývojářů a počítačových techniků pracuje s počítačem pomocí vzdáleného přístupu (prostřednictvím protokolu SSH), a tak mají k dispozici pouze příkazový řádek. Při vývoji by jim tak mohla přijít vhod fotobanka, která poskytuje obrázky právě prostřednictvím příkazového řádku (tj. obrázek stáhne a poskytne jeho jednoduchý náhled v tzv. ASCII art prostřednictvím příkazového řádku).

Samozřejmě samotné technické řešení programu není ničím inovativní (přestože kód pro vykreslování je zcela vytvořen autorem seminární práce) a samotný koncept CLI fotobanky je pravděpodobně v praxi s těží využitelný. Účelem programu je spíše ilustrovat, že zobrazení licence nějaké CLI služby je jistě věc, se kterou je možné se setkat. Samotná funkce vykreslení obrázku pomocí ASCII znaků je jenom přidaná hodnota a “proof of concept”.

Během vytváření seminární práce bylo také nutné počítat s tím, že osoba ji hodnotící nemusí mít zkušenost s prací s takovým typem aplikace. Program tak byl opatřen aparátem, který při detekování, že program nebyl spuštěn v příkazovém řádku, automaticky příkazový řádek spustí a zobrazí uživateli základní instrukce k užití. Program takto i automaticky nainstaluje vše potřebné pro jeho chod.

# Zdrojový kód

Druhým krokem bylo sepsat zdrojový kód pro výše popsané funkce. Prvním krokem bylo vytvořit funkci, která dokáže předat text licence uživateli nějakým jednoduše stravitelným způsobem. Většina emulátorů příkazových řádků (tzv. terminal emulators) jsou schopny zobrazit všechny barvy na standardním RGB spektru. Jednotlivé výstupní funkce tak byly opatřeny aparátem, který zvýrazňuje povolení a zákazy, jež jsou vyjádřeny v ujednáních licenční smlouvy, korespondujícími barvami.

Pro dokázání, že licenční smlouva fotobanky opravdu může být vyjádřena v CLI programu, byl dále vytvořen jednoduchý proces, který dokáže konvertovat obrázky v jpg formátu na obrázky v podobě textu, resp. konvertovat jednotlivé pixely rastrového obrázku na korespondující písmeno nebo jiný ACII charakter.

Při tomto procesu byl využit kód z repositáře [nfnt/resize](https://github.com/nfnt/resize) pod ISC licencí (ekvivalent MIT licence), který umožnil jednoduché zmenšování a zvětšování vstupních obrázků. Samotný program je napsán v jazyce Go za použití jeho standardní knihovny, a to pod licencí BSD‑3‑Clause.

Dalším krokem byla implementace výše uvedených funkcionalit takovým způsobem, aby je mohl užívat koncový uživatel. Tomuto procesu předcházela stručná rešerše vývoje CLI programů v jazyce Go. Po rešerši byla zvolena knihovna [spf13/cobra](https://github.com/spf13/cobra) pod licencí Apache‑2.0. Tato knihovna poskytuje základní “kostru” (tzv. framework) pro umožnění implementace výše popsaných funkcí ve standardu, který následují nejpoužívanější CLI programy.

Program nemá žádný „backend“ – tedy žádnou serverovou část, ke které by se koncový uživatel prostřednictvím programu připojoval a která by uživateli poskytovala obrázky. Pro potřeby seminární práce by taková funkcionalita byla zbytečná. V tuto chvíli je program „distribuován“ společně s ukázkovými obrázky, které byly získány pod alespoň creative commons licencí (tj. „nekomerční“ bezplatná licence).

# Implementace programu koncovému uživateli

Jako největší překážka se projevila nativní antivirová ochranu operačního systému Windows, která značně znesnadňovala pokusy o implementaci programu. Binární soubory totiž v dnešním kyberbezpečnostním standardu vyžadují přinejmenším uznávaný nebo až kvalifikovaný digitální podpis, a to certifikátem, který k tomu byl přímo určen (tzv. code‑signing certificate) a vydán příslušnou certifikační autoritou; nestačí osobní certifikát, který poskytuje Universita Karlova svým studentům. Takový certifikát by byl pro potřeby seminární práce přílišnou finanční zátěží, a tak byl zvolen takový přístup, který minimalizuje šanci, že antivirus program koncovému uživateli automaticky odstraní.

Bohužel kompletní eliminace upozornění operačního systému na “potenciálně škodlivý software” při spouštění programu nebyla možná. V rámci prevence antivirové, technické nebo lidské chyby byl nakonec zvolen takový přístup, který zcela zabrání možným problémům s implementací, a to pomocí vzdáleného přístupu do virtuálního serveru (pomocí protokolu SSH), na kterém již byla aplikace úspěšně implementována.

# Sepsání licence

Posledním krokem bylo sepsání “licence” k obrázkům ve fotobance. Proto, že se dle zadání seminární práce nemuselo jednat o vyčerpávající právní text, byl zvolen přístup, který bude pro fiktivního koncového uživatele co nejpochopitelnější – tedy prostá čeština a příklady vysvětlující jednotlivé pojmy, pokud uživatel zadá číslo odstavce, který potřebuje vysvětlit.

Při sepisování „licence“ byla užita umělá inteligence Copilot, která využívá jazykový model GPT‑4o. Specificky byla použita dle tabulky níže. Dále byla umělá inteligence použita při sepisování instalačního scriptu.

Úvodní prompt: _Odpovídej stručně a pouze tím, co po tobě budu explicitně vyžadovat._

| Prompt | Odpověď |
| --- | --- |
| _Uveď několik příkladů komerčního užití obrázku._ | Reklamní kampaně (billboardy, plakáty, online bannery).<br><br>Obaly produktů (potraviny, elektronika, kosmetika).<br><br>Webové stránky a e-shopy.<br><br>Firemní prezentace a marketingové materiály.<br><br>Sociální sítě (příspěvky, placené reklamy).<br><br>Merchandising (trička, hrnky, plakáty).<br><br>Knihy, časopisy a brožury (ilustrace, obálky). |
| _Uveď několik příkladů úpravy obrázku._ | Ořezání a změna velikosti.<br><br>Úprava jasu, kontrastu a saturace.<br><br>Retušování nedokonalostí.<br><br>Přidání textu nebo loga.<br><br>Aplikace filtrů nebo efektů.<br><br>Odstranění nebo nahrazení pozadí.<br><br>Skládání více obrázků (koláže).<br><br>Převod do jiného formátu (např. JPEG na PNG). |
| _Uveď několik příkladů osobního užití obrázku._ | Tapeta na mobil nebo počítač.<br><br>Sdílení na osobních sociálních sítích.<br><br>Tisk do fotoalba nebo na stěnu.<br><br>Výroba osobních přání nebo pozvánek.<br><br>Použití v osobním blogu nebo portfoliu.<br><br>Inspirace pro vlastní kresbu nebo tvorbu. |
| _Rozveď do pár vět pravidlo „U obrázku nemusíte uvádět autora.“_ | Pravidlo „U obrázku nemusíte uvádět autora“ znamená, že není povinností připisovat autorství při použití obrázku, pokud to není vyžadováno licenčními podmínkami. To se často týká obrázků ve veřejné doméně nebo těch s licencemi umožňujícími volné užití bez uvedení zdroje (např. některé verze licencí CC0). U obrázků chráněných autorským právem však obvykle uvedení autora nebo zdroje vyžaduje zákon nebo licenční dohoda. Nejasnosti je třeba vždy ověřit v podmínkách užití konkrétního díla. |
| _Co by obvykle řekla fotobanka, aby uživatel uvedl jméno autora obrázku (přestože to nevyžaduje licence). Jednou nebo dvěma větami._ | Fotobanka by obvykle doporučila uvést jméno autora jako gesto uznání a podpory jejich práce, i když to licence nevyžaduje. Mohla by také zmínit, že takový přístup může přispět k větší viditelnosti autora. |
| _Write a batch command that opens a new terminal window with a specified message._ | REM Open a new terminal window with the message start cmd /k "echo Pro spusteni programu zadejte 'fotobanka banka' pro zobrazeni licence zadejte 'fotobanka licence'" |
| _Do the same with powershell_ | # Launch Command Prompt with the specified message Start-Process cmd -ArgumentList "/k echo Pro spusteni programu zadejte 'fotobanka banka' pro zobrazeni licence zadejte 'fotobanka licence'" |
