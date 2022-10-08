// Code generated by go generate; DO NOT EDIT.
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
	"github.com/haproxytech/config-parser/v4/common"
	"github.com/haproxytech/config-parser/v4/errors"
	"github.com/haproxytech/config-parser/v4/types"
)

func (p *UseFcgiApp) Init() {
	p.data = nil
	p.preComments = []string{}
}

func (p *UseFcgiApp) GetParserName() string {
	return "use-fcgi-app"
}

func (p *UseFcgiApp) Get(createIfNotExist bool) (common.ParserData, error) {
	if p.data == nil {
		if createIfNotExist {
			p.data = &types.UseFcgiApp{}
			return p.data, nil
		}
		return nil, errors.ErrFetch
	}
	return p.data, nil
}

func (p *UseFcgiApp) GetPreComments() ([]string, error) {
	return p.preComments, nil
}

func (p *UseFcgiApp) SetPreComments(preComments []string) {
	p.preComments = preComments
}

func (p *UseFcgiApp) GetOne(index int) (common.ParserData, error) {
	if index > 0 {
		return nil, errors.ErrFetch
	}
	if p.data == nil {
		return nil, errors.ErrFetch
	}
	return p.data, nil
}

func (p *UseFcgiApp) Delete(index int) error {
	p.Init()
	return nil
}

func (p *UseFcgiApp) Insert(data common.ParserData, index int) error {
	return p.Set(data, index)
}

func (p *UseFcgiApp) Set(data common.ParserData, index int) error {
	if data == nil {
		p.Init()
		return nil
	}
	switch newValue := data.(type) {
	case *types.UseFcgiApp:
		p.data = newValue
	case types.UseFcgiApp:
		p.data = &newValue
	default:
		return errors.ErrInvalidData
	}
	return nil
}

func (p *UseFcgiApp) PreParse(line string, parts []string, preComments []string, comment string) (string, error) {
	changeState, err := p.Parse(line, parts, comment)
	if err == nil && preComments != nil {
		p.preComments = append(p.preComments, preComments...)
	}
	return changeState, err
}

func (p *UseFcgiApp) ResultAll() ([]common.ReturnResultLine, []string, error) {
	res, err := p.Result()
	return res, p.preComments, err
}
