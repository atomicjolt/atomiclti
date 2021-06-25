package webpack

/**
 * This package is a modification of this blog post (as such I have also included the license):
 * https://kimrgrey.medium.com/integration-of-create-react-app-into-golang-server-in-2020-1aff6e93ee5a
 * Demo repo: https://github.com/trickstersio/go-create-react-app
 *
 *	MIT License
 *
 *	Copyright (c) 2018 Sergey Tsvetkov
 *
 *	Permission is hereby granted, free of charge, to any person obtaining a copy
 *	of this software and associated documentation files (the "Software"), to deal
 *	in the Software without restriction, including without limitation the rights
 *	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *	copies of the Software, and to permit persons to whom the Software is
 *	furnished to do so, subject to the following conditions:
 *
 *	The above copyright notice and this permission notice shall be included in all
 *	copies or substantial portions of the Software.
 *
 *	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *	SOFTWARE.
 **/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

type Manifest struct {
	Files       Files `json:"files"`
	Entrypoints Entrypoints
}

type Files map[string]string
type Entrypoints []string

func NewFromDevServer(manifestUrl string) (*Manifest, error) {
	manifest := &Manifest{}

	res, err := http.Get(manifestUrl)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(content, manifest); err != nil {
		return nil, err
	}

	return manifest, nil
}

func NewFromBuildPath(buildPath string) (*Manifest, error) {
	manifest := &Manifest{}
	assetsManifestPath := path.Join(buildPath, "asset-manifest.json")

	if _, err := os.Stat(assetsManifestPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("Did not find manifest at %s.", assetsManifestPath)
	}

	content, err := ioutil.ReadFile(assetsManifestPath)

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(content, &manifest); err != nil {
		return nil, err
	}

	return manifest, nil
}

func (e Entrypoints) Scripts() Entrypoints {
	var scripts Entrypoints

	for _, f := range e {
		if strings.HasSuffix(f, ".js") {
			scripts = append(scripts, f)
		}
	}

	return scripts
}

func (e Entrypoints) Styles() Entrypoints {
	var styles Entrypoints

	for _, f := range e {
		if strings.HasSuffix(f, ".css") {
			styles = append(styles, f)
		}
	}

	return styles
}
