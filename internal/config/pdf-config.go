package config

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)



var PortraitConfig = &pdfcpu.Import{
    PageSize: "A4",                  // defaults to portrait
	PageDim: &types.Dim{Width: 595, Height: 842},
    Pos:      types.Full,         // scale to fit page
    Scale:    1.0,
}

var LandscapeConfig = &pdfcpu.Import{
    PageDim:  &types.Dim{Width: 842, Height: 595},  // manually swapped dimensions
    UserDim:  true,                  // tells pdfcpu to honor PageDim
    Pos:      types.Full,
    Scale:    1.0,
}

var DefaultConfig = &pdfcpu.Import{
    PageSize: "A4",                  // defaults to portrait
	PageDim: &types.Dim{Width: 595, Height: 842},
    Pos:      types.Full,         // scale to fit page
    Scale:    1.0,
}