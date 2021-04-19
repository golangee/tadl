// Copyright 2021 Torben Schinke
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ast

// Service can be many things, like any stub. It defines members for
// configuration, injection and custom fields. Afterwards methods
// are defined just like in an interface.
type Service struct {
	// Doc contains a summary, arbitrary text lines, captions, sections and more.
	Doc       DocTypeBlock        `parser:"@@"`
	Name      Ident               `"service" @@ "{"`
	Configure []*FieldWithDefault `("configure" "{" @@* "}")?`
	Inject    []*Field            `("inject" "{" @@* "}")?`
	Private   []*Field            `("private" "{" @@* "}")?`
	Methods   []*Method           ` @@* "}"`
}