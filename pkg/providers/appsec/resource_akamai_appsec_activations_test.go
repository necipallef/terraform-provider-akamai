package appsec

import (
	"encoding/json"
	"testing"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/v7/pkg/appsec"
	"github.com/akamai/terraform-provider-akamai/v5/pkg/common/testutils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAkamaiActivations_res_basic(t *testing.T) {
	t.Run("match by Activations ID", func(t *testing.T) {
		client := &appsec.Mock{}

		removeActivationsResponse := appsec.RemoveActivationsResponse{}
		err := json.Unmarshal(testutils.LoadFixtureBytes(t, "testdata/TestResActivations/ActivationsDelete.json"), &removeActivationsResponse)
		require.NoError(t, err)

		getActivationsResponse := appsec.GetActivationsResponse{}
		err = json.Unmarshal(testutils.LoadFixtureBytes(t, "testdata/TestResActivations/Activations.json"), &getActivationsResponse)
		require.NoError(t, err)

		createActivationsResponse := appsec.CreateActivationsResponse{}
		err = json.Unmarshal(testutils.LoadFixtureBytes(t, "testdata/TestResActivations/Activations.json"), &createActivationsResponse)
		require.NoError(t, err)

		client.On("GetActivations",
			mock.Anything,
			appsec.GetActivationsRequest{ActivationID: 547694},
		).Return(&getActivationsResponse, nil)

		client.On("CreateActivations",
			mock.Anything,
			appsec.CreateActivationsRequest{
				Action:             "ACTIVATE",
				Network:            "STAGING",
				Note:               "",
				NotificationEmails: []string{"user@example.com"},
				ActivationConfigs: []struct {
					ConfigID      int `json:"configId"`
					ConfigVersion int `json:"configVersion"`
				}{{ConfigID: 43253, ConfigVersion: 7}}},
		).Return(&createActivationsResponse, nil)

		client.On("RemoveActivations",
			mock.Anything,
			appsec.RemoveActivationsRequest{
				ActivationID:       547694,
				Action:             "DEACTIVATE",
				Network:            "STAGING",
				Note:               "",
				NotificationEmails: []string{"user@example.com"},
				ActivationConfigs: []struct {
					ConfigID      int `json:"configId"`
					ConfigVersion int `json:"configVersion"`
				}{{ConfigID: 43253, ConfigVersion: 7}}},
		).Return(&removeActivationsResponse, nil)

		useClient(client, func() {
			resource.Test(t, resource.TestCase{
				IsUnitTest:        false,
				ProviderFactories: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: testutils.LoadFixtureString(t, "testdata/TestResActivations/match_by_id.tf"),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttr("akamai_appsec_activations.test", "id", "547694"),
						),
					},
				},
			})
		})

		client.AssertExpectations(t)
	})

}
