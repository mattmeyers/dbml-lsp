package lsp

// TODO: This should be the union type number | string
type ProgressToken interface{}

type DocumentURI string

var EOL = []string{"\n", "\r\n", "\r"}

type Position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}

type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}

type Location struct {
	URI   DocumentURI `json:"uri"`
	Range Range       `json:"range"`
}

type LocationLink struct {
	OriginSelectionRange Range       `json:"originSelectionRange,omitempty"`
	TargetURI            DocumentURI `json:"targetUri"`
	TargetRange          Range       `json:"targetRange"`
	TargetSelectionRange Range       `json:"targetSelectionRange"`
}

type Diagnostic struct {
	Range    Range              `json:"range"`
	Severity DiagnosticSeverity `json:"severity,omitempty"`
	// TODO: Code can be either a string or an int. Make this a concrete type like ID
	Code               interface{}                    `json:"code,omitempty"`
	Source             string                         `json:"source,omitempty"`
	Message            string                         `json:"message"`
	Tags               []DiagnosticTag                `json:"tags,omitempty"`
	RelatedInformation []DiagnosticRelatedInformation `json:"relatedInformation,omitempty"`
}

type DiagnosticSeverity int

const (
	DSError DiagnosticSeverity = iota + 1
	DSWarning
	DSInformation
	DSHint
)

type DiagnosticTag int

const (
	DTUnnecessary DiagnosticTag = iota + 1
	DTDeprecated
)

type DiagnosticRelatedInformation struct {
	Location Location `json:"location"`
	Message  string   `json:"message"`
}

type Command struct {
	Title     string        `json:"title"`
	Command   string        `json:"command"`
	Arguments []interface{} `json:"arguments,omitempty"`
}

type TextEdit struct {
	Range   Range  `json:"range"`
	NewText string `json:"newText"`
}

type TextDocumentEdit struct {
	TextDocument VersionedTextDocumentIdentifier `json:"textDocument"`
	Edits        []TextEdit                      `json:"edits"`
}

type CreateFileOptions struct {
	Overwrite      bool `json:"overwrite,omitempty"`
	IgnoreIfExists bool `json:"ignoreIfExists,omitempty"`
}

type CreateFile struct {
	// This must be set to "create"
	Kind    string              `json:"kind"`
	URI     DocumentURI         `json:"uri"`
	Options []CreateFileOptions `json:"options,omitempty"`
}

type RenameFileOptions struct {
	Overwrite      bool `json:"overwrite,omitempty"`
	IgnoreIfExists bool `json:"ignoreIfExists,omitempty"`
}

type RenameFile struct {
	// This must be set to "rename"
	Kind    string              `json:"kind"`
	OldURI  DocumentURI         `json:"oldUri"`
	NewURI  DocumentURI         `json:"newUri"`
	Options []RenameFileOptions `json:"options,omitempty"`
}

type DeleteFileOptions struct {
	Recursive         bool `json:"recursive,omitempty"`
	IgnoreIfNotExists bool `json:"ignoreIfNotExists,omitempty"`
}

type DeleteFile struct {
	// This must be set to "delete"
	Kind    string              `json:"kind"`
	URI     DocumentURI         `json:"uri"`
	Options []DeleteFileOptions `json:"options,omitempty"`
}

// TODO: Figure out how to define this with interfaces
type WorkspaceEdit struct{}

// TODO: Define this
type WorkspaceEditClientCapabilites struct{}

type TextDocumentIdentifier struct {
	URI DocumentURI `json:"uri"`
}

type TextDocumentItem struct {
	URI        DocumentURI `json:"uri"`
	LanguageID string      `json:"LanguageId"`
	Version    int         `json:"version"`
	Text       string      `json:"text"`
}

type VersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier
	Version *int `json:"version"`
}

type TextDocumentPositionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Position     Position               `json:"position"`
}

type DocumentFilter struct {
	Language string `json:"language,omitempty"`
	Scheme   string `json:"scheme,omitempty"`
	Pattern  string `json:"pattern,omitempty"`
}

type DocumentSelector []DocumentFilter

type StaticRegistrationOptions struct {
	ID string `json:"id,omitempty"`
}

type TextDocumentRegistrationOptions struct {
	DocumentSelector DocumentSelector `json:"documentSelector"`
}

type MarkupKind string

const (
	PlainText MarkupKind = "plaintext"
	Markdown  MarkupKind = "markdown"
)

type MarkupContent struct {
	Kind  MarkupKind `json:"kind"`
	Value string     `json:"value"`
}

type WorkDoneProgressBegin struct {
	// Must be set to "begin"
	Kind        string  `json:"kind"`
	Title       string  `json:"title"`
	Cancellable bool    `json:"cancellable,omitempty"`
	Message     string  `json:"message,omitempty"`
	Percentage  float64 `json:"percentage,omitempty"`
}

type WorkDoneProgressReport struct {
	// Must be set to "report"
	Kind        string  `json:"kind"`
	Cancellable bool    `json:"cancellable,omitempty"`
	Message     string  `json:"message,omitempty"`
	Percentage  float64 `json:"percentage,omitempty"`
}

type WorkDoneProgressEnd struct {
	// Must be set to "end"
	Kind    string `json:"kind"`
	Message string `json:"message,omitempty"`
}

type WorkDoneProgressParams struct {
	WorkDoneToken ProgressToken `json:"workDoneToken,omitempty"`
}

type WorkDoneProgressOptions struct {
	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}

type PartialResultParams struct {
	PartialResultToken ProgressToken `json:"partialResultToken,omitempty"`
}
