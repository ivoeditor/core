package core

type Command struct {
	Name    string
	Payload map[string]interface{}
}

func (cmd Command) String() string {
	return cmd.Name
}
