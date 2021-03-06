/*
 * Copyright 2018 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package riff_buildpack

import (
	"fmt"
	"path/filepath"

	"github.com/buildpack/libbuildpack"
	"github.com/cloudfoundry/libjavabuildpack"
)

// Metadata represents the contents of the riff.toml file in an application root
type Metadata struct {
	Handler string `toml:"handler"`
}

// String makes Metadata satisfy the Stringer interface.
func (m Metadata) String() string {
	return fmt.Sprintf("Metadata{ Handler: %s }", m.Handler)
}

// NewMetadata creates a new Metadata from the contents of $APPLICATION_ROOT/riff.toml.  If that file does not exist,
// the second return value is false.
func NewMetadata(application libbuildpack.Application, logger libjavabuildpack.Logger) (Metadata, bool, error) {
	f := filepath.Join(application.Root, "riff.toml")

	exists, err := libjavabuildpack.FileExists(f)
	if err != nil {
		return Metadata{}, false, err
	}

	if !exists {
		return Metadata{}, false, nil
	}

	var metadata Metadata
	err = libjavabuildpack.FromTomlFile(f, &metadata)
	if err != nil {
		return Metadata{}, false, err
	}

	logger.Debug("Riff metadata: %s", metadata)
	return metadata, true, nil
}
