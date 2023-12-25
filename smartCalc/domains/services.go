package domains

type EqualModel struct {
	EqualString string
	Prepared    string
	XEqualStr   string
	ResultStr   string
	XEqual      float64
	Result      float64
	Equation    EquationModel
}

type EquationModel struct {
	Err       error
	Equation  string
	Checked   string
	Prepared  string
	ResultStr string
	Result    float64
}
