package common

// All tests are be linked[1] to the integreatly-test-cases[2] repo by using the same ID
// 1. https://gitlab.cee.redhat.com/integreatly-qe/integreatly-test-cases#how-to-automate-a-test-case-and-link-it-back
// 2. https://gitlab.cee.redhat.com/integreatly-qe/integreatly-test-cases
var (
	ALL_TESTS = []TestCase{
		// Add all tests that can be executed prior to a completed installation here
		{"Verify RHMI CRD Exists", TestIntegreatlyCRDExists},
		{"Verify RHMI Config CRD Exists", TestRHMIConfigCRDExists},
	}

	HAPPY_PATH_TESTS = []TestCase{
		// Add all happy path tests to be executed after RHMI installation is completed here
		{"A01 - Verify that all stages in the integreatly-operator CR report completed", TestIntegreatlyStagesStatus}, // Keep test as first on the list, as it ensures that all products are reported as complete
		{"Test RHMI installation CR metric", TestRHMICRMetrics},
		{"A03 - Verify all namespaces have been created with the correct name", TestNamespaceCreated},
		{"A18 - Verify RHMI Config CRs Successful", TestRHMIConfigCRs},
		{"A22 - Verify RHMI Config Updates CRO Strategy Override Config Map", TestRHMIConfigCROStrategyOverride},
		{"B03 - Verify RHMI Developer User Permissions are Correct", TestRHMIDeveloperUserPermissions},
		{"F02 - Verify PodDisruptionBudgets exist", TestIntegreatlyPodDisruptionBudgetsExist},
		{"E01 - Verify Grafana Route is accessible", TestGrafanaExternalRouteAccessible},
		{"E05 - Verify Grafana Route returns dashboards", TestGrafanaExternalRouteDashboardExist},
		{"B06 - Verify users with no email get default email", TestDefaultUserEmail},
		{"Verify Network Policy allows cross NS access to SVC", TestNetworkPolicyAccessNSToSVC},
		{"F08 - Verify Replicas Scale correctly in RHSSO and user SSO", TestReplicasInRHSSOAndUserSSO},
	}

	MANAGED_PRODUCT_TESTS = []TestCase{
		{"A05 - Verify product operator version", TestProductOperatorVersions},
		{"A06 - Verify PVC", TestPVClaims},
		{"A07 - Verify product versions", TestProductVersions},
		{"A15 - Verify Stateful Set resources have the expected replicas", TestStatefulSetsExpectedReplicas},
		{"A09 - Verify Subscription Install Plan Strategy", TestSubscriptionInstallPlanType},
		{"A08 - Verify all products routes are created", TestIntegreatlyRoutesExist},
		{"A14 - Verify Deployment Config resources have the expected replicas", TestDeploymentConfigExpectedReplicas},
		{"A13 - Verify Deployment resources have the expected replicas", TestDeploymentExpectedReplicas},
		{"A10 - Verify CRO Postgres CRs Successful", TestCROPostgresSuccessfulState},
		{"A11 - Verify CRO Redis CRs Successful", TestCRORedisSuccessfulState},
		{"A12 - Verify CRO BlobStorage CRs Successful", TestCROBlobStorageSuccessfulState},
		{"A16 - Custom first broker login flow", TestAuthDelayFirstBrokerLogin},
		{"Verify servicemonitors are cloned in monitoring namespace and rolebindings are created", TestServiceMonitorsCloneAndRolebindingsExist},
		{"Verify Alerts are not firing during or after installation apart from DeadMansSwitch", TestIntegreatlyAlertsFiring},
		{"B04 - Verify Dedicated Admin User Permissions are Correct", TestDedicatedAdminUserPermissions},
		{"B05 - Verify Codeready CRUDL permissions", TestCodereadyCrudlPermisssions},
		{"C01 - Verify Alerts are not pending or firing apart from DeadMansSwitch", TestIntegreatlyAlertsPendingOrFiring},
		{"C03 - Verify that alerting mechanism works", TestIntegreatlyAlertsMechanism},
		{"C04 - Verify Alerts exist", TestIntegreatlyAlertsExist},
		{"E02 - Verify that all dashboards are installed and all the graphs are filled with data", TestDashboardsData},
		{"E03 - Verify dashboards exist", TestIntegreatlyDashboardsExist},
		{"F05 - Verify Replicas Scale correctly in Threescale", TestReplicasInThreescale},
		{"F06 - Verify Replicas Scale correctly in Apicurito", TestReplicasInApicurito},
		{"H03 - Verify 3scale CRUDL permissions", Test3ScaleCrudlPermissions},
		{"H05 - Verify Fuse CRUDL permissions", TestFuseCrudlPermissions},
	}

	SELF_MANAGED_PRODUCT_TESTS = []TestCase{
		{"Test ApicurioRegistry API", TestApicurioRegistryAPI},
	}

	DESTRUCTIVE_TESTS = []TestCase{
		// Add all destructive tests here that should not be executed as part of the happy path tests
		{"J03 - Verify namespaces restored when deleted", TestNamespaceRestoration},
	}
)
