package config

var Config *Conf

// Conf 配置结构体
type Conf struct {
	Mysql        MysqlConfig        // MySQL数据库配置
	Redis        RedisConfig        // Redis配置
	TestNet      TestNetConfig      // 测试网络配置
	MainNet      MainNetConfig      // 主网配置
	Token        TokenConfig        // Token配置
	Email        EmailConfig        // 邮件配置
	DefaultAdmin DefaultAdminConfig // 默认管理员配置
	Threshold    ThresholdConfig    // 阈值配置
	Jwt          JwtConfig          // JWT配置
	Env          EnvConfig          // 环境配置
}

// EnvConfig 环境配置结构体
type EnvConfig struct {
	Port               string `toml:"port"`               // 服务端口
	Version            string `toml:"version"`            // 版本号
	Protocol           string `toml:"protocol"`           // 协议
	DomainName         string `toml:"domain_name"`        // 域名
	TaskDuration       int64  `toml:"task_duration"`      // 任务持续时间
	WssTimeoutDuration int64  `toml:"wss_timeout_duration"` // WSS超时时间
	TaskExtendDuration int64  `toml:"task_extend_duration"` // 任务延长时间
}

// ThresholdConfig 阈值配置结构体
type ThresholdConfig struct {
	PledgePoolTokenThresholdBnb string `toml:"pledge_pool_token_threshold_bnb"` // BNB质押池token阈值
}

// EmailConfig 邮件配置结构体
type EmailConfig struct {
	Username string   `toml:"username"` // 邮箱用户名
	Pwd      string   `toml:"pwd"`      // 邮箱密码
	Host     string   `toml:"host"`     // 邮箱服务器地址
	Port     string   `toml:"port"`     // 邮箱服务器端口
	From     string   `toml:"from"`     // 发件人地址
	Subject  string   `toml:"subject"`  // 邮件主题
	To       []string `toml:"to"`       // 收件人列表
	Cc       []string `toml:"cc"`       // 抄送人列表
}

// DefaultAdminConfig 默认管理员配置结构体
type DefaultAdminConfig struct {
	Username string `toml:"username"` // 管理员用户名
	Password string `toml:"password"` // 管理员密码
}

// JwtConfig JWT配置结构体
type JwtConfig struct {
	SecretKey  string `toml:"secret_key"`  // JWT密钥
	ExpireTime int    `toml:"expire_time"` // JWT过期时间（秒）
}

// TokenConfig Token配置结构体
type TokenConfig struct {
	LogoUrl string `toml:"logo_url"` // Logo图片URL
}

// MysqlConfig MySQL配置结构体
type MysqlConfig struct {
	Address      string `toml:"address"`       // 数据库地址
	Port         string `toml:"port"`          // 数据库端口
	DbName       string `toml:"db_name"`       // 数据库名称
	UserName     string `toml:"user_name"`     // 数据库用户名
	Password     string `toml:"password"`      // 数据库密码
	MaxOpenConns int    `toml:"max_open_conns"` // 最大打开连接数
	MaxIdleConns int    `toml:"max_idle_conns"` // 最大空闲连接数
	MaxLifeTime  int    `toml:"max_life_time"`  // 连接最大生命周期
}

// TestNetConfig 测试网络配置结构体
type TestNetConfig struct {
	ChainId              string `toml:"chain_id"`               // 链ID
	NetUrl               string `toml:"net_url"`                // 网络URL
	PlgrAddress          string `toml:"plgr_address"`          // PLGR地址
	PledgePoolToken      string `toml:"pledge_pool_token"`      // 质押池Token地址
	BscPledgeOracleToken string `toml:"bsc_pledge_oracle_token"` // BSC质押预言机Token地址
}

// MainNetConfig 主网配置结构体
type MainNetConfig struct {
	ChainId              string `toml:"chain_id"`               // 链ID
	NetUrl               string `toml:"net_url"`                // 网络URL
	PlgrAddress          string `toml:"plgr_address"`          // PLGR地址
	PledgePoolToken      string `toml:"pledge_pool_token"`      // 质押池Token地址
	BscPledgeOracleToken string `toml:"bsc_pledge_oracle_token"` // BSC质押预言机Token地址
}

// RedisConfig Redis配置结构体
type RedisConfig struct {
	Address     string `toml:"address"`      // Redis服务器地址
	Port        string `toml:"port"`         // Redis服务器端口
	Db          int    `toml:"db"`           // Redis数据库编号
	Password    string `toml:"password"`     // Redis密码
	MaxIdle     int    `toml:"max_idle"`     // 最大空闲连接数
	MaxActive   int    `toml:"max_active"`   // 最大活动连接数
	IdleTimeout int    `toml:"idle_timeout"` // 空闲超时时间
}
