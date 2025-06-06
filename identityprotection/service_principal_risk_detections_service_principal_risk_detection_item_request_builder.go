// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package identityprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i817f881c915e69dd137be7bad331c8136ce800e35c919d4f8934998c1448647b "github.com/sysucjf111/msgraph-sdk-go/models"
    idae218f47fa3b03633eafcaa0318872d0b9b7f3b9f3a816347f92aeb0932e506 "github.com/sysucjf111/msgraph-sdk-go/models/odataerrors"
)

// ServicePrincipalRiskDetectionsServicePrincipalRiskDetectionItemRequestBuilder provides operations to manage the servicePrincipalRiskDetections property of the microsoft.graph.identityProtectionRoot entity.
type ServicePrincipalRiskDetectionsServicePrincipalRiskDetectionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServicePrincipalRiskDetectionsServicePrincipalRiskDetectionItemRequestBuilderGetQueryParameters read the properties and relationships of a servicePrincipalRiskDetection object.
type ServicePrincipalRiskDetectionsServicePrincipalRiskDetectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ServicePrincipalRiskDetectionsServicePrincipalRiskDetectionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalRiskDetectionsServicePrincipalRiskDetectionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServicePrincipalRiskDetectionsServicePrincipalRiskDetectionItemRequestBuilderGetQueryParameters
}
// NewServicePrincipalRiskDetectionsServicePrincipalRiskDetectionItemRequestBuilderInternal instantiates a new ServicePrincipalRiskDetectionsServicePrincipalRiskDetectionItemRequestBuilder and sets the default values.
func NewServicePrincipalRiskDetectionsServicePrincipalRiskDetectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalRiskDetectionsServicePrincipalRiskDetectionItemRequestBuilder) {
    m := &ServicePrincipalRiskDetectionsServicePrincipalRiskDetectionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityProtection/servicePrincipalRiskDetections/{servicePrincipalRiskDetection%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewServicePrincipalRiskDetectionsServicePrincipalRiskDetectionItemRequestBuilder instantiates a new ServicePrincipalRiskDetectionsServicePrincipalRiskDetectionItemRequestBuilder and sets the default values.
func NewServicePrincipalRiskDetectionsServicePrincipalRiskDetectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalRiskDetectionsServicePrincipalRiskDetectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServicePrincipalRiskDetectionsServicePrincipalRiskDetectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get read the properties and relationships of a servicePrincipalRiskDetection object.
// returns a ServicePrincipalRiskDetectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceprincipalriskdetection-get?view=graph-rest-1.0
func (m *ServicePrincipalRiskDetectionsServicePrincipalRiskDetectionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ServicePrincipalRiskDetectionsServicePrincipalRiskDetectionItemRequestBuilderGetRequestConfiguration)(i817f881c915e69dd137be7bad331c8136ce800e35c919d4f8934998c1448647b.ServicePrincipalRiskDetectionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": idae218f47fa3b03633eafcaa0318872d0b9b7f3b9f3a816347f92aeb0932e506.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i817f881c915e69dd137be7bad331c8136ce800e35c919d4f8934998c1448647b.CreateServicePrincipalRiskDetectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i817f881c915e69dd137be7bad331c8136ce800e35c919d4f8934998c1448647b.ServicePrincipalRiskDetectionable), nil
}
// ToGetRequestInformation read the properties and relationships of a servicePrincipalRiskDetection object.
// returns a *RequestInformation when successful
func (m *ServicePrincipalRiskDetectionsServicePrincipalRiskDetectionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ServicePrincipalRiskDetectionsServicePrincipalRiskDetectionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ServicePrincipalRiskDetectionsServicePrincipalRiskDetectionItemRequestBuilder when successful
func (m *ServicePrincipalRiskDetectionsServicePrincipalRiskDetectionItemRequestBuilder) WithUrl(rawUrl string)(*ServicePrincipalRiskDetectionsServicePrincipalRiskDetectionItemRequestBuilder) {
    return NewServicePrincipalRiskDetectionsServicePrincipalRiskDetectionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
