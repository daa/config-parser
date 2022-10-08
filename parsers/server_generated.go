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

func (p *Server) Init() {
	p.data = []types.Server{}
	p.preComments = []string{}
}

func (p *Server) GetParserName() string {
	return "server"
}

func (p *Server) Get(createIfNotExist bool) (common.ParserData, error) {
	if len(p.data) == 0 && !createIfNotExist {
		return nil, errors.ErrFetch
	}
	return p.data, nil
}

func (p *Server) GetPreComments() ([]string, error) {
	return p.preComments, nil
}

func (p *Server) SetPreComments(preComments []string) {
	p.preComments = preComments
}

func (p *Server) GetOne(index int) (common.ParserData, error) {
	if index < 0 || index >= len(p.data) {
		return nil, errors.ErrFetch
	}
	return p.data[index], nil
}

func (p *Server) Delete(index int) error {
	if index < 0 || index >= len(p.data) {
		return errors.ErrFetch
	}
	copy(p.data[index:], p.data[index+1:])
	p.data[len(p.data)-1] = types.Server{}
	p.data = p.data[:len(p.data)-1]
	return nil
}

func (p *Server) Insert(data common.ParserData, index int) error {
	if data == nil {
		return errors.ErrInvalidData
	}
	switch newValue := data.(type) {
	case []types.Server:
		p.data = newValue
	case *types.Server:
		if index > -1 {
			if index > len(p.data) {
				return errors.ErrIndexOutOfRange
			}
			p.data = append(p.data, types.Server{})
			copy(p.data[index+1:], p.data[index:])
			p.data[index] = *newValue
		} else {
			p.data = append(p.data, *newValue)
		}
	case types.Server:
		if index > -1 {
			if index > len(p.data) {
				return errors.ErrIndexOutOfRange
			}
			p.data = append(p.data, types.Server{})
			copy(p.data[index+1:], p.data[index:])
			p.data[index] = newValue
		} else {
			p.data = append(p.data, newValue)
		}
	default:
		return errors.ErrInvalidData
	}
	return nil
}

func (p *Server) Set(data common.ParserData, index int) error {
	if data == nil {
		p.Init()
		return nil
	}
	switch newValue := data.(type) {
	case []types.Server:
		p.data = newValue
	case *types.Server:
		if index > -1 && index < len(p.data) {
			p.data[index] = *newValue
		} else if index == -1 {
			p.data = append(p.data, *newValue)
		} else {
			return errors.ErrIndexOutOfRange
		}
	case types.Server:
		if index > -1 && index < len(p.data) {
			p.data[index] = newValue
		} else if index == -1 {
			p.data = append(p.data, newValue)
		} else {
			return errors.ErrIndexOutOfRange
		}
	default:
		return errors.ErrInvalidData
	}
	return nil
}

func (p *Server) PreParse(line string, parts []string, preComments []string, comment string) (string, error) {
	changeState, err := p.Parse(line, parts, comment)
	if err == nil && preComments != nil {
		p.preComments = append(p.preComments, preComments...)
	}
	return changeState, err
}

func (p *Server) Parse(line string, parts []string, comment string) (string, error) {
	if parts[0] == "server" {
		data, err := p.parse(line, parts, comment)
		if err != nil {
			if _, ok := err.(*errors.ParseError); ok {
				return "", err
			}
			return "", &errors.ParseError{Parser: "Server", Line: line}
		}
		p.data = append(p.data, *data)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "Server", Line: line}
}

func (p *Server) ResultAll() ([]common.ReturnResultLine, []string, error) {
	res, err := p.Result()
	return res, p.preComments, err
}
