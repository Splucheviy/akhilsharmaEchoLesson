package handlers_test

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type EndtoEndSuite struct {
	suite.Suite
}

func TestEndtoEndSuite(t *testing.T) {
	suite.Run(t, new(EndtoEndSuite))
}

func (s *EndtoEndSuite) TestPostHandler() {
	c := http.Client{}
	r, _ := c.Get("http://localhost:8080/posts")
	s.Equal(http.StatusOK, r.StatusCode)
}

func (s *EndtoEndSuite) TestPostNoResult() {
	c := http.Client{}
	r, _ := c.Get("http://localhost:8080/post/1000")
	s.Equal(http.StatusOK, r.StatusCode)
	b, _ := io.ReadAll(r.Body)
	s.JSONEq(`{"status" : "ok", "data":[]}`, string(b))

}
