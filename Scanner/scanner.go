package Scanner

import (
	"bufio"
	"fmt"
	"os"
)
/* 每次从标准输入读取一行 */
func ScannerFromStdin() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		fmt.Printf("%d:%s", n, line)
	}
}

/* 每次从文件读取一行
 * 优化点data = ioutil.ReadFile(filename)是一次性全部读入内存
 * 使用strings.Split(string(data), "\n")切割
 */
func ScannerFromFile() error {
	if f, err := os.Open("filename"); err == nil {
		defer f.Close()
		input := bufio.NewScanner(f)
		for input.Scan() {
			fmt.Println(input.Text())
		}
	}
	return nil
}
/* 编程方法
 * 1.bufio.NewScanner构造新的结构体(对象)
 * 2.Scan()移动游标
 * 3.Text()方法获得Scan结果
 * 扩展：有点类似编译原理提到的词法分析器,每次按照正则获得下一个token
 */
