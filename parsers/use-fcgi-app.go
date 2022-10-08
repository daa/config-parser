/*
Copyright 2019 HAProxy Technologies

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package parsers

import (
	"fmt"

	"github.com/haproxytech/config-parser/v4/common"
	"github.com/haproxytech/config-parser/v4/errors"
	"github.com/haproxytech/config-parser/v4/types"
)

type UseFcgiApp struct {
	data        *types.UseFcgiApp
	preComments []string // comments that appear before the the actual line
}

func (s *UseFcgiApp) Parse(line string, parts []string, comment string) (string, error) {
	if len(parts) != 2 {
		return "", &errors.ParseError{Parser: "use-fcgi-app", Line: line}
	}
	if parts[0] != "use-fcgi-app" {
		return "", &errors.ParseError{Parser: "use-fcgi-app", Line: line}
	}
	s.data = &types.UseFcgiApp{
		Name:    parts[1],
		Comment: comment,
	}
	return "", nil
}

func (s *UseFcgiApp) Result() ([]common.ReturnResultLine, error) {
	if s.data == nil {
		return nil, errors.ErrFetch
	}
	return []common.ReturnResultLine{
		{
			Data:    fmt.Sprintf("use-fcgi-app %s", s.data.Name),
			Comment: s.data.Comment,
		},
	}, nil
}
