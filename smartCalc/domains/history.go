package domains

import "fmt"

type HistoryItem struct {
	Mode     string `json:"mode"`
	Equation string `json:"equation"`
	Result   string `json:"result"`
	Entrys   string `json:"entrys"`

	XEqual string `json:"xEqual"`
	XFrom  string `json:"xFrom"`
	XTo    string `json:"xTo"`
	YFrom  string `json:"yFrom"`
	YTo    string `json:"yTo"`
}

const debug = false

func DbgPrint(s string) {
	if debug {
		fmt.Println(s)
	}
}
