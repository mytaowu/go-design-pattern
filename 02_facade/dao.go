package _2_facade

import "fmt"

type DAO struct{}

func (d *DAO) Generate() {
	config := GetManagerInstance().GetConfigDao()
	if config.NeedGenDAO {
		// 按照要求生成代码
		fmt.Println("正在生成dao层代码文件")
	}
}
