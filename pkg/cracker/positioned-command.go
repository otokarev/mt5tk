package cracker

type positionedCommand struct {
	Pos int
	Command
}

type commandArray []positionedCommand

func newCommandArray(cmds []Command) commandArray {
	r := make([]positionedCommand, len(cmds))

	for i, cmd := range cmds {
		r[i].Pos = i
		r[i].Command = cmd
	}

	return r
}

func (a commandArray) toCommandSlice() []Command {
	r := make([]Command, len(a))

	for _, pc := range a {
		r[pc.Pos] = pc.Command
	}

	return r
}
