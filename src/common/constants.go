package common

type FileSize = uint

const (
	KB FileSize = 1 << 10
	MB          = KB << 10
	GB          = MB << 10
)

const (
	GitHubMaxFileSize            FileSize = 99 * MB
	GitHubMaxRecommendedFileSize          = 49 * MB
	GithubRepoMaxSize                     = 1 * GB
)

const (
	SplitSuffix string = "part"
)
