package handlers

import (
	"context"

	lsctx "github.com/hashicorp/terraform-ls/internal/context"
	ilsp "github.com/hashicorp/terraform-ls/internal/lsp"
	lsp "github.com/hashicorp/terraform-ls/internal/protocol"
)

func (h *logHandler) TextDocumentLink(ctx context.Context, params lsp.DocumentLinkParams) ([]lsp.DocumentLink, error) {
	fs, err := lsctx.DocumentStorage(ctx)
	if err != nil {
		return nil, err
	}

	// cc, err := lsctx.ClientCapabilities(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	mf, err := lsctx.ModuleFinder(ctx)
	if err != nil {
		return nil, err
	}

	file, err := fs.GetDocument(ilsp.FileHandlerFromDocumentURI(params.TextDocument.URI))
	if err != nil {
		return nil, err
	}

	mod, err := mf.ModuleByPath(file.Dir())
	if err != nil {
		return nil, err
	}

	schema, err := mf.SchemaForPath(file.Dir())
	if err != nil {
		return nil, err
	}

	d, err := mod.DecoderWithSchema(schema)
	if err != nil {
		return nil, err
	}

	links, err := d.LinksInFile(file.Filename())
	if err != nil {
		return nil, err
	}

	return ilsp.Links(links), nil
}
