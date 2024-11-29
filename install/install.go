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
package install

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"unsafe"
)

func OpenCmd() {
	cmd := exec.Command("cmd", "/c", "start", "cmd", "/k", "echo Pro spuštění programu zadejte 'fotobanka banka' a pro zobrazení licence fotobanky zadejte 'fotobanka licence'")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		fmt.Println("Chyba při spouštění příkazového řádku:", err)
		return
	}
}

// Determines whether the current process has a parent process. Used to determine if the process is running in a console window.
func IsInTerminal() bool {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	procGetConsoleProcessList := kernel32.NewProc("GetConsoleProcessList")

	var processList [2]uint32
	count, _, _ := procGetConsoleProcessList.Call(uintptr(unsafe.Pointer(&processList[0])), uintptr(len(processList)))

	return !(count == 1)
}

func AreImgsInPwd(pwd string) bool {
	_, err := os.Stat(filepath.Join(pwd + "imgs"))
	return err == nil
}

func DownloadImgs() {
	url := "https://www.r3quie.com/imgs.zip"
	destDir := filepath.Join(os.Getenv("APPDATA"), "fotobanka")
	zipFile := filepath.Join(os.TempDir(), "imgs.zip")

	// Download the zip file
	fmt.Println("Stahuji imgs.zip...")
	err := downloadFile(zipFile, url)
	if err != nil {
		fmt.Println("Chyba při stahování:", err)
		return
	}

	// Create destination directory if it doesn't exist
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		err = os.MkdirAll(destDir, os.ModePerm)
		if err != nil {
			fmt.Println("Chyba při vytváření adresáře:", err)
			return
		}
	}

	// Extract the zip file
	err = unzip(zipFile, destDir)
	if err != nil {
		fmt.Println("Chyba při rozbalování:", err)
		return
	}

	// Clean up
	err = os.Remove(zipFile)
	if err != nil {
		fmt.Println("Chyba při odstraňování dočasného souboru:", err)
	}
}

func downloadFile(filepath string, url string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func unzip(src string, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}
	return nil
}
