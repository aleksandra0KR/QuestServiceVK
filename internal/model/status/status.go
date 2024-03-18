package status

type Status string

const (
	Done      Status = "done"
	InProcess Status = "in process"
	Available Status = "available"
)
