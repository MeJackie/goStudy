package main

import (
	//"flag"
	"fmt"
	"sync"

	//"github.com/gopl.io/ch5/links"
	//"log"
	//"time"

	//"github.com/gopl.io/ch8/thumbnail"
	"os"
)



func main() {
	ch := make(chan int, 10)
	sync.Mutex

	b:= <-ch
	fmt.Println(b)
	//ch <- 1
	//close(ch)
	//ch <- 2
	//select {
	//case c:=<-ch:
	//	fmt.Println("获取数据", c)
	//default:
	//	b:= <-ch
	//	fmt.Println("default", b)
	//}
	os.Exit(1)
	//workList := make(chan []string)
	//unseekLists := make(chan string)
	//
	//// 消费为爬取连接
	//for i := 0; i < 20; i++ {
	//	go func() {
	//		for link := range unseekLists {
	//			findLink := crawl(link)
	//			go func() {workList <- findLink}()  // 避免死锁
	//		}
	//	}()
	//}
	//
	//// 查找为爬取连接
	//seek := make(map[string]bool)
	//for list := range workList {
	//	for _, link := range list {
	//		if !seek[link] {
	//			seek[link] = true
	//			unseekLists <- link
	//		}
	//	}
	//}
	//
	//
	//var n int
	//
	//n++
	//go func() {worklist <- os.Args[1:]}()
	//
	//fmt.Println(n)
	//fmt.Println(<-worklist)
}



//
//func makeThumbnails(filenames []string) {
//	ch := make(chan struct{})
//	for _, f := range filenames {
//		go func(f string) {
//			thumbnail.ImageFile(f)
//			ch <- struct{}{}
//		}(f)
//	}
//
//	// len(filenames)个数据
//	for range filenames {
//		<-ch
//	}
//}
