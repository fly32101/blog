package config

// Config 应用配置
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port string
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	Charset  string
}

// NewConfig 创建默认配置
func NewConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Port: "8080",
		},
		Database: DatabaseConfig{
			Driver:   "mysql",
			Host:     "", // 使用本地数据库
			Port:     "",
			Username: "",
			Password: "", // 无密码
			DBName:   "",
			Charset:  "utf8mb4",
		},
	}
}

// GetDSN 获取数据库连接字符串
func (dc *DatabaseConfig) GetDSN() string {
	return dc.Username + ":" + dc.Password + "@tcp(" + dc.Host + ":" + dc.Port + ")/" + dc.DBName + "?charset=" + dc.Charset + "&parseTime=true&loc=Local"
}
