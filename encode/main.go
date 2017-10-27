package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"code.google.com/p/mahonia"
)

var (
	infile  = "in.txt"
	outfile = "out.txt"
	src     = "utf8"
	dst     = "tcvn3"
)

func main() {
	flag.StringVar(&infile, "file", infile, "file")
	flag.StringVar(&src, "src", src, "src")
	flag.StringVar(&dst, "dst", dst, "dst")
	flag.StringVar(&outfile, "out", outfile, "out")
	flag.Parse()

	decoder := mahonia.NewDecoder(src)
	if decoder == nil {
		fmt.Println("unknow src code ", src)
		return
	}
	encoder := mahonia.NewEncoder(dst)
	if encoder == nil {
		fmt.Println("unknow src code ", dst)
		return
	}

	f, err := os.Open(infile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, bytes2, err := decoder.Translate(bytes, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	strout := encoder.ConvertString(string(bytes2))

	f2, err := os.OpenFile(outfile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f2.Close()

	f2.Write([]byte(strout))
}
