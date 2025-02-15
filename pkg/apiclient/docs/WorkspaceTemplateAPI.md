# \WorkspaceTemplateAPI

All URIs are relative to *http://localhost:3986*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWorkspaceTemplate**](WorkspaceTemplateAPI.md#DeleteWorkspaceTemplate) | **Delete** /workspace-template/{templateName} | Delete workspace template data
[**FindWorkspaceTemplate**](WorkspaceTemplateAPI.md#FindWorkspaceTemplate) | **Get** /workspace-template/{templateName} | Find a workspace template
[**GetDefaultWorkspaceTemplate**](WorkspaceTemplateAPI.md#GetDefaultWorkspaceTemplate) | **Get** /workspace-template/default/{gitUrl} | Get default workspace templates by git url
[**ListWorkspaceTemplates**](WorkspaceTemplateAPI.md#ListWorkspaceTemplates) | **Get** /workspace-template | List workspace templates
[**SaveWorkspaceTemplate**](WorkspaceTemplateAPI.md#SaveWorkspaceTemplate) | **Put** /workspace-template | Set workspace template data
[**SetDefaultWorkspaceTemplate**](WorkspaceTemplateAPI.md#SetDefaultWorkspaceTemplate) | **Patch** /workspace-template/{templateName}/set-default | Set workspace template to default



## DeleteWorkspaceTemplate

> DeleteWorkspaceTemplate(ctx, templateName).Force(force).Execute()

Delete workspace template data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/apiclient"
)

func main() {
	templateName := "templateName_example" // string | Template name
	force := true // bool | Force (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkspaceTemplateAPI.DeleteWorkspaceTemplate(context.Background(), templateName).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceTemplateAPI.DeleteWorkspaceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateName** | **string** | Template name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkspaceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | Force | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindWorkspaceTemplate

> WorkspaceTemplate FindWorkspaceTemplate(ctx, templateName).Execute()

Find a workspace template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/apiclient"
)

func main() {
	templateName := "templateName_example" // string | Template name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceTemplateAPI.FindWorkspaceTemplate(context.Background(), templateName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceTemplateAPI.FindWorkspaceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindWorkspaceTemplate`: WorkspaceTemplate
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceTemplateAPI.FindWorkspaceTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateName** | **string** | Template name | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindWorkspaceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkspaceTemplate**](WorkspaceTemplate.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultWorkspaceTemplate

> WorkspaceTemplate GetDefaultWorkspaceTemplate(ctx, gitUrl).Execute()

Get default workspace templates by git url



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/apiclient"
)

func main() {
	gitUrl := "gitUrl_example" // string | Git URL

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceTemplateAPI.GetDefaultWorkspaceTemplate(context.Background(), gitUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceTemplateAPI.GetDefaultWorkspaceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultWorkspaceTemplate`: WorkspaceTemplate
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceTemplateAPI.GetDefaultWorkspaceTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gitUrl** | **string** | Git URL | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultWorkspaceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkspaceTemplate**](WorkspaceTemplate.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkspaceTemplates

> []WorkspaceTemplate ListWorkspaceTemplates(ctx).Execute()

List workspace templates



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/apiclient"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceTemplateAPI.ListWorkspaceTemplates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceTemplateAPI.ListWorkspaceTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkspaceTemplates`: []WorkspaceTemplate
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceTemplateAPI.ListWorkspaceTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspaceTemplatesRequest struct via the builder pattern


### Return type

[**[]WorkspaceTemplate**](WorkspaceTemplate.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveWorkspaceTemplate

> SaveWorkspaceTemplate(ctx).WorkspaceTemplate(workspaceTemplate).Execute()

Set workspace template data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/apiclient"
)

func main() {
	workspaceTemplate := *openapiclient.NewCreateWorkspaceTemplateDTO(map[string]string{"key": "Inner_example"}, "Name_example", "RepositoryUrl_example") // CreateWorkspaceTemplateDTO | Workspace template

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkspaceTemplateAPI.SaveWorkspaceTemplate(context.Background()).WorkspaceTemplate(workspaceTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceTemplateAPI.SaveWorkspaceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveWorkspaceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceTemplate** | [**CreateWorkspaceTemplateDTO**](CreateWorkspaceTemplateDTO.md) | Workspace template | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDefaultWorkspaceTemplate

> SetDefaultWorkspaceTemplate(ctx, templateName).Execute()

Set workspace template to default



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/apiclient"
)

func main() {
	templateName := "templateName_example" // string | Template name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkspaceTemplateAPI.SetDefaultWorkspaceTemplate(context.Background(), templateName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceTemplateAPI.SetDefaultWorkspaceTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateName** | **string** | Template name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDefaultWorkspaceTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

