package lsp

import (
	"github.com/hashicorp/hcl-lang/lang"
	lsp "github.com/hashicorp/terraform-ls/internal/protocol"
)

func Links(links []lang.Link) []lsp.DocumentLink {
	docLinks := make([]lsp.DocumentLink, len(links))

	for i, link := range links {
		docLinks[i] = lsp.DocumentLink{
			Range:  HCLRangeToLSP(link.Range),
			Target: link.URI,
		}
	}

	return docLinks
}
