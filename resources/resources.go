// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package resources

import (
	"embed"
	"fmt"
	"io"
)

//go:embed alterations.txt namelist.txt
var resourceFS embed.FS

func GetResourceFile(path string) (io.Reader, error) {
	file, err := resourceFS.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to obtain the embedded file: %s: %v", path, err)
	}
	return file, err
}

func GetResourceFileData(path string) ([]byte, error) {
	return resourceFS.ReadFile(path)
}
