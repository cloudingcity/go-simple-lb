package server

import (
	"bytes"
	"net"
	"net/url"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type ControllerTestSuite struct {
	suite.Suite
	u1  *url.URL
	u2  *url.URL
	u3  *url.URL
	out *bytes.Buffer
}

func (suite *ControllerTestSuite) SetupTest() {
	u1, _ := url.Parse("http://localhost:1234")
	u2, _ := url.Parse("http://localhost:1235")
	u3, _ := url.Parse("http://localhost:1236")
	suite.u1 = u1
	suite.u2 = u2
	suite.u3 = u3

	suite.out = &bytes.Buffer{}
	logrus.SetOutput(suite.out)
}

func (suite *ControllerTestSuite) TestGetNext() {
	c := NewController()
	c.SetupServers(suite.u1, suite.u2, suite.u3)

	suite.Equal(suite.u1, c.getNext().url)
	suite.Equal(suite.u2, c.getNext().url)
	suite.Equal(suite.u3, c.getNext().url)
}

func (suite *ControllerTestSuite) TestGetNextFail() {
	c := NewController()
	suite.Nil(c.getNext())
}

func (suite *ControllerTestSuite) TestHealthCheck() {
	c := NewController()
	c.SetupServers(suite.u1, suite.u2, suite.u3)

	c.HealthCheck()
	suite.Equal(0, c.upIDs.Len())
	suite.Equal(3, c.downIDs.Len())

	ln, err := net.Listen("tcp4", "localhost:1234")
	suite.NoError(err)

	c.HealthCheck()
	suite.Equal(1, c.upIDs.Len())
	suite.Equal(2, c.downIDs.Len())

	ln.Close()
}

func (suite *ControllerTestSuite) TestDown() {
	c := NewController()
	c.SetupServers(suite.u1, suite.u2, suite.u3)

	c.down(1)
	suite.Equal(2, c.upIDs.Len())
	suite.Equal(1, c.downIDs.Len())
	suite.Contains(suite.out.String(), "[http://localhost:1234] down")
}

func TestControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ControllerTestSuite))
}
