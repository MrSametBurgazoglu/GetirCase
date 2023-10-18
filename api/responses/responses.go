package responses

type ResponseTypes uint16

const (
	Success ResponseTypes = iota
	ParseJsonError
	ValidationError
)

var ResponseMassages = map[ResponseTypes]string{
	Success:        "success",
	ParseJsonError: "Error occured when parsing request json",
}

func (receiver ResponseTypes) Message() string {
	return ResponseMassages[receiver]
}
