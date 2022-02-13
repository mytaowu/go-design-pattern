package _2_facade

import "fmt"

type Business struct{}

func (b *Business) Generate() {
	config := GetManagerInstance().GetConfigDao()
	if config.NeedGenBusiness {
		// 按照要求生成代码
		fmt.Println("正在生成逻辑层代码文件")
	}
}
