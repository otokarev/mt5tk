package cracker

type Command struct {
	arg      Argument
	proc     Proc
	result   Result
	error    error
	tries    int
	maxTries int
}

func NewCommand(proc Proc, arg Argument, maxTries int) Command {
	return Command{
		arg:      arg,
		proc:     proc,
		maxTries: maxTries,
	}
}

func (c *Command) Result() Result {
	return c.result
}

func (c *Command) Error() error {
	return c.error
}

func (c *Command) run(resource Resource) {
	c.tries++
	result, err := c.proc(resource, c.arg)
	if err != nil {
		c.error = err
	} else {
		c.error = nil
		c.result = result
	}
}

func (c *Command) shouldTryAgain() bool {
	if c.error != nil {
		if c.tries >= c.maxTries {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}
