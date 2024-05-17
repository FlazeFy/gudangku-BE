package main

import (
	"fmt"
	"gudangku/packages/tests"
	"testing"

	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHistory(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GudangKu API Testing Suite - History")
}

var _ = Describe("GudangKu API Testing - History", func() {
	const method = "get"
	const token = "33|0DWfzepjZqA1Utxi3X9KQ40vcmKmZdJIatAJtmnq8d0f169f"
	const local_url = "http://127.0.0.1:1323"

	It(fmt.Sprintf("%s - All History", method), func() {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			Get(local_url + "/api/v1/history")

		tests.ValidateResponse(resp, err)
	})
})
