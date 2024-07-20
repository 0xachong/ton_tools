package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/xssnick/tonutils-go/ton/wallet"
)

func main() {
	fmt.Print("请输入要生成助记词个数:")

	// 创建一个新的 bufio.Reader 读取标准输入
	reader := bufio.NewReader(os.Stdin)

	// 读取一行字符串，直到用户按下回车键
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	input = strings.ReplaceAll(input, "\n", "")
	input = strings.ReplaceAll(input, "\r", "")
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < num; i++ {
		mnemonic := wallet.NewSeed()
		fmt.Println(strings.Join(mnemonic, " "))
	}
	fmt.Println("按任意键退出...")
	// 创建一个新的 bufio.Reader 读取标准输入
	reader = bufio.NewReader(os.Stdin)
	// 读取一个字节，这将等待用户输入
	_, _ = reader.ReadByte()
	fmt.Println("Exiting...")
}