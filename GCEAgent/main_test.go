//  Copyright 2017 Google Inc. All Rights Reserved.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

// +build test

package main

import (
	"errors"
	"net"
	"testing"
)

var errRegNotExist = errors.New("error")

func TestContainsString(t *testing.T) {
	table := []struct {
		a     string
		slice []string
		want  bool
	}{
		{"a", []string{"a", "b"}, true},
		{"c", []string{"a", "b"}, false},
	}
	for _, tt := range table {
		if got, want := containsString(tt.a, tt.slice), tt.want; got != want {
			t.Errorf("containsString(%s, %v) incorrect return: got %v, want %t", tt.a, tt.slice, got, want)
		}
	}
}

func resetPwd(username, pwd string) error {
	return nil
}

func createAdminUser(username, pwd string) error {
	return nil
}

func readRegMultiString(key, name string) ([]string, error) {
	return nil, nil
}

func writeRegMultiString(key, name string, value []string) error {
	return nil
}

func addIPAddress(ip, mask net.IP, index int) error {
	return nil
}

func deleteIPAddress(ip net.IP) error {
	return nil
}
