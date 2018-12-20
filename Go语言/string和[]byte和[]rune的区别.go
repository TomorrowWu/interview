package main

import "fmt"

func printStirng(s string) {
	fmt.Print("\nprintStirng:")
	for _, char := range s {
		fmt.Printf("%d,", char)
	}
}

func printBytes(s string) {
	fmt.Print("\nprintBytes:")
	bs := []byte(s)
	for _, b := range bs {
		fmt.Printf("%v,", b)
	}
}

func printRunes(s string) {
	fmt.Print("\nprintRunes:")
	rs := []rune(s)
	for _, r := range rs {
		fmt.Printf("%v,", r)
	}
}

func main() {
	s1 := "中文"
	printStirng(s1) //20013,25991
	printBytes(s1)  //228,184,173,230,150,135,
	printRunes(s1)  //20013,25991,

	s2 := "abc"
	printStirng(s2) //97,98,99,
	printBytes(s2)  //97,98,99,
	printRunes(s2)  //97,98,99,

	//遍历string相当于遍历[]int32,等同于rune（int32）,int32相当于4字节的长度，中文英文字符都在长度内
	//处理中英文时不同，英文，一个字符即一个字节，中文的UTF-8是一个字符3个字节
}
