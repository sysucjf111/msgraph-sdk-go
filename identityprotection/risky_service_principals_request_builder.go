// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package identityprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i817f881c915e69dd137be7bad331c8136ce800e35c919d4f8934998c1448647b "github.com/sysucjf111/msgraph-sdk-go/models"
    idae218f47fa3b03633eafcaa0318872d0b9b7f3b9f3a816347f92aeb0932e506 "github.com/sysucjf111/msgraph-sdk-go/models/odataerrors"
)

// RiskyServicePrincipalsRequestBuilder provides operations to manage the riskyServicePrincipals property of the microsoft.graph.identityProtectionRoot entity.
type RiskyServicePrincipalsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RiskyServicePrincipalsRequestBuilderGetQueryParameters retrieve the properties and relationships of riskyServicePrincipal objects.
type RiskyServicePrincipalsRequestBuilderGetQueryParameters struct {
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
// RiskyServicePrincipalsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RiskyServicePrincipalsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RiskyServicePrincipalsRequestBuilderGetQueryParameters
}
// ByRiskyServicePrincipalId provides operations to manage the riskyServicePrincipals property of the microsoft.graph.identityProtectionRoot entity.
// returns a *RiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilder when successful
func (m *RiskyServicePrincipalsRequestBuilder) ByRiskyServicePrincipalId(riskyServicePrincipalId string)(*RiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if riskyServicePrincipalId != "" {
        urlTplParams["riskyServicePrincipal%2Did"] = riskyServicePrincipalId
    }
    return NewRiskyServicePrincipalsRiskyServicePrincipalItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewRiskyServicePrincipalsRequestBuilderInternal instantiates a new RiskyServicePrincipalsRequestBuilder and sets the default values.
func NewRiskyServicePrincipalsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RiskyServicePrincipalsRequestBuilder) {
    m := &RiskyServicePrincipalsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityProtection/riskyServicePrincipals{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewRiskyServicePrincipalsRequestBuilder instantiates a new RiskyServicePrincipalsRequestBuilder and sets the default values.
func NewRiskyServicePrincipalsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RiskyServicePrincipalsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRiskyServicePrincipalsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *RiskyServicePrincipalsCountRequestBuilder when successful
func (m *RiskyServicePrincipalsRequestBuilder) Count()(*RiskyServicePrincipalsCountRequestBuilder) {
    return NewRiskyServicePrincipalsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve the properties and relationships of riskyServicePrincipal objects.
// returns a RiskyServicePrincipalCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identityprotectionroot-list-riskyserviceprincipals?view=graph-rest-1.0
func (m *RiskyServicePrincipalsRequestBuilder) Get(ctx context.Context, requestConfiguration *RiskyServicePrincipalsRequestBuilderGetRequestConfiguration)(i817f881c915e69dd137be7bad331c8136ce800e35c919d4f8934998c1448647b.RiskyServicePrincipalCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": idae218f47fa3b03633eafcaa0318872d0b9b7f3b9f3a816347f92aeb0932e506.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i817f881c915e69dd137be7bad331c8136ce800e35c919d4f8934998c1448647b.CreateRiskyServicePrincipalCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i817f881c915e69dd137be7bad331c8136ce800e35c919d4f8934998c1448647b.RiskyServicePrincipalCollectionResponseable), nil
}
// ToGetRequestInformation retrieve the properties and relationships of riskyServicePrincipal objects.
// returns a *RequestInformation when successful
func (m *RiskyServicePrincipalsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RiskyServicePrincipalsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RiskyServicePrincipalsRequestBuilder when successful
func (m *RiskyServicePrincipalsRequestBuilder) WithUrl(rawUrl string)(*RiskyServicePrincipalsRequestBuilder) {
    return NewRiskyServicePrincipalsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
