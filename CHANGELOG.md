# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.4.2, 2026-10-27

- The development image is based on Debian (`golang:1.26.10-trixie`) instead of Alpine. Alpine keeps only the latest build of each package, so the exact pins broke the image build whenever a package was updated. System packages now come from the snapshot.debian.org date the base image was built from, which keeps them fixed for a given base image digest. jq 1.8.2, shellcheck 0.11.0 and shfmt 3.13.1 keep their versions, and Semgrep runs in its own virtual environment.

## v0.4.1, 2026-10-27

- Documentation only. The doc comment on `CreateDataConnectorRequest.Headers` and `UpdateDataConnectorRequest.Headers` now warns against putting credentials in header values, because they appear unmasked in execution logs except `Authorization`, and points to `token_ids` instead.

## v0.4.0, 2026-10-24

- Breaking: `DeleteCustomObjectInstancesByExternalId` and `DeleteCustomObjectInstancesById` change arguments on `Client`, `ClientWithResponses` and both interfaces. `DeleteCustomObjectInstancesByExternalId` drops one string argument, `DeleteCustomObjectInstancesById` gains one, and `NewDeleteCustomObjectInstancesByExternalIdRequest` and `NewDeleteCustomObjectInstancesByIdRequest` change the same way. `DeleteCustomObjectInstancesByIdParams.ExternalId` moves to `DeleteCustomObjectInstancesByExternalIdParams.ExternalId`.
- Breaking: `UpdateNewsItem` and `UpdateNewsItemWithResponse` now take `UpdateNewsItemRequest` as the body. `UpdateNewsItemJSONRequestBody` changes from `NewsItemRequest` to it. Added `UpdateNewsItemRequestState` with `Draft` and `Live` values.
- Breaking: `CustomObjectInstancesPaginatedList` loses `Data`, `Pages`, `TotalCount` and `Type`, and is no longer comparable. `ListCustomObjectInstances200JSONResponseBody` is a new union with `AsCustomObjectInstance`, `AsCustomObjectInstanceNotFound`, `AsCustomObjectInstancesPaginatedList` and matching `From`, `Merge` and JSON methods.
- Breaking: `CreateOrUpdateCustomObjectInstanceRequest.CustomAttributes` is now a map of `CreateOrUpdateCustomObjectInstanceRequest_CustomAttributes_AdditionalProperties` instead of `map[string]string`.
- Breaking: `NewsItem.Labels`, `NewsItem.Reactions` and `NewsItemRequest.Reactions` are now `*[]string` instead of `*[]*string`. `Predicate.Value` is now `interface{}`. `PaginatedResponse.Pages` is now `*PagesLink`. Removed `PaginatedResponseTypeConversationList`.
- Breaking: removed `BulkContentActionsResponse.JSON403` and `GetJSON403`.
- Added `NewsItemRequest.CoverImageUrl`, `Predicate.Predicates`, `InternalArticleListItem.Tags`, `Locale` and `JsonBlocks` on the internal article create and update requests, and `AdminId` on `DetachTagFromContentSnippetParams` and `DetachTagFromInternalArticleParams`.
- Added `Page` and `PerPage` on `ListInternalArticlesParams`, `SearchInternalArticlesParams`, `ListLiveNewsfeedItemsParams`, `ListNewsItemsParams` and `ListNewsfeedsParams`.
- Added `JSON400`, `JSON401`, `JSON404`, `JSON422` and `JSON403` response fields with `GetJSON` methods on many content snippet, internal article, news item, newsfeed, folder, audience, conversation, dynamic content, custom object and search responses.

## v0.3.0, 2026-10-23

- Breaking: `AskFinResponse.JSON200`, `ReplyToFinResponse.JSON200`, `EscalateFinConversationResponse.JSON200` and `ListFinCapabilitiesResponse.JSON200` have new body types, and `SseSubscriptionUrl` is removed from the Ask Fin body. `CreateContactResponse.JSON200` and `MergeContactResponse.JSON200` also have new body types, as do their `GetJSON200` methods.
- Breaking: `AttachSubscriptionTypeToContactJSONBody.ConsentType` is now `*string`. `FinAgentMessage.Timestamp` is now `*time.Time`. `UpdateExternalPageRequest.ExternalId` is now `string` and `Url` is now `*string`. `ListAllCompaniesParams.Order` is now `*ListAllCompaniesParamsOrder`.
- Breaking: `Contact.Avatar` is now `*string`. `Company.Notes` is now a struct with `Data`, `HasMore` and a total count. `CompanyList.Pages` is now `*PagesLink` and `CompanyScroll.Pages` is now `*map[string]interface{}`.
- Breaking: `CreateDataConnectorRequest` and `UpdateDataConnectorRequest` `DataInputs` and `Headers` have new item types, and header `Name` and `Value` are now plain strings. `RunFinProcedureJSONBody.Message` is now an inline struct with a `RunFinProcedureJSONBodyMessageAuthor` author.
- Breaking: removed `CompanyNote.Author`, `Company` and `CreatedAt`, `ContactAttachedCompanies.Companies` (use `Data`), `UpdateExternalPageRequest.FinAvailability`, `ContactsBulkJobStatePending`, and `ExportWorkflowResponse.JSON403` with `GetJSON403`.
- Breaking: removed `CreateExternalPageRequestLocale`, `ExternalPageLocale`, `UpdateExternalPageRequestLocale` and their `En` constants and `Valid` methods.
- Added `ResponseFormat` with `Html` and `Markdown` values on the Fin start, reply and ask request bodies.
- Added `CompanyUserList`, `CompanySearchResponse`, `ContactSubscriptionTypes`, and the `AsCompanyUserList`, `FromCompanyUserList` and `MergeCompanyUserList` union methods on the company retrieve body. `RetrieveCompanyParams` and `RetrieveACompanyByIdParams` gain `Type`, and the latter also `Page` and `PerPage`.
- Added contact fields `SmsConsent`, `UnsubscribedFromSms`, `LanguageOverride`, `Referrer`, `Utm*` fields, `OptedInSubscriptionTypes`, `OptedOutSubscriptionTypes`, `EnabledPushMessaging`, and `ContinentCode` and `CountryCode` on `ContactLocation`.
- Added `PerPage`, `StartingAfter`, `Page`, `Order` and `EscapeIdentity` params on the contact, company, external page and content import source list calls, and `IncludeMergeHistory` on `ListAttachedContactsParams`.
- Added `EnforceAuth` and `InternalDescription` on data connector types, `TransformCodeBlock` on `DataConnectorDetail`, new data connector error type and source constants, `ExternalPage.Tags`, `AiAgentAvailability`, `AiCopilotAvailability` and `AiSalesAgentAvailability` on external page requests, `CreateOrUpdateCompanyRequest.Id`, `UpdateCompanyRequest.UpdateLastRequestAt` and `CompanyNote.Url`.
- Added `JSON401`, `JSON403`, `JSON404` and `JSON422` response fields on several contact, content import source, external page and Fin responses.

## v0.2.0, 2026-10-22

- Added `GetSubscriptionType` on `Client` and `ClientWithResponses`, with `GetSubscriptionTypeParams`, `GetSubscriptionTypeResponse`, `NewGetSubscriptionTypeRequest` and `ParseGetSubscriptionTypeResponse`.
- Added `DataEvent`, `DataEventList` and `LisDataEvents200JSONResponseBody`, a union with `AsDataEventList`, `AsDataEventSummary` and matching `From`, `Merge` and JSON methods.
- Added `CallTranscriptTurn`, and `ShowCallTranscriptResponse` now has `JSON200`. `Headers200` and `ShowCallTranscriptResponse200Headers` are removed.
- Added article fields `AudienceIds`, `ContentId`, `Locale` and `ParentIds`, `Page` and `PerPage` on the list and search params, `SearchArticlesParamsState`, and `JSON400` and `JSON422` response fields on several article and collection responses.
- Breaking: `Pages` on article, collection, internal article and help center redirect lists is now `*PagesLink` instead of `*CursorPages`. `CursorPages.Next` is now an inline struct instead of `*StartingAfterPaging`.
- Breaking: `CustomAttributes` is now a struct with `AITopicAndSubtopic`, `CXScoreReasons` and `AdditionalProperties`. `AiTopicSubtopicList`, `CxScoreReasonList` and their `As`, `From` and `Merge` methods are removed. Added `AiTopicSubtopic` and `CxScoreReason`.
- Breaking: `CustomAttributes` on company requests is now a map of union types instead of `map[string]string`. This affects `Company`, `CreateOrUpdateCompanyRequest` and `UpdateCompanyRequest`.
- Breaking: `CreateArticleRequest.AuthorId` and `Title` and `CreateCollectionRequest.Name` are now pointers. `CreateArticleRequest.ParentType` and `UpdateArticleRequest.ParentType` are removed.
- Breaking: `CreateDataEventSummariesRequest.EventSummaries` is now a slice of structs, and the type is no longer comparable. `DataEventSummary` loses `Email`, `Events`, `IntercomUserId`, `Type` and `UserId`.
- Breaking: `GetCustomObjectInstancesByIdResponse.JSON200` and `RetrieveCompanyResponse.JSON200` have new body types.
- Breaking: removed `ArticleStatistics.Conversions` (use `Conversations`), `AwayStatusReason.CreatedAt` and `UpdatedAt`, `ConversationPartMetadata.QuickReplyUuid`, `ConversationTeammates.Teammates` (use `Admins`), `PhoneSwitch.Phone`, `Visitor.LasRequestAt` (use `LastRequestAt`) and `TicketTypeAttributeList.TicketTypeAttributes`.
- Breaking: `InternalArticleListItem.Id` is now `*int` and its `Type` is removed. `CreateTicketTypeRequest.Category` is no longer a pointer. `Error.Errors` items have `Data` instead of `Field`.
- Breaking: renamed `SlaAppliedSlaStatusCancelled` to `SlaAppliedSlaStatusCanceled` and `UpdateContentImportSourceRequestSyncBehaviorAutomated` to `UpdateContentImportSourceRequestSyncBehaviorAutomatic`.
- Raised the Go requirement in `go.mod` from 1.26.9 to 1.26.10 to pick up the `crypto/x509` fix for GO-2026-6809. The development image and the setup reference match.

## v0.1.1, 2026-10-17

- Documentation only. The agent skill explains how to look up every endpoint, parameter and model with `go doc` or on pkg.go.dev, and declares its license.
- The skill no longer lists the API credentials under OpenClaw's `requires.env`, which kept OpenClaw from loading it until they were set. It requires only `go`; the credentials are still documented as what your program reads.
- The README points to the full API reference, and the examples in the README and the skill were checked against the current code.

## v0.1.0, 2026-10-17

- Added `intercom`, a REST client generated with oapi-codegen from Intercom's OpenAPI document (Intercom API, spec version 2.17, OpenAPI 3.0.1). It covers all 258 operations across 182 paths, including admins, contacts, companies, conversations, messages, articles, tickets and exports.
- Number fields are `float64` instead of oapi-codegen's default `float32`.
- The spec patch adds missing reporting job path parameters, resolves two inline enum Go name collisions, collapses single-branch contact body unions, and names four other inline union schemas so typed requests retain their JSON methods and serialize their fields.
- `make generate` selects the newest numbered stable spec from Intercom's public repository and regenerates the client. Initial source: https://raw.githubusercontent.com/intercom/Intercom-OpenAPI/main/descriptions/2.17/api.intercom.io.yaml. The spec is not committed.
- Added the `go-intercom` agent skill and Claude Code and Codex plugin manifests.
