package linux

// ディストリの一覧
// 上の方が優先度が高い
// 派生ディストリは親ディストリより上に配置する
var DistroList []*Linux = []*Linux{
	// Arch系
	Manjaro,
	Arch,
	// Debian系
	LinuxMint,
	Ubuntu,
	Debian,
	// Red Hat系
	Fedora,
	CentOS,
	RHEL,
	// SUSE系
	OpenSUSETumbleweed,
	OpenSUSELeap,
	// 独立系
	Alpine,
	Gentoo,
	// フォールバック
	Other,
}
