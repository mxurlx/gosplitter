package common

var Flags = map[string]map[string][]string{
	".": {
		"help":    {"h", "false", "Show help message"},
		"version": {"v", "false", "Show version"},
	},
	"split": {
		"input":     {"1", "<mandatory>", "Input file"},
		"output":    {"o", "", "Output directory"},
		"suffix":    {"s", "part", "Chunk suffix"},
		"chunksize": {"c", "4096", "Chunk size"},
	},
	"merge": {
		"input":  {"1", "<mandatory>", "Input directory"},
		"output": {"o", "", "Output file"},
	},
}
