package common

type FileSize = uint

const (
	KB FileSize = 1 << 10
	MB FileSize = KB << 10
	GB FileSize = MB << 10
)

const (
	GitHubMaxFileSize            FileSize = 99 * MB
	GitHubMaxRecommendedFileSize FileSize = 49 * MB
	GithubRepoMaxSize            FileSize = 1 * GB
)

const (
	SplitSuffix string = "part"
)

const (
	DefaultReadBufferSize FileSize = 16 * MB
)
