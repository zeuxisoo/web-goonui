package setting

import (
    "os"
    "os/exec"
    "path/filepath"
    "net/url"
    "strings"

    "gopkg.in/ini.v1"
    "github.com/go-macaron/session"

    "github.com/zeuxisoo/go-goonui/app/kernels/log"
    "github.com/zeuxisoo/go-goonui/app/helpers"
)

var (
    AppPath         string

    Address         string
    Port            string

    AppUrl          string
    AppSubUrl       string
    AppSubUrlDepth  int

    StaticRootPath  string

    SecretKey       string

    SessionConfig   session.Options

    CsrfCookieName  string

    Configure       *ini.File
)

func init() {
    var err error

    if AppPath, err = executablePath(); err != nil {
        log.Fatalf("Fail to get app path: %v\n", err)
    }

    AppPath = strings.Replace(AppPath, "\\", "/", -1)
}

func executablePath() (string, error) {
    file, err := exec.LookPath(os.Args[0])

    if err != nil {
        return "", err
    }

    return filepath.Abs(file)
}

func AppDirectory() (string, error) {
    i := strings.LastIndex(AppPath, "/")

    if i == -1 {
        return AppPath, nil
    }

    return AppPath[:i], nil
}

func NewSetting() {
    // Load configuration file
    appDirectory, err := AppDirectory()

    if err != nil {
        log.Fatalf("Fail to get application directory: %v", err)
    }

    Configure, err = ini.Load("conf/app.ini")

    if err != nil {
        log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
    }

    // Server configuration
    sectionServer := Configure.Section("server")

    // => Set AppUrl
    AppUrl = sectionServer.Key("ROOT_URL").MustString("http://localhost:8080/")

    if AppUrl[len(AppUrl)-1] != '/' {
        AppUrl += "/"
    }

    url, err := url.Parse(AppUrl)

    if err != nil {
        log.Fatalf("Invalid ROOT_URL '%s': %s", AppUrl, err)
    }

    AppSubUrl      = strings.TrimSuffix(url.Path, "/")
    AppSubUrlDepth = strings.Count(AppSubUrl, "/")

    // => Set others
    Address        = sectionServer.Key("ADDRESS").MustString("127.0.0.1")
    Port           = sectionServer.Key("PORT").MustString("8080")
    StaticRootPath = sectionServer.Key("STATIC_ROOT_PATH").MustString(appDirectory)

    // Security configuration
    sectionSecurity := Configure.Section("security")

    SecretKey      = sectionSecurity.Key("SECRET_KEY").MustString(helpers.GetRandomString(15))
    CsrfCookieName = sectionSecurity.Key("CSRF_COOKIE_NAME").MustString("_csrf")
}

func NewSessionService() {
    sessionSection := Configure.Section("session")

    providerConfig := sessionSection.Key("PROVIDER_CONFIG").String()
    providerConfig  = strings.Trim(providerConfig, "\" ")

    SessionConfig.Provider       = sessionSection.Key("PROVIDER").In("memory", []string{"memory", "file", "redis", "mysql"})
    SessionConfig.ProviderConfig = providerConfig
    SessionConfig.CookieName     = sessionSection.Key("COOKIE_NAME").MustString("goonui")
    SessionConfig.CookiePath     = AppSubUrl
    SessionConfig.Secure         = sessionSection.Key("COOKIE_SECURE").MustBool()
    SessionConfig.Gclifetime     = sessionSection.Key("GC_INTERVAL_TIME").MustInt64(86400)
    SessionConfig.Maxlifetime    = sessionSection.Key("SESSION_LIFE_TIME").MustInt64(86400)
}
