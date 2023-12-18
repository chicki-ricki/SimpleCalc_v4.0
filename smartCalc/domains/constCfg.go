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
	TempFileDir  string `json:"TempFileDir"`  // Directory for Tempfile
	TempGraph    string `json:"TempGraph"`    // tempfile name
	HistoryFile  string `json:"HistoryFile"`  // historyfile path (with directory)
	XWindowGraph uint32 `json:"XWindowGraph"` // Graph window size X
	YWindowGraph uint32 `json:"YWindowGraph"` // Graph window size Y
	DarkTheme    string `json:"DarkTheme"`    //Dark mode "yes" or "no"
	IconPath     string `json:"IconPath"`     //iconfile name
	TypePath     string `json:"TypePath"`     //Typefile name
	Debug        bool   `json:"Debug"`        //debug mode with output additional info to terminal -  true or false
}

var (
	Os      = runtime.GOOS   // "windows", "darwin", "linux"
	Arch    = runtime.GOARCH // "amd64", "386", "arm"
	WD, _   = os.Getwd()     // get workdirectory
	testDir = testDirFind(WD)

	// Path for config in linux system
	// [0] - main; [1] - optional
	ConfigLinuxPath = []string{
		"/etc/clevercalc/clevercalcLinux.cfg",
		testDir + "/config/clevercalcLinuxTest.cfg"}

	// Path for config in Mac system
	// [0] - main; [1] - for test
	ConfigMacPath = []string{
		"/Applications/clevercalc.app/Contents/Resources/config/clevercalcMacOS.cfg",
		testDir + "/config/clevercalcMacOSTest.cfg"}

	Config *Cfg = InitConfig("") // Handling config path by type config name in quotes (but this way not recommend)
)

// Create and write new config for Mac
func createNewMacConfig() *Cfg {
	var c Cfg

	c.WorkDir = WD
	c.AssetsDir = c.WorkDir + "/assets/"
	c.LogDir = c.WorkDir + "/log/"
	c.TempFileDir = c.WorkDir + "/tmp/"
	c.TempGraph = "tempGraph.png"
	c.HistoryFile = c.WorkDir + "/var/history.json"
	c.XWindowGraph = 600
	c.YWindowGraph = 600
	c.DarkTheme = "no"
	c.IconPath = c.AssetsDir + "Icon.png"
	c.TypePath = c.AssetsDir + "protosans56.ttf"
	c.Debug = false

	fmt.Println("Create config with inner data")

	return &c
}

// Create and write new config for Linux
func createNewLinuxConfig() *Cfg {
	var c Cfg

	c.WorkDir = WD
	c.AssetsDir = c.WorkDir + "/assets/"
	c.LogDir = c.WorkDir + "/log/"
	c.TempFileDir = c.WorkDir + "/tmp/"
	c.TempGraph = "tempGraph.png"
	c.HistoryFile = c.WorkDir + "/var/history.json"
	c.XWindowGraph = 600
	c.YWindowGraph = 600
	c.DarkTheme = "yes"
	c.IconPath = c.AssetsDir + "Icon.png"
	c.TypePath = c.AssetsDir + "protosans56.ttf"
	c.Debug = false

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

	// fmt.Println(WD)
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

		// Create New config from inner information if load unsuccesful
		return createNewLinuxConfig()

	case "darwin":
		for _, path := range ConfigMacPath {
			if err := readConfig(path, &c); err == nil {
				return &c
			}
		}
		// Create New config from inner information if load unsuccesful
		return createNewMacConfig()
	}

	return createNewLinuxConfig()
}

func testDirFind(WD string) string {
	if len(WD) > 20 {
		return WD[:len(WD)-19]
	}
	return WD
}