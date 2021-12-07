package redirect

import (
	"github.com/stretchr/testify/assert"
)

func (suite *RedirectSuite) Test_Options() {
	options := &options{}

	ops := []Option{
		WithSetIgnoreURI("/users/helloshaohua/info", true),
		WithSetIgnoreURI("/users/helloshaohua/status", true),
	}

	for _, o := range ops {
		o.apply(options)
	}
	assert.Len(suite.T(), options.ignore, 2)
}
