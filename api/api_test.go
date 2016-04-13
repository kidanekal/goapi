package api_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"encoding/json"
	"net/http"

	"github.com/kidanekal/goapi/api"
)

var _ = Describe("Api", func() {

	Context("Health check", func() {

		It("GET returns 200", func() {

			resp, err := http.Get(server.URL + "/health")

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
		})

		It("HEAD returns 405", func() {

			resp, err := http.Head(server.URL + "/health")

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(405))
		})
	})
})
