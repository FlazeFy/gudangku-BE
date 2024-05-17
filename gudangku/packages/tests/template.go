package tests

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	. "github.com/onsi/gomega"
)

func ValidateResponse(resp *resty.Response, err error) {
	Expect(err).To(BeNil())
	Expect(resp.StatusCode()).To(Equal(http.StatusOK))
	Expect(resp.Body()).NotTo(BeNil())
	fmt.Println(resp.String())
}
