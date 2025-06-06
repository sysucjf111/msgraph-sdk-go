// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package auditlogs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AuditLogsRequestBuilder builds and executes requests for operations under \auditLogs
type AuditLogsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewAuditLogsRequestBuilderInternal instantiates a new AuditLogsRequestBuilder and sets the default values.
func NewAuditLogsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuditLogsRequestBuilder) {
    m := &AuditLogsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/auditLogs", pathParameters),
    }
    return m
}
// NewAuditLogsRequestBuilder instantiates a new AuditLogsRequestBuilder and sets the default values.
func NewAuditLogsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuditLogsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuditLogsRequestBuilderInternal(urlParams, requestAdapter)
}
// DirectoryAudits provides operations to manage the directoryAudits property of the microsoft.graph.auditLogRoot entity.
// returns a *DirectoryAuditsRequestBuilder when successful
func (m *AuditLogsRequestBuilder) DirectoryAudits()(*DirectoryAuditsRequestBuilder) {
    return NewDirectoryAuditsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Provisioning provides operations to manage the provisioning property of the microsoft.graph.auditLogRoot entity.
// returns a *ProvisioningRequestBuilder when successful
func (m *AuditLogsRequestBuilder) Provisioning()(*ProvisioningRequestBuilder) {
    return NewProvisioningRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SignIns provides operations to manage the signIns property of the microsoft.graph.auditLogRoot entity.
// returns a *SignInsRequestBuilder when successful
func (m *AuditLogsRequestBuilder) SignIns()(*SignInsRequestBuilder) {
    return NewSignInsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
