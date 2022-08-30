package Phoenix_Graphics

import (
	"bufio"
	"fmt"
	"log"
	Options "main/modules/Options"
	"os"
)

const (
	Banner_file1 = "text/banners/banner_long.txt"
	Banner_file2 = "text/banners/banner_long2.txt"
	Banner_file3 = "text/banners/banner_verti.txt"
	Banner_file4 = "text/banners/banner_verti2.txt"
)

func Open_File(banner_file, color string, cl bool) {
	switch cl {
	case true:
		fmt.Println("\x1b[H\x1b[2J\x1b[3J")
	}
	var Opt Options.Options
	if Opt.Screen_Resolution == "" {
		banner_file = Banner_file1
	}
	if Opt.Screen_Resolution == "Verticle" || Opt.Screen_Resolution == "verticle" {
		banner_file = Banner_file3
	}
	if Opt.Screen_Resolution == "Landscape" || Opt.Screen_Resolution == "landscape" {
		banner_file = Banner_file2
	}
	f, x := os.Open(banner_file)
	if x != nil {
		log.Fatal(x)
	} else {
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			fmt.Println(color, scanner.Text())
		}
	}
}
