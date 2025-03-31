package app

import (
	"Qscan/lib/sflag"
	"fmt"
	"github.com/gookit/color"
	"os"
	"strings"
	"time"
)

type args struct {
	USAGE, HELP, LOGO, SYNTAX string

	Help, Debug, ClosePing, Check, CloseColor, Scan bool
	ScanVersion, DownloadQQwry, CloseCDN            bool
	Output, Proxy, Encoding                         string
	Port, ExcludedPort                              []int
	Path, Host, Target                              []string
	OutputJson, OutputCSV                           string
	Spy, Touch                                      string
	Top, Threads, Timeout                           int
	//hydra模块
	Hydra, HydraUpdate             bool
	HydraUser, HydraPass, HydraMod []string
	//exploit模块
	Exploit, DnsLog, PocFull bool
	Cookie                   string
	PocNum, WebTimeout       int
}

var (
	green  = []*color.Style256{color.S256(46), color.S256(47), color.S256(48), color.S256(49), color.S256(50), color.S256(51)}
	pink   = []*color.Style256{color.S256(214), color.S256(215), color.S256(216), color.S256(217), color.S256(218), color.S256(219)}
	yellow = []*color.Style256{color.S256(226), color.S256(227), color.S256(228), color.S256(229), color.S256(230), color.S256(231)}
)

var (
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36"
	Accept    = "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
)

type PocInfo struct {
	Target  string
	PocName string
}

type HostInfo struct {
	Host    string
	Ports   string
	Url     string
	Infostr []string
}

var Args = args{}

// Parse 初始化参数
func (o *args) Parse() {
	//自定义Usage
	sflag.SetUsage(o.LOGO)
	//定义参数
	o.define()
	//实例化参数值
	sflag.Parse()
	//输出LOGO
	o.printBanner()
}

// 定义参数
func (o *args) define() {
	sflag.BoolVar(&o.Help, "h", false)
	sflag.BoolVar(&o.Help, "help", false)
	sflag.BoolVar(&o.Debug, "debug", false)
	sflag.BoolVar(&o.Debug, "d", false)
	//spy模块
	sflag.AutoVarString(&o.Spy, "spy", "None")
	//hydra模块
	sflag.BoolVar(&o.Hydra, "hydra", false)
	sflag.BoolVar(&o.Hydra, "hY", false)
	sflag.BoolVar(&o.HydraUpdate, "hydra-update", false)
	sflag.StringSpliceVar(&o.HydraUser, "hydra-user")
	sflag.StringSpliceVar(&o.HydraPass, "hydra-pass")
	sflag.StringSpliceVar(&o.HydraMod, "hydra-mod")
	//exploit
	sflag.BoolVar(&o.Exploit, "eX", false)
	sflag.BoolVar(&o.Exploit, "exploit", false)
	sflag.StringVar(&o.Cookie, "cookie", "")
	sflag.IntVar(&o.PocNum, "num", 20)
	sflag.IntVar(&o.WebTimeout, "wt", 5)
	sflag.BoolVar(&o.DnsLog, "dns", false)
	sflag.BoolVar(&o.PocFull, "full", false)
	//kscan模块
	sflag.StringSpliceVar(&o.Target, "target")
	sflag.StringSpliceVar(&o.Target, "t")
	sflag.IntSpliceVar(&o.Port, "port")
	sflag.IntSpliceVar(&o.Port, "p")
	sflag.IntSpliceVar(&o.ExcludedPort, "eP")
	sflag.IntSpliceVar(&o.ExcludedPort, "excluded-port")
	sflag.StringSpliceVar(&o.Path, "path")
	sflag.StringSpliceVar(&o.Host, "host")
	sflag.StringVar(&o.Proxy, "proxy", "")
	sflag.IntVar(&o.Top, "top", 400)
	sflag.IntVar(&o.Threads, "threads", 800)
	sflag.IntVar(&o.Timeout, "timeout", 3)
	sflag.BoolVar(&o.ClosePing, "Pn", false)
	sflag.BoolVar(&o.Check, "check", false)
	sflag.BoolVar(&o.ScanVersion, "sV", false)
	//CDN检测
	sflag.BoolVar(&o.CloseCDN, "Dn", false)
	sflag.BoolVar(&o.DownloadQQwry, "download-qqwry", false)

	//输出模块
	sflag.StringVar(&o.Encoding, "encoding", "utf-8")
	sflag.StringVar(&o.Output, "o", "")
	sflag.StringVar(&o.Output, "output", "")
	sflag.StringVar(&o.OutputJson, "oJ", "")
	sflag.StringVar(&o.OutputCSV, "oC", "")
	sflag.BoolVar(&o.CloseColor, "Cn", false)
}

func (o *args) SetLogo(logo string) {
	o.LOGO = logo
}

func (o *args) SetUsage(usage string) {
	o.USAGE = usage
}

func (o *args) SetSyntax(syntax string) {
	o.SYNTAX = syntax
}

func (o *args) SetHelp(help string) {
	o.HELP = help
}

// CheckArgs 校验参数真实性
func (o *args) CheckArgs() {
	if len(o.Port) > 0 && o.Top != 400 {
		fmt.Print("--port、--top参数不能同时使用")
		os.Exit(0)
	}
	//判断内容
	if o.Top != 0 && (o.Top > 1000 || o.Top < 1) {
		fmt.Print("TOP参数输入错误,TOP参数应为1-1000之间的整数。")
		os.Exit(0)
	}
	if o.Proxy != "" && sflag.ProxyStrVerification(o.Proxy) {
		fmt.Print("--proxy参数输入错误，其格式应为：http://ip:port，支持socks5/4")
	}
	if o.Threads != 0 && o.Threads > 2048 {
		fmt.Print("--threads参数最大值为2048")
		os.Exit(0)
	}
}

// 输出LOGO
func (o *args) printBanner() {
	if len(os.Args) == 1 {
		fmt.Println(gradient(o.LOGO, pink))
		fmt.Print(o.USAGE)
		os.Exit(0)
	}
	if o.Help {
		fmt.Println(gradient(o.LOGO, pink))
		fmt.Print(o.USAGE)
		fmt.Print(o.HELP)
		os.Exit(0)
	}
	//打印logo
	fmt.Println(gradient(o.LOGO, pink))
}

func gradient(text string, coloRR []*color.Style256) string {
	lines := strings.Split(text, "\n")

	var output string

	t := len(text) / 6
	i := 0
	j := 0
	for l := 0; l < len(lines); l++ {
		str := strings.Split(lines[l], "")
		for _, x := range str {
			j++
			output += coloRR[i].Sprint(x)
			if j > t {
				i++
				j = 0
			}
		}
		if len(lines) != 0 {
			output += "\n"
		}
	}

	return strings.TrimRight(output, "\n")
}

// CheckErrs 检查是否为需要重试的错误
func CheckErrs(err error) error {
	if err == nil {
		return nil
	}

	// 已知需要重试的错误列表
	errs := []string{
		"closed by the remote host", "too many connections",
		"EOF", "A connection attempt failed",
		"established connection failed", "connection attempt failed",
		"Unable to read", "is not allowed to connect to this",
		"no pg_hba.conf entry",
		"No connection could be made",
		"invalid packet size",
		"bad connection",
	}

	// 检查错误是否匹配
	errLower := strings.ToLower(err.Error())
	for _, key := range errs {
		if strings.Contains(errLower, strings.ToLower(key)) {
			time.Sleep(3 * time.Second)
			return err
		}
	}

	return nil
}

var Userdict = map[string][]string{
	"elastic":   {"elastic", "admin", "kibana"},
	"rabbitmq":  {"guest", "admin", "administrator", "rabbit", "rabbitmq", "root"},
	"kafka":     {"admin", "kafka", "root", "test"},
	"activemq":  {"admin", "root", "activemq", "system", "user"},
	"ldap":      {"admin", "administrator", "root", "cn=admin", "cn=administrator", "cn=manager"},
	"smtp":      {"admin", "root", "postmaster", "mail", "smtp", "administrator"},
	"imap":      {"admin", "mail", "postmaster", "root", "user", "test"},
	"pop3":      {"admin", "root", "mail", "user", "test", "postmaster"},
	"zabbix":    {"Admin", "admin", "guest", "user"},
	"rsync":     {"rsync", "root", "admin", "backup"},
	"cassandra": {"cassandra", "admin", "root", "system"},
	"neo4j":     {"neo4j", "admin", "root", "test"},
}

var Passwords = []string{"123456", "admin", "admin123", "root", "", "pass123", "pass@123", "password", "Password", "P@ssword123", "123123", "654321", "111111", "123", "1", "admin@123", "Admin@123", "admin123!@#", "{user}", "{user}1", "{user}111", "{user}123", "{user}@123", "{user}_123", "{user}#123", "{user}@111", "{user}@2019", "{user}@123#4", "P@ssw0rd!", "P@ssw0rd", "Passw0rd", "qwe123", "12345678", "test", "test123", "123qwe", "123qwe!@#", "123456789", "123321", "666666", "a123456.", "123456~a", "123456!a", "000000", "1234567890", "8888888", "!QAZ2wsx", "1qaz2wsx", "abc123", "abc123456", "1qaz@WSX", "a11111", "a12345", "Aa1234", "Aa1234.", "Aa12345", "a123456", "a123123", "Aa123123", "Aa123456", "Aa12345.", "sysadmin", "system", "1qaz!QAZ", "2wsx@WSX", "qwe123!@#", "Aa123456!", "A123456s!", "sa123456", "1q2w3e", "Charge123", "Aa123456789", "elastic123"}
