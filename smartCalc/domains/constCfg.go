package domains

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
)

type Cfg struct {
	WorkDir      string `json:"WorkDir"`      // Directory with programm
	AssetsDir    string `json:"AssetsDir"`    // Directory with help.txt and type
	LogDir       string `json:"LogDir"`       // Directory for log
	LogFile      string `json:"LogFile"`      // LogFile name
	TempFileDir  string `json:"TempFileDir"`  // Directory for Tempfile
	TempGraph    string `json:"TempGraph"`    // tempfile name
	HistoryFile  string `json:"HistoryFile"`  // historyfile path (with directory)
	XWindowGraph uint32 `json:"XWindowGraph"` // Graph window size X
	YWindowGraph uint32 `json:"YWindowGraph"` // Graph window size Y
	DarkTheme    string `json:"DarkTheme"`    //Dark mode "yes" or "no"
	IconPath     string `json:"IconPath"`     //iconfile name
	TypePath     string `json:"TypePath"`     //Typefile name
	Debug        int    `json:"Debug"`        //debug mode with output additional info to terminal -  true or false
}

var (
	Os      = runtime.GOOS   // "windows", "darwin", "linux"
	Arch    = runtime.GOARCH // "amd64", "386", "arm"
	WD, _   = os.Getwd()     // get workdirectory
	testDir = "../"

	// Path for config in linux system
	ConfigLinuxPath = []string{
		"/etc/smartCalc/smartCalcLinux.cfg",
		testDir + "/conf/smartCalcLinuxTest.cfg",
		"../conf/smartCalcLinuxTest.cfg",
	}

	// Path for config in Mac system
	ConfigMacPath = []string{
		testDir + "conf/smartCalcTest.cfg",
	}

	Config *Cfg = InitConfig("") // Handling config path by type config name in quotes (but this way not recommend)

	NeccessoryFiles = []string{
		Config.LogDir + Config.LogFile,
		Config.HistoryFile,
		Config.TempFileDir + Config.TempGraph,
		Config.TypePath,
		"./static/css/index_style.css",
		"./static/css/startCalculate_style.css",
		"./static/fonts/protosans56.ttf",
		"./static/img/background.png",
		"./static/img/backgroundMathw4.jpg",
		"./static/js/main.js",
		"./views/calculate/startCalculate.html",
		"./views/index.tpl",
	}
)

// Create and write new config
func createNewConfig() *Cfg {
	var c Cfg

	c.WorkDir = "."
	c.AssetsDir = c.WorkDir + "/assets/"
	c.LogDir = c.WorkDir + "/logs/"
	c.LogFile = "smartCalc.log"
	c.TempFileDir = c.WorkDir + "/static/tmp/"
	c.TempGraph = "tempGraph.png"
	c.HistoryFile = c.WorkDir + "/var/history.json"
	c.XWindowGraph = 600
	c.YWindowGraph = 600
	c.DarkTheme = "no"
	c.IconPath = c.AssetsDir + "Icon.png"
	c.TypePath = c.WorkDir + "/static/fonts/" + "protosans56.ttf"
	c.Debug = 3

	fmt.Println("Create config with inner data")

	return &c
}

// Read config file from disk
func readConfig(fileName string, c *Cfg) error {
	var err error
	if dataFromFile, err := os.ReadFile(fileName); err == nil {
		if err = json.Unmarshal(dataFromFile, c); err == nil {
			fmt.Println("Load config from:", fileName)
			return err
		}
	} else {
		return err
	}
	return err
}

// Inicialize config
func InitConfig(fileName string) *Cfg {
	var c Cfg

	// Try load config from handle path
	if fileName != "" {
		if err := readConfig(fileName, &c); err == nil {
			return &c
		}
	}

	// Try load config from installing programm path or test path
	switch Os {
	case "linux":
		for _, path := range ConfigLinuxPath {
			if err := readConfig(path, &c); err == nil {
				return &c
			}
		}

	case "darwin":
		for _, path := range ConfigMacPath {
			if err := readConfig(path, &c); err == nil {
				return &c
			}
		}
	}

	return createNewConfig()
}
