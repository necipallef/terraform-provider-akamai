package botman

import (
	"testing"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/v7/pkg/botman"
	"github.com/akamai/terraform-provider-akamai/v5/pkg/test"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stretchr/testify/mock"
)

func TestDataResponseAction(t *testing.T) {
	t.Run("DataResponseAction", func(t *testing.T) {

		mockedBotmanClient := &botman.Mock{}
		response := botman.GetResponseActionListResponse{
			ResponseActions: []map[string]interface{}{
				{"actionId": "b85e3eaa-d334-466d-857e-33308ce416be", "testKey": "testValue1"},
				{"actionId": "69acad64-7459-4c1d-9bad-672600150127", "testKey": "testValue2"},
				{"actionId": "cc9c3f89-e179-4892-89cf-d5e623ba9dc7", "testKey": "testValue3"},
				{"actionId": "10c54ea3-e3cb-4fc0-b0e0-fa3658aebd7b", "testKey": "testValue4"},
				{"actionId": "4d64d85a-a07f-485a-bbac-24c60658a1b8", "testKey": "testValue5"},
			},
		}
		expectedJSON := `
{
	"responseActions":[
		{"actionId":"b85e3eaa-d334-466d-857e-33308ce416be", "testKey":"testValue1"},
		{"actionId":"69acad64-7459-4c1d-9bad-672600150127", "testKey":"testValue2"},
		{"actionId":"cc9c3f89-e179-4892-89cf-d5e623ba9dc7", "testKey":"testValue3"},
		{"actionId":"10c54ea3-e3cb-4fc0-b0e0-fa3658aebd7b", "testKey":"testValue4"},
		{"actionId":"4d64d85a-a07f-485a-bbac-24c60658a1b8", "testKey":"testValue5"}
	]
}`
		mockedBotmanClient.On("GetResponseActionList",
			mock.Anything,
			botman.GetResponseActionListRequest{ConfigID: 43253, Version: 15},
		).Return(&response, nil)

		useClient(mockedBotmanClient, func() {

			resource.Test(t, resource.TestCase{
				IsUnitTest:        true,
				ProviderFactories: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: test.Fixture("testdata/TestDataResponseAction/basic.tf"),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttr("data.akamai_botman_response_action.test", "json", compactJSON(expectedJSON))),
					},
				},
			})
		})

		mockedBotmanClient.AssertExpectations(t)
	})
	t.Run("DataResponseAction filter by id", func(t *testing.T) {

		mockedBotmanClient := &botman.Mock{}
		response := botman.GetResponseActionListResponse{
			ResponseActions: []map[string]interface{}{
				{"actionId": "cc9c3f89-e179-4892-89cf-d5e623ba9dc7", "testKey": "testValue3"},
			},
		}
		expectedJSON := `
{
	"responseActions":[
		{"actionId":"cc9c3f89-e179-4892-89cf-d5e623ba9dc7", "testKey":"testValue3"}
	]
}`
		mockedBotmanClient.On("GetResponseActionList",
			mock.Anything,
			botman.GetResponseActionListRequest{ConfigID: 43253, Version: 15, ActionID: "cc9c3f89-e179-4892-89cf-d5e623ba9dc7"},
		).Return(&response, nil)

		useClient(mockedBotmanClient, func() {

			resource.Test(t, resource.TestCase{
				IsUnitTest:        true,
				ProviderFactories: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: test.Fixture("testdata/TestDataResponseAction/filter_by_id.tf"),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttr("data.akamai_botman_response_action.test", "json", compactJSON(expectedJSON))),
					},
				},
			})
		})

		mockedBotmanClient.AssertExpectations(t)
	})
}
