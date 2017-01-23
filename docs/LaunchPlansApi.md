# \LaunchPlansApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLaunchPlan**](LaunchPlansApi.md#AddLaunchPlan) | **Post** /launchplan/ | Create Instances Using Launch Plans


# **AddLaunchPlan**
> LaunchPlanResponse AddLaunchPlan($body)

Create Instances Using Launch Plans

A launch plan is a JSON-formatted file that defines the properties of one or more instances. You can use a launch plan to quickly create and start multiple, non-persistent instances in Oracle Compute Cloud Service.<p>A launch plan specifies the provisioning sequence and attributes of the instances that you want to create. Note that while you can reuse your launch plan JSON file to create new instances based on the attributes and provisioning sequence specified in the JSON file, the launch plan itself doesn't persist in Oracle Compute Cloud Service.<p>For information about the attributes you can specify in a launch plan JSON file, see <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=STCSG-GUID-433AEA3E-F569-45BB-8373-6108524EE25E\">Creating Instances Using Launch Plans</a> in <em>Using Oracle Compute Cloud Service (IaaS)</em>.<p><b>Required Role: </b>To complete this task, you must have the <code>Compute_Operations</code> role. If this role isn't assigned to you or you're not sure, then ask your system administrator to ensure that the role is assigned to you in Oracle Cloud My Services. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=MMOCS-GUID-54C2E747-7D5B-451C-A39C-77936178EBB6\">Modifying User Roles</a> in <em>Managing and Monitoring Oracle Cloud</em>.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LaunchPlanPostRequest**](LaunchPlanPostRequest.md)|  | 

### Return type

[**LaunchPlanResponse**](LaunchPlan-response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/oracle-compute-v3+json
 - **Accept**: application/oracle-compute-v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

