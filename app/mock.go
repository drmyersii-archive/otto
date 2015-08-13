package app

// Mock is a mock implementation of the App interface.
type Mock struct {
	CompileCalled  bool
	CompileContext *Context
	CompileResult  *CompileResult
	CompileErr     error

	DevCalled  bool
	DevContext *Context
	DevErr     error
}

func (m *Mock) Compile(ctx *Context) (*CompileResult, error) {
	m.CompileCalled = true
	m.CompileContext = ctx
	return m.CompileResult, m.CompileErr
}

func (m *Mock) Dev(ctx *Context) error {
	m.DevCalled = true
	m.DevContext = ctx
	return m.DevErr
}