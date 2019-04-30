/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package file

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"strings"

	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/uri"
)

// Scheme for file
const Scheme = "file"

// Artifact returns a file *api.Artifact from a given u
func Artifact(u *uri.URI) (*api.Artifact, error) {

	if u.Scheme != "" && u.Scheme != Scheme {
		return nil, nil
	}

	path := strings.TrimPrefix(u.Opaque, "//")

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return nil, err
	}
	checksum := h.Sum(nil)

	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}

	return &api.Artifact{
		Name: stat.Name(),
		Hash: hex.EncodeToString(checksum),
		Size: uint64(stat.Size()),
	}, nil
}
