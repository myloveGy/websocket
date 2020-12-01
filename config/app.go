package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/jinxing-go/mysql"

	config2 "websocket/pkg/config"
)

type Redis struct {
	Addr     string `toml:"addr" json:"addr"`
	Password string `toml:"password" json:"password"`
	DB       int    `toml:"db" json:"db"`
	Prefix   string `toml:"prefix" json:"prefix"`
}

type Config struct {
	AppName     string `toml:"app_name"`
	StoragePath string `toml:"storage_path"`
	Debug       string `toml:"debug"`
	Token       string `toml:"token"`
	StaticUrl   string `toml:"static_url"` // 前端地址
	Env         string
	Path        string
	Address     string `toml:"address"` // 启用的端口
	StartTime   time.Time
	IsTesting   bool
	PrivateKey  []byte
	DB          map[string]*mysql.Config `toml:"database"` // 数据库信息
	Redis       map[string]*Redis        `toml:"redis"`
}

var App = &Config{
	StartTime: time.Now(),
}

func init() {
	ArgsInit()
	load(ExtCliArgs)
}

// 判断是否为测试执行
func isTestMode() bool {
	// test执行文件的路径后缀带.test，生产环境的可执行文件，不可能带.test后缀
	if strings.HasSuffix(os.Args[0], ".test") {
		return true
	}

	testVars := map[string]bool{
		"-test.v":   true,
		"-test.run": true,
	}

	for _, str := range os.Args {
		if testVars[str] {
			return true
		}
	}

	return false
}

func getAppPath(path string) string {

	if path == "" {
		// 由于go test执行路径是临时目录，因此寻找配置文件要从编译路径查找
		if isTestMode() {
			App.IsTesting = true
			_, file, _, _ := runtime.Caller(0)
			path = filepath.Dir(filepath.Dir(file))
		} else {
			path = "./"
		}
	}
	return path
}

func mustCheckError(err error) {
	if err != nil {
		log.Fatalf("config error %s", err.Error())
	}
}

func load(args map[string]string) {
	App.Path = getAppPath(args["config"])

	// load config
	conf, err := config2.LoadFile(filepath.Join(App.Path, "config.toml"))
	mustCheckError(err)

	// 加载 env 文件
	err = conf.Env(filepath.Join(App.Path, ".env"))
	mustCheckError(err)

	err = conf.Unmarshal(App)
	mustCheckError(err)

	if !filepath.IsAbs(App.StoragePath) {
		App.StoragePath = filepath.Join(App.Path, App.StoragePath)
	}
}
