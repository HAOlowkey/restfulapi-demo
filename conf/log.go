package conf

type LogFormat string

const (
	TextFormat LogFormat = "text"
	JsonFormat LogFormat = "json"
)

type LogTo string

const (
	ToFile   LogTo = "file"
	ToStdout LogTo = "stdout"
)
