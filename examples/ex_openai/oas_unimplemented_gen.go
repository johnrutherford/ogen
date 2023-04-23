// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CancelFineTune implements cancelFineTune operation.
//
// Immediately cancel a fine-tune job.
//
// POST /fine-tunes/{fine_tune_id}/cancel
func (UnimplementedHandler) CancelFineTune(ctx context.Context, params CancelFineTuneParams) (r FineTune, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateAnswer implements createAnswer operation.
//
// Answers the specified question using the provided documents and examples.
// The endpoint first [searches](/docs/api-reference/searches) over provided documents or files to
// find relevant context. The relevant context is combined with the provided examples and question to
// create the prompt for [completion](/docs/api-reference/completions).
//
// Deprecated: schema marks this operation as deprecated.
//
// POST /answers
func (UnimplementedHandler) CreateAnswer(ctx context.Context, req *CreateAnswerRequest) (r *CreateAnswerResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateChatCompletion implements createChatCompletion operation.
//
// Creates a completion for the chat message.
//
// POST /chat/completions
func (UnimplementedHandler) CreateChatCompletion(ctx context.Context, req *CreateChatCompletionRequest) (r *CreateChatCompletionResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateClassification implements createClassification operation.
//
// Classifies the specified `query` using provided examples.
// The endpoint first [searches](/docs/api-reference/searches) over the labeled examples
// to select the ones most relevant for the particular query. Then, the relevant examples
// are combined with the query to construct a prompt to produce the final label via the
// [completions](/docs/api-reference/completions) endpoint.
// Labeled examples can be provided via an uploaded `file`, or explicitly listed in the
// request using the `examples` parameter for quick tests and small scale use cases.
//
// Deprecated: schema marks this operation as deprecated.
//
// POST /classifications
func (UnimplementedHandler) CreateClassification(ctx context.Context, req *CreateClassificationRequest) (r *CreateClassificationResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateCompletion implements createCompletion operation.
//
// Creates a completion for the provided prompt and parameters.
//
// POST /completions
func (UnimplementedHandler) CreateCompletion(ctx context.Context, req *CreateCompletionRequest) (r *CreateCompletionResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateEdit implements createEdit operation.
//
// Creates a new edit for the provided input, instruction, and parameters.
//
// POST /edits
func (UnimplementedHandler) CreateEdit(ctx context.Context, req *CreateEditRequest) (r *CreateEditResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateEmbedding implements createEmbedding operation.
//
// Creates an embedding vector representing the input text.
//
// POST /embeddings
func (UnimplementedHandler) CreateEmbedding(ctx context.Context, req *CreateEmbeddingRequest) (r *CreateEmbeddingResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateFile implements createFile operation.
//
// Upload a file that contains document(s) to be used across various endpoints/features. Currently,
// the size of all the files uploaded by one organization can be up to 1 GB. Please contact us if you
// need to increase the storage limit.
//
// POST /files
func (UnimplementedHandler) CreateFile(ctx context.Context, req *CreateFileRequestMultipart) (r OpenAIFile, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateFineTune implements createFineTune operation.
//
// Creates a job that fine-tunes a specified model from a given dataset.
// Response includes details of the enqueued job including job status and the name of the fine-tuned
// models once complete.
// [Learn more about Fine-tuning](/docs/guides/fine-tuning).
//
// POST /fine-tunes
func (UnimplementedHandler) CreateFineTune(ctx context.Context, req *CreateFineTuneRequest) (r FineTune, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateImage implements createImage operation.
//
// Creates an image given a prompt.
//
// POST /images/generations
func (UnimplementedHandler) CreateImage(ctx context.Context, req *CreateImageRequest) (r ImagesResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateImageEdit implements createImageEdit operation.
//
// Creates an edited or extended image given an original image and a prompt.
//
// POST /images/edits
func (UnimplementedHandler) CreateImageEdit(ctx context.Context, req *CreateImageEditRequestMultipart) (r ImagesResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateImageVariation implements createImageVariation operation.
//
// Creates a variation of a given image.
//
// POST /images/variations
func (UnimplementedHandler) CreateImageVariation(ctx context.Context, req *CreateImageVariationRequestMultipart) (r ImagesResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateModeration implements createModeration operation.
//
// Classifies if text violates OpenAI's Content Policy.
//
// POST /moderations
func (UnimplementedHandler) CreateModeration(ctx context.Context, req *CreateModerationRequest) (r *CreateModerationResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateSearch implements createSearch operation.
//
// The search endpoint computes similarity scores between provided query and documents. Documents can
// be passed directly to the API if there are no more than 200 of them.
// To go beyond the 200 document limit, documents can be processed offline and then used for
// efficient retrieval at query time. When `file` is set, the search endpoint searches over all the
// documents in the given file and returns up to the `max_rerank` number of documents. These
// documents will be returned along with their search scores.
// The similarity score is a positive score that usually ranges from 0 to 300 (but can sometimes go
// higher), where a score above 200 usually means the document is semantically similar to the query.
//
// Deprecated: schema marks this operation as deprecated.
//
// POST /engines/{engine_id}/search
func (UnimplementedHandler) CreateSearch(ctx context.Context, req *CreateSearchRequest, params CreateSearchParams) (r *CreateSearchResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateTranscription implements createTranscription operation.
//
// Transcribes audio into the input language.
//
// POST /audio/transcriptions
func (UnimplementedHandler) CreateTranscription(ctx context.Context, req *CreateTranscriptionRequestMultipart) (r *CreateTranscriptionResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateTranslation implements createTranslation operation.
//
// Translates audio into into English.
//
// POST /audio/translations
func (UnimplementedHandler) CreateTranslation(ctx context.Context, req *CreateTranslationRequestMultipart) (r *CreateTranslationResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteFile implements deleteFile operation.
//
// Delete a file.
//
// DELETE /files/{file_id}
func (UnimplementedHandler) DeleteFile(ctx context.Context, params DeleteFileParams) (r *DeleteFileResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteModel implements deleteModel operation.
//
// Delete a fine-tuned model. You must have the Owner role in your organization.
//
// DELETE /models/{model}
func (UnimplementedHandler) DeleteModel(ctx context.Context, params DeleteModelParams) (r *DeleteModelResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// DownloadFile implements downloadFile operation.
//
// Returns the contents of the specified file.
//
// GET /files/{file_id}/content
func (UnimplementedHandler) DownloadFile(ctx context.Context, params DownloadFileParams) (r string, _ error) {
	return r, ht.ErrNotImplemented
}

// ListEngines implements listEngines operation.
//
// Lists the currently available (non-finetuned) models, and provides basic information about each
// one such as the owner and availability.
//
// Deprecated: schema marks this operation as deprecated.
//
// GET /engines
func (UnimplementedHandler) ListEngines(ctx context.Context) (r *ListEnginesResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// ListFiles implements listFiles operation.
//
// Returns a list of files that belong to the user's organization.
//
// GET /files
func (UnimplementedHandler) ListFiles(ctx context.Context) (r *ListFilesResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// ListFineTuneEvents implements listFineTuneEvents operation.
//
// Get fine-grained status updates for a fine-tune job.
//
// GET /fine-tunes/{fine_tune_id}/events
func (UnimplementedHandler) ListFineTuneEvents(ctx context.Context, params ListFineTuneEventsParams) (r *ListFineTuneEventsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// ListFineTunes implements listFineTunes operation.
//
// List your organization's fine-tuning jobs.
//
// GET /fine-tunes
func (UnimplementedHandler) ListFineTunes(ctx context.Context) (r *ListFineTunesResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// ListModels implements listModels operation.
//
// Lists the currently available models, and provides basic information about each one such as the
// owner and availability.
//
// GET /models
func (UnimplementedHandler) ListModels(ctx context.Context) (r *ListModelsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// RetrieveEngine implements retrieveEngine operation.
//
// Retrieves a model instance, providing basic information about it such as the owner and
// availability.
//
// Deprecated: schema marks this operation as deprecated.
//
// GET /engines/{engine_id}
func (UnimplementedHandler) RetrieveEngine(ctx context.Context, params RetrieveEngineParams) (r Engine, _ error) {
	return r, ht.ErrNotImplemented
}

// RetrieveFile implements retrieveFile operation.
//
// Returns information about a specific file.
//
// GET /files/{file_id}
func (UnimplementedHandler) RetrieveFile(ctx context.Context, params RetrieveFileParams) (r OpenAIFile, _ error) {
	return r, ht.ErrNotImplemented
}

// RetrieveFineTune implements retrieveFineTune operation.
//
// Gets info about the fine-tune job.
// [Learn more about Fine-tuning](/docs/guides/fine-tuning).
//
// GET /fine-tunes/{fine_tune_id}
func (UnimplementedHandler) RetrieveFineTune(ctx context.Context, params RetrieveFineTuneParams) (r FineTune, _ error) {
	return r, ht.ErrNotImplemented
}

// RetrieveModel implements retrieveModel operation.
//
// Retrieves a model instance, providing basic information about the model such as the owner and
// permissioning.
//
// GET /models/{model}
func (UnimplementedHandler) RetrieveModel(ctx context.Context, params RetrieveModelParams) (r Model, _ error) {
	return r, ht.ErrNotImplemented
}
