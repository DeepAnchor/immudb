/*
Copyright 2021 CodeNotary, Inc. All rights reserved.

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

package server

import "errors"

var ErrUnknowMessageType = errors.New("found an unknown message type on the wire")
var ErrDBNotprovided = errors.New("database name not provided")
var ErrUsernameNotprovided = errors.New("user name not provided")
var ErrPwNotprovided = errors.New("password not provided")
var ErrDBNotExists = errors.New("selected db doesn't exists")
var ErrUsernameNotFound = errors.New("user not found")
