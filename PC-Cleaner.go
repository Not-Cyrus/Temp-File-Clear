package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var (
	paths = map[int]string{
		0: os.Getenv("TEMP"),
		1: os.Getenv("APPDATA") + `\Discord\Cache`,
		2: os.Getenv("LOCALAPPDATA") + `\CrashDumps`,
		3: os.Getenv("LOCALAPPDATA") + `\Steam\htmlcache\Cache`,
		4: os.Getenv("APPDATA") + `\discordptsb\Cache`,
		5: os.Getenv("APPDATA") + `\discordcanary\Cache`,
	}
	totalFiles = 0
)

func main() {
	for _, path := range paths {
		files, _ := ioutil.ReadDir(path)
		for _, file := range files {
			err := os.Remove(path + `\` + file.Name()) // you see, I could use os.RemoveAll but I rather not make hacky code to remake the directory
			if err == nil {
				totalFiles++
			}
		}
	}
	fmt.Println(fmt.Sprintf("Deleted %d files", totalFiles))
	for i := 5; i > 0; i-- {
		fmt.Println(fmt.Sprintf("Closing in %d seconds, You can close it yourself if you want", i))
		time.Sleep(time.Duration(1) * time.Second)
	}
	os.Exit(0)
}
