# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

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
