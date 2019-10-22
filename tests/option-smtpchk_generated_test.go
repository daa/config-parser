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

package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/haproxytech/config-parser/v2/parsers"
)


func TestOptionSmtpchkNormal0(t *testing.T) {
	parser := &parsers.OptionSmtpchk{}
	line := strings.TrimSpace("option smtpchk")
	err := ProcessLine(line, parser)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	result, err := parser.Result()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	var returnLine string
	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}
	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}
}
func TestOptionSmtpchkNormal1(t *testing.T) {
	parser := &parsers.OptionSmtpchk{}
	line := strings.TrimSpace("no option smtpchk")
	err := ProcessLine(line, parser)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	result, err := parser.Result()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	var returnLine string
	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}
	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}
}
func TestOptionSmtpchkNormal2(t *testing.T) {
	parser := &parsers.OptionSmtpchk{}
	line := strings.TrimSpace("option smtpchk HELO mydomain.org")
	err := ProcessLine(line, parser)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	result, err := parser.Result()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	var returnLine string
	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}
	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}
}
func TestOptionSmtpchkNormal3(t *testing.T) {
	parser := &parsers.OptionSmtpchk{}
	line := strings.TrimSpace("option smtpchk EHLO mydomain.org")
	err := ProcessLine(line, parser)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	result, err := parser.Result()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	var returnLine string
	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}
	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}
}
func TestOptionSmtpchkNormal4(t *testing.T) {
	parser := &parsers.OptionSmtpchk{}
	line := strings.TrimSpace("option smtpchk # comment")
	err := ProcessLine(line, parser)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	result, err := parser.Result()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	var returnLine string
	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}
	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}
}
func TestOptionSmtpchkNormal5(t *testing.T) {
	parser := &parsers.OptionSmtpchk{}
	line := strings.TrimSpace("option smtpchk HELO mydomain.org # comment")
	err := ProcessLine(line, parser)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	result, err := parser.Result()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	var returnLine string
	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}
	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}
}
func TestOptionSmtpchkFail0(t *testing.T) {
	parser := &parsers.OptionSmtpchk{}
	line := strings.TrimSpace("---")
	err := ProcessLine(line, parser)
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error for line [%s]", line))
	}
	_, err = parser.Result()
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error on result for line [%s]", line))
	}
}
func TestOptionSmtpchkFail1(t *testing.T) {
	parser := &parsers.OptionSmtpchk{}
	line := strings.TrimSpace("--- ---")
	err := ProcessLine(line, parser)
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error for line [%s]", line))
	}
	_, err = parser.Result()
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error on result for line [%s]", line))
	}
}
