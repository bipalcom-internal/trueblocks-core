package migrate

// TODO: BOGUS - MIGRATION STUFF
func CheckBackLevelIndex(chain string, missingOkay bool) {
	// fileName := config.GetPathToIndex(chain) + "finalized/000000000-000000000.bin"
	// 	if !file.FileExists(fileName) && missingOkay {
	// 		return
	// 	}
	// 	ok, err := index.HasValidHeader(chain, fileName)
	// 	if ok && err == nil {
	// 		return
	// 	}

	// 	const BackLevelVersion string = `

	//   A back-level version of an index file was found at

	//     {[{FILE}]}

	//   Please carefully follow all migrations up to and including {0}
	//   before proceeding.

	//   See https://github.com/TrueBlocks/trueblocks-core/blob/develop/MIGRATIONS.md

	//   [{VERSION}]

	// `
	// 	msg := strings.Replace(BackLevelVersion, "{0}", "{v0.40.0-beta}", -1)
	// 	msg = strings.Replace(msg, "[{VERSION}]", version.LibraryVersion, -1)
	// 	msg = strings.Replace(msg, "[{FILE}]", fileName, -1)
	// 	msg = strings.Replace(msg, "{", colors.Green, -1)
	// 	msg = strings.Replace(msg, "}", colors.Off, -1)
	// 	log.Fatalf(msg)
}
