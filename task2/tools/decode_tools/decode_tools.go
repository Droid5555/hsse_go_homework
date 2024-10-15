package decode_tools

type Input struct {
	Request string `json:"inputString"`
}

type Output struct {
	Response string `json:"outputString"`
}

func NewInput(s string) *Input {
	d := &Input{}
	d.Request = s
	return d
}
