package mock

import (
	"github.com/spf13/afero"
)

func AferoFs() afero.Fs {
	fs := afero.NewMemMapFs()

	dirs := []string{
		"media/photos",
		"media/video",
		"media/music",
		"text/plain",
		"text/pdf",
		"tree/a/a_a",
		"tree/a/a_b",
		"tree/a/a_b/a_b_a",
		"tree/a/a_c/a_c_a",
	}
	for _, dir := range dirs {
		err := fs.Mkdir(dir, 0755)
		if err != nil {
			panic(err)
		}

	}
	return fs
}

func AferoHttpFs() *afero.HttpFs {
	fs := AferoFs()
	return afero.NewHttpFs(fs)
}
