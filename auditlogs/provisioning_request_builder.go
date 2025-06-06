// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package auditlogs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i817f881c915e69dd137be7bad331c8136ce800e35c919d4f8934998c1448647b "github.com/sysucjf111/msgraph-sdk-go/models"
    idae218f47fa3b03633eafcaa0318872d0b9b7f3b9f3a816347f92aeb0932e506 "github.com/sysucjf111/msgraph-sdk-go/models/odataerrors"
)

// ProvisioningRequestBuilder provides operations to manage the provisioning property of the microsoft.graph.auditLogRoot entity.
type ProvisioningRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ProvisioningRequestBuilderGetQueryParameters get all provisioning events that occurred in your tenant, such as the deletion of a group in a target application or the creation of a user when provisioning user accounts from your HR system. 
type ProvisioningRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ProvisioningRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProvisioningRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ProvisioningRequestBuilderGetQueryParameters
}
// ByProvisioningObjectSummaryId provides operations to manage the provisioning property of the microsoft.graph.auditLogRoot entity.
// returns a *ProvisioningProvisioningObjectSummaryItemRequestBuilder when successful
func (m *ProvisioningRequestBuilder) ByProvisioningObjectSummaryId(provisioningObjectSummaryId string)(*ProvisioningProvisioningObjectSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if provisioningObjectSummaryId != "" {
        urlTplParams["provisioningObjectSummary%2Did"] = provisioningObjectSummaryId
    }
    return NewProvisioningProvisioningObjectSummaryItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewProvisioningRequestBuilderInternal instantiates a new ProvisioningRequestBuilder and sets the default values.
func NewProvisioningRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProvisioningRequestBuilder) {
    m := &ProvisioningRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/auditLogs/provisioning{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewProvisioningRequestBuilder instantiates a new ProvisioningRequestBuilder and sets the default values.
func NewProvisioningRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ProvisioningRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewProvisioningRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ProvisioningCountRequestBuilder when successful
func (m *ProvisioningRequestBuilder) Count()(*ProvisioningCountRequestBuilder) {
    return NewProvisioningCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get all provisioning events that occurred in your tenant, such as the deletion of a group in a target application or the creation of a user when provisioning user accounts from your HR system. 
// returns a ProvisioningObjectSummaryCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/provisioningobjectsummary-list?view=graph-rest-1.0
func (m *ProvisioningRequestBuilder) Get(ctx context.Context, requestConfiguration *ProvisioningRequestBuilderGetRequestConfiguration)(i817f881c915e69dd137be7bad331c8136ce800e35c919d4f8934998c1448647b.ProvisioningObjectSummaryCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": idae218f47fa3b03633eafcaa0318872d0b9b7f3b9f3a816347f92aeb0932e506.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i817f881c915e69dd137be7bad331c8136ce800e35c919d4f8934998c1448647b.CreateProvisioningObjectSummaryCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i817f881c915e69dd137be7bad331c8136ce800e35c919d4f8934998c1448647b.ProvisioningObjectSummaryCollectionResponseable), nil
}
// ToGetRequestInformation get all provisioning events that occurred in your tenant, such as the deletion of a group in a target application or the creation of a user when provisioning user accounts from your HR system. 
// returns a *RequestInformation when successful
func (m *ProvisioningRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ProvisioningRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ProvisioningRequestBuilder when successful
func (m *ProvisioningRequestBuilder) WithUrl(rawUrl string)(*ProvisioningRequestBuilder) {
    return NewProvisioningRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
