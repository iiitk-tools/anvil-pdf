package pdffunc

import (
	"bytes"
	"io"

	"github.com/iiitk-tools/anvil-pdf/internal/config"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func ConvertImagesToPDF(imgs []io.Reader, orientation string) (*bytes.Buffer, error) {
	var importConfig *pdfcpu.Import
	pdfBuff := new(bytes.Buffer)

	switch orientation {
	case "portrait":
		importConfig = config.PortraitConfig
	case "landscape":
		importConfig = config.LandscapeConfig
	default:
		importConfig = config.DefaultConfig
	}

	// Use pdfcpu.NewDefaultConfiguration()
	err := api.ImportImages(nil, pdfBuff, imgs, importConfig, model.NewDefaultConfiguration())
	if err != nil {
		return nil, err
	}

	return pdfBuff, nil
}
