package receiptsPkg

/*-------------------------------------------------------------------------------------------
 * qblocks - fast, easily-accessible, fully-decentralized data from blockchains
 * copyright (c) 2016, 2021 TrueBlocks, LLC (http://trueblocks.io)
 *
 * This program is free software: you may redistribute it and/or modify it under the terms
 * of the GNU General Public License as published by the Free Software Foundation, either
 * version 3 of the License, or (at your option) any later version. This program is
 * distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even
 * the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * General Public License for more details. You should have received a copy of the GNU General
 * Public License along with this program. If not, see http://www.gnu.org/licenses/.
 *-------------------------------------------------------------------------------------------*/
/*
 * Parts of this file were generated with makeClass --run. Edit only those parts of
 * the code inside of 'EXISTING_CODE' tags.
 */

// EXISTING_CODE
import (
	"net/http"

	"github.com/spf13/cobra"
)

// EXISTING_CODE

var Options ReceiptsOptions

func RunReceipts(cmd *cobra.Command, args []string) error {
	Options.Transactions = args
	opts := Options

	// EXISTING_CODE
	// EXISTING_CODE

	err := opts.ValidateReceipts()
	if err != nil {
		return err
	}

	// EXISTING_CODE
	return opts.Globals.PassItOn("getReceipts", opts.ToDashStr())
	// EXISTING_CODE
}

func ServeReceipts(w http.ResponseWriter, r *http.Request) {
	opts := FromRequest(w, r)

	// EXISTING_CODE
	// EXISTING_CODE

	err := opts.ValidateReceipts()
	if err != nil {
		opts.Globals.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	// EXISTING_CODE
	opts.Globals.PassItOn("getReceipts", opts.ToDashStr())
	// EXISTING_CODE
}

// EXISTING_CODE
// EXISTING_CODE
