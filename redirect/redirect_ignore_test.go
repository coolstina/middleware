package redirect

import "github.com/stretchr/testify/assert"

func (suite *RedirectSuite) Test_Ignore_URIIsIgnore_Want_False() {
	var ignore Ignore
	actual := ignore.URIIsIgnore("hello world")
	assert.Equal(suite.T(), false, actual)
}

func (suite *RedirectSuite) Test_Ignore_URIIsIgnore_Want_True() {
	var ignore = Ignore{"/users/helloshaohua/info": true}
	actual := ignore.URIIsIgnore("/users/helloshaohua/info")
	assert.Equal(suite.T(), true, actual)
}
