// Kodex (Community Edition - CE) - Privacy & Security Engineering Platform
// Copyright (C) 2020  KIProtect GmbH (HRB 208395B) - Germany
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package fixtures

import (
	"github.com/kiprotect/kodex/helpers"
	"github.com/kiprotect/kodex"
)

type Settings struct {
}

func (c Settings) Setup(fixtures map[string]interface{}) (interface{}, error) {
	// we set the loglevel to 'debug' so we can see which settings files are being loaded
	kodex.Log.SetLevel(kodex.DebugLogLevel)
	return helpers.Settings(helpers.SettingsPaths())
}

func (c Settings) Teardown(fixture interface{}) error {
	return nil
}
