// Copyright 2018 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package crane_test

import (
	"fmt"

	"github.com/schidstorm/go-containerregistry/pkg/crane"
)

func Example() {
	c := map[string][]byte{
		"/binary": []byte("binary contents"),
	}
	i, _ := crane.Image(c)
	d, _ := i.Digest()
	fmt.Println(d)
	// Output: sha256:09fb0c6289cefaad8c74c7e5fd6758ad6906ab8f57f1350d9f4eb5a7df45ff8b
}
