package decode_tools

type Input struct {
	InputString string `json:"inputString"`
}

type Output struct {
	OutputString string `json:"outputString"`
}

func NewInput(s string) *Input {
	d := &Input{}
	d.InputString = s
	return d
}
