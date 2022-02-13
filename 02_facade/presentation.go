package _2_facade

import "fmt"

type Presentation struct{}

func (p *Presentation) Generate() {
	config := GetManagerInstance().GetConfigDao()
	if config.NeedGenPresentation {
		// 按照要求生成代码
		fmt.Println("正在生成表现层代码文件")
	}
}
