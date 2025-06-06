// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package identityprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i817f881c915e69dd137be7bad331c8136ce800e35c919d4f8934998c1448647b "github.com/sysucjf111/msgraph-sdk-go/models"
    idae218f47fa3b03633eafcaa0318872d0b9b7f3b9f3a816347f92aeb0932e506 "github.com/sysucjf111/msgraph-sdk-go/models/odataerrors"
)

// RiskyUsersItemHistoryRequestBuilder provides operations to manage the history property of the microsoft.graph.riskyUser entity.
type RiskyUsersItemHistoryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RiskyUsersItemHistoryRequestBuilderGetQueryParameters get the riskyUserHistoryItems from the history navigation property.
type RiskyUsersItemHistoryRequestBuilderGetQueryParameters struct {
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
// RiskyUsersItemHistoryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RiskyUsersItemHistoryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RiskyUsersItemHistoryRequestBuilderGetQueryParameters
}
// ByRiskyUserHistoryItemId provides operations to manage the history property of the microsoft.graph.riskyUser entity.
// returns a *RiskyUsersItemHistoryRiskyUserHistoryItemItemRequestBuilder when successful
func (m *RiskyUsersItemHistoryRequestBuilder) ByRiskyUserHistoryItemId(riskyUserHistoryItemId string)(*RiskyUsersItemHistoryRiskyUserHistoryItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if riskyUserHistoryItemId != "" {
        urlTplParams["riskyUserHistoryItem%2Did"] = riskyUserHistoryItemId
    }
    return NewRiskyUsersItemHistoryRiskyUserHistoryItemItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewRiskyUsersItemHistoryRequestBuilderInternal instantiates a new RiskyUsersItemHistoryRequestBuilder and sets the default values.
func NewRiskyUsersItemHistoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RiskyUsersItemHistoryRequestBuilder) {
    m := &RiskyUsersItemHistoryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityProtection/riskyUsers/{riskyUser%2Did}/history{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewRiskyUsersItemHistoryRequestBuilder instantiates a new RiskyUsersItemHistoryRequestBuilder and sets the default values.
func NewRiskyUsersItemHistoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RiskyUsersItemHistoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRiskyUsersItemHistoryRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *RiskyUsersItemHistoryCountRequestBuilder when successful
func (m *RiskyUsersItemHistoryRequestBuilder) Count()(*RiskyUsersItemHistoryCountRequestBuilder) {
    return NewRiskyUsersItemHistoryCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the riskyUserHistoryItems from the history navigation property.
// returns a RiskyUserHistoryItemCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/riskyuser-list-history?view=graph-rest-1.0
func (m *RiskyUsersItemHistoryRequestBuilder) Get(ctx context.Context, requestConfiguration *RiskyUsersItemHistoryRequestBuilderGetRequestConfiguration)(i817f881c915e69dd137be7bad331c8136ce800e35c919d4f8934998c1448647b.RiskyUserHistoryItemCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": idae218f47fa3b03633eafcaa0318872d0b9b7f3b9f3a816347f92aeb0932e506.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i817f881c915e69dd137be7bad331c8136ce800e35c919d4f8934998c1448647b.CreateRiskyUserHistoryItemCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i817f881c915e69dd137be7bad331c8136ce800e35c919d4f8934998c1448647b.RiskyUserHistoryItemCollectionResponseable), nil
}
// ToGetRequestInformation get the riskyUserHistoryItems from the history navigation property.
// returns a *RequestInformation when successful
func (m *RiskyUsersItemHistoryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RiskyUsersItemHistoryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RiskyUsersItemHistoryRequestBuilder when successful
func (m *RiskyUsersItemHistoryRequestBuilder) WithUrl(rawUrl string)(*RiskyUsersItemHistoryRequestBuilder) {
    return NewRiskyUsersItemHistoryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
