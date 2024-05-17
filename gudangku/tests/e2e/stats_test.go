package main

import (
	"fmt"
	"gudangku/packages/tests"
	"testing"

	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestStats(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GudangKu API Testing Suite - Stats")
}

var _ = Describe("GudangKu API Testing - Stats", func() {
	const method = "get"
	const token = "33|0DWfzepjZqA1Utxi3X9KQ40vcmKmZdJIatAJtmnq8d0f169f"
	const local_url = "http://127.0.0.1:1323"

	It(fmt.Sprintf("%s - Total Inventory By Category", method), func() {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			Get(local_url + "/api/v1/stats/total_inventory_by_category")

		tests.ValidateResponse(resp, err)
	})

	It(fmt.Sprintf("%s - Total Inventory By Favorite", method), func() {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			Get(local_url + "/api/v1/stats/total_inventory_by_favorite")

		tests.ValidateResponse(resp, err)
	})

	It(fmt.Sprintf("%s - Total Inventory By Room", method), func() {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			Get(local_url + "/api/v1/stats/total_inventory_by_room")

		tests.ValidateResponse(resp, err)
	})
})
