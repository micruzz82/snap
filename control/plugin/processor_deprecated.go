/*          **  DEPRECATED  **
For more information, see our deprecation notice
on Github: https://github.com/micruzz82/snap/issues/1289
*/

/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015 Intel Corporation

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

package plugin

import "github.com/micruzz82/snap/core/ctypes"

// Processor plugin
type ProcessorPlugin interface {
	Plugin
	Process(contentType string, content []byte, config map[string]ctypes.ConfigValue) (string, []byte, error)
}
