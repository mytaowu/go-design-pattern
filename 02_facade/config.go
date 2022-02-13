package _2_facade

// ConfigModel 配置类
type ConfigModel struct {
	// NeedGenPresentation 是否生成表现层
	NeedGenPresentation bool
	// NeedGenBusiness 是否生成逻辑层
	NeedGenBusiness bool
	// NeedGenDAO 是否生成DAO层
	NeedGenDAO bool
}
