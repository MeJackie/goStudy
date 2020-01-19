package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
)

// 原始i/o操作
func read1(path string) {
	fi, err := os.Open(path)
	if err != err {
		panic(err)
	}
	defer fi.Close()
	buf := make([]byte, 1024)
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
	}
}

// bufio包i/o操作
func read2(path string) {
	fi, err := os.Open(path)
	if err != err {
		panic(err)
	}
	defer fi.Close()
	buf := make([]byte, 1024)
	r := bufio.NewReader(fi)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
	}
}

// ioutil包i/o操作
func read3(path string) {
	fi, err := os.Open(path)
	if err != err {
		panic(err)
	}
	defer fi.Close()
	_, err = ioutil.ReadAll(fi)
}

func main() {

}
