// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package identityprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i817f881c915e69dd137be7bad331c8136ce800e35c919d4f8934998c1448647b "github.com/sysucjf111/msgraph-sdk-go/models"
    idae218f47fa3b03633eafcaa0318872d0b9b7f3b9f3a816347f92aeb0932e506 "github.com/sysucjf111/msgraph-sdk-go/models/odataerrors"
)

// RiskyServicePrincipalsItemHistoryRiskyServicePrincipalHistoryItemItemRequestBuilder provides operations to manage the history property of the microsoft.graph.riskyServicePrincipal entity.
type RiskyServicePrincipalsItemHistoryRiskyServicePrincipalHistoryItemItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RiskyServicePrincipalsItemHistoryRiskyServicePrincipalHistoryItemItemRequestBuilderGetQueryParameters represents the risk history of Microsoft Entra service principals.
type RiskyServicePrincipalsItemHistoryRiskyServicePrincipalHistoryItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RiskyServicePrincipalsItemHistoryRiskyServicePrincipalHistoryItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RiskyServicePrincipalsItemHistoryRiskyServicePrincipalHistoryItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RiskyServicePrincipalsItemHistoryRiskyServicePrincipalHistoryItemItemRequestBuilderGetQueryParameters
}
// NewRiskyServicePrincipalsItemHistoryRiskyServicePrincipalHistoryItemItemRequestBuilderInternal instantiates a new RiskyServicePrincipalsItemHistoryRiskyServicePrincipalHistoryItemItemRequestBuilder and sets the default values.
func NewRiskyServicePrincipalsItemHistoryRiskyServicePrincipalHistoryItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RiskyServicePrincipalsItemHistoryRiskyServicePrincipalHistoryItemItemRequestBuilder) {
    m := &RiskyServicePrincipalsItemHistoryRiskyServicePrincipalHistoryItemItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityProtection/riskyServicePrincipals/{riskyServicePrincipal%2Did}/history/{riskyServicePrincipalHistoryItem%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewRiskyServicePrincipalsItemHistoryRiskyServicePrincipalHistoryItemItemRequestBuilder instantiates a new RiskyServicePrincipalsItemHistoryRiskyServicePrincipalHistoryItemItemRequestBuilder and sets the default values.
func NewRiskyServicePrincipalsItemHistoryRiskyServicePrincipalHistoryItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RiskyServicePrincipalsItemHistoryRiskyServicePrincipalHistoryItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRiskyServicePrincipalsItemHistoryRiskyServicePrincipalHistoryItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get represents the risk history of Microsoft Entra service principals.
// returns a RiskyServicePrincipalHistoryItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RiskyServicePrincipalsItemHistoryRiskyServicePrincipalHistoryItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RiskyServicePrincipalsItemHistoryRiskyServicePrincipalHistoryItemItemRequestBuilderGetRequestConfiguration)(i817f881c915e69dd137be7bad331c8136ce800e35c919d4f8934998c1448647b.RiskyServicePrincipalHistoryItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": idae218f47fa3b03633eafcaa0318872d0b9b7f3b9f3a816347f92aeb0932e506.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i817f881c915e69dd137be7bad331c8136ce800e35c919d4f8934998c1448647b.CreateRiskyServicePrincipalHistoryItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i817f881c915e69dd137be7bad331c8136ce800e35c919d4f8934998c1448647b.RiskyServicePrincipalHistoryItemable), nil
}
// ToGetRequestInformation represents the risk history of Microsoft Entra service principals.
// returns a *RequestInformation when successful
func (m *RiskyServicePrincipalsItemHistoryRiskyServicePrincipalHistoryItemItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RiskyServicePrincipalsItemHistoryRiskyServicePrincipalHistoryItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RiskyServicePrincipalsItemHistoryRiskyServicePrincipalHistoryItemItemRequestBuilder when successful
func (m *RiskyServicePrincipalsItemHistoryRiskyServicePrincipalHistoryItemItemRequestBuilder) WithUrl(rawUrl string)(*RiskyServicePrincipalsItemHistoryRiskyServicePrincipalHistoryItemItemRequestBuilder) {
    return NewRiskyServicePrincipalsItemHistoryRiskyServicePrincipalHistoryItemItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
