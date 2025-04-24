package main

type Symbols struct {
	Lock                 string
	Network              string
	Separator            string
	SeparatorThin        string
	SeparatorReverse     string
	SeparatorReverseThin string

	RepoDetached   string
	RepoAhead      string
	RepoBehind     string
	RepoStaged     string
	RepoNotStaged  string
	RepoUntracked  string
	RepoConflicted string
	RepoStashed    string
}

type Theme struct {
	Reset          uint8
	UsernameFg     uint8
	UsernameBg     uint8
	UsernameRootBg uint8

	// The foreground-background mapping is precomputed and stored in a map for improved performance
	// The old script used to brute-force this at runtime
	HostnameColorizedFgMap map[uint8]uint8

	HomeSpecialDisplay bool
	HomeFg             uint8
	HomeBg             uint8
	AliasFg            uint8
	AliasBg            uint8
	PathFg             uint8
	PathBg             uint8
	CwdFg              uint8
	SeparatorFg        uint8

	ReadonlyFg uint8
	ReadonlyBg uint8

	SshFg uint8
	SshBg uint8

	KubeClusterFg   uint8
	KubeClusterBg   uint8
	KubeNamespaceFg uint8
	KubeNamespaceBg uint8

	DotEnvFg uint8
	DotEnvBg uint8

	RepoCleanFg uint8
	RepoCleanBg uint8
	RepoDirtyFg uint8
	RepoDirtyBg uint8

	JobsFg uint8
	JobsBg uint8

	CmdPassedFg uint8
	CmdPassedBg uint8
	CmdFailedFg uint8
	CmdFailedBg uint8

	GitAheadFg      uint8
	GitAheadBg      uint8
	GitBehindFg     uint8
	GitBehindBg     uint8
	GitStagedFg     uint8
	GitStagedBg     uint8
	GitNotStagedFg  uint8
	GitNotStagedBg  uint8
	GitUntrackedFg  uint8
	GitUntrackedBg  uint8
	GitConflictedFg uint8
	GitConflictedBg uint8
	GitStashedFg    uint8
	GitStashedBg    uint8

	TFWsFg uint8
	TFWsBg uint8

	ShellVarFg uint8
	ShellVarBg uint8

	NodeFg uint8
	NodeBg uint8
}
