package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	pat := "*/*.env"
	if len(os.Args) > 1 {
		pat = os.Args[1]
	}
	envfiles, err := filepath.Glob(pat)
	if err != nil {
		panic(err)
	}
	// fmt.Println(envfiles)

	for _, fn := range envfiles {
		f, err := os.Open(fn)
		if err != nil {
			panic(err)
		}

		fmt.Println("Reading", fn)
		b, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}

		str := string(b)
		lines := strings.Split(str, "\n")
		newfilename := strings.Replace(fn, ".", "_", -1)
		fmt.Println("Writing", newfilename)
		newf, err := os.OpenFile(newfilename, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			panic(err)
		}

		for _, l := range lines {
			segs := strings.Split(strings.Replace(l, `"`, "", -1), "=")
			if len(segs) < 2 {
				fmt.Fprintln(newf, "")
				continue
			}
			fmt.Fprintf(newf, `export %s="%s"`, segs[0], strings.Join(segs[1:], "="))
			fmt.Fprintln(newf)
		}
		err = newf.Close()
		if err != nil {
			panic(err)
		}
	}
}
