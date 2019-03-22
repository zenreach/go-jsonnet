package iporter

func codeToPV(i *interpreter, filename string, code string) *cachedThunk {
	node, err := snippetToAST(filename, code)
	if err != nil {
		// TODO(sbarzowski) we should wrap (static) error here
		// within a RuntimeError. Because whether we get this error or not
		// actually depends on what happens in Runtime (whether import gets
		// evaluated).
		// The same thinking applies to external variables.
		return &cachedThunk{err: err}
	}
	env := makeInitialEnv(filename, i.baseStd)
	return &cachedThunk{
		env:     &env,
		body:    node,
		content: nil,
	}
}
