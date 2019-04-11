package render

import (
	"github.com/VirtusLab/go-extended/pkg/files"
	"github.com/VirtusLab/render/renderer"
	"github.com/VirtusLab/render/renderer/parameters"
)

func Render(file, config string) (string, error) {
	params, err := parameters.FromFiles([]string{config})
	if err != nil {
		return "", err
	}
	template, err := files.ReadInput(file)
	if err != nil {
		return "", err
	}
	var opts []string
	return render(string(template), opts, params)
}

func render(template string, opts []string, params parameters.Parameters) (string, error) {
	return renderer.New(
		renderer.WithOptions(opts...),
		renderer.WithParameters(params),
		renderer.WithSprigFunctions(),
		renderer.WithExtraFunctions(),
		renderer.WithCryptFunctions(),
	).Render(template)
}
