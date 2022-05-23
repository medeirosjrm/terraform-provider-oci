// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"

	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"

	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

/**
  Dependency variables:
      hostname = var.stack_mon_hostname_resource1
      management_agent_id = var.stack_mon_management_agent_id_resource1
      hostname2 = var.stack_mon_hostname_resource2
      management_agent_id2 = var.stack_mon_management_agent_id_resource2
*/
var (
	MonitoredResourcesAssociateMonitoredResourceConfig = MonitoredResourcesAssociateMonitoredResourceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resources_associate_monitored_resource", "test_monitored_resources_associate_monitored_resource", acctest.Required, acctest.Create, monitoredResourcesAssociateMonitoredResourceRepresentation)

	monitoredResourcesAssociateMonitoredResourceRepresentation = map[string]interface{}{
		"association_type":        acctest.Representation{RepType: acctest.Required, Create: `contains`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"destination_resource_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_stack_monitoring_monitored_resource.test_destination_resource.id}`},
		"source_resource_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_stack_monitoring_monitored_resource.test_source_resource.id}`},
	}

	MonitoredResourcesAssociateMonitoredResourceResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource", "test_destination_resource", acctest.Optional, acctest.Create, monitoredResourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resource", "test_source_resource", acctest.Optional, acctest.Create, monitoredResourceRepresentation2)
)

// issue-routing-tag: stack_monitoring/default
func TestStackMonitoringMonitoredResourcesAssociateMonitoredResourceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStackMonitoringMonitoredResourcesAssociateMonitoredResourceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managementAgentId1 := utils.GetEnvSettingWithBlankDefault("stack_mon_management_agent_id_resource1")
	if managementAgentId1 == "" {
		t.Skip("Setting environmental variable stack_mon_management_agent_id_resource1 that represents management agent with resource monitoring plugin is pre-requisite for this test")
	}
	managementAgentId1VariableStr := fmt.Sprintf("variable \"stack_mon_management_agent_id_resource1\" { default = \"%s\" }\n", managementAgentId1)

	hostname1 := utils.GetEnvSettingWithBlankDefault("stack_mon_hostname_resource1")
	if hostname1 == "" {
		t.Skip("Setting environmental variable stack_mon_hostname_resource1 that host accessible by agent defined by stack_mon_management_agent_id_resource1 variable is pre-requisite for this test")
	}
	hostname1VariableStr := fmt.Sprintf("variable \"stack_mon_hostname_resource1\" { default = \"%s\" }\n", hostname1)

	managementAgentId2 := utils.GetEnvSettingWithBlankDefault("stack_mon_management_agent_id_resource2")
	if managementAgentId2 == "" {
		t.Skip("Setting environmental variable stack_mon_management_agent_id_resource2 that represents management agent with resource monitoring plugin is pre-requisite for this test")
	}
	managementAgentId2VariableStr := fmt.Sprintf("variable \"stack_mon_management_agent_id_resource2\" { default = \"%s\" }\n", managementAgentId2)

	hostname2 := utils.GetEnvSettingWithBlankDefault("stack_mon_hostname_resource2")
	if hostname2 == "" {
		t.Skip("Setting environmental variable stack_mon_hostname_resource2 that host accessible by agent defined by stack_mon_management_agent_id_resource2 variable is pre-requisite for this test")
	}
	hostname2VariableStr := fmt.Sprintf("variable \"stack_mon_hostname_resource2\" { default = \"%s\" }\n", hostname2)

	resourceName := "oci_stack_monitoring_monitored_resources_associate_monitored_resource.test_monitored_resources_associate_monitored_resource"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+MonitoredResourcesAssociateMonitoredResourceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resources_associate_monitored_resource", "test_monitored_resources_associate_monitored_resource", acctest.Required, acctest.Create, monitoredResourcesAssociateMonitoredResourceRepresentation), "stackmonitoring", "monitoredResourcesAssociateMonitoredResource", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + managementAgentId1VariableStr + hostname1VariableStr + managementAgentId2VariableStr + hostname2VariableStr + MonitoredResourcesAssociateMonitoredResourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_monitored_resources_associate_monitored_resource", "test_monitored_resources_associate_monitored_resource", acctest.Required, acctest.Create, monitoredResourcesAssociateMonitoredResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "association_type", "contains"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "destination_resource_id"),
				resource.TestCheckResourceAttrSet(resourceName, "source_resource_id"),
				resource.TestCheckResourceAttrSet(resourceName, "tenant_id"),
				resource.TestCheckResourceAttr(resourceName, "source_resource_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source_resource_details.0.name", "terraformSecondaryResource"),
				resource.TestCheckResourceAttr(resourceName, "source_resource_details.0.type", "host"),
				resource.TestCheckResourceAttr(resourceName, "destination_resource_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "destination_resource_details.0.name", "terraformResource"),
				resource.TestCheckResourceAttr(resourceName, "destination_resource_details.0.type", "host"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
	})
}