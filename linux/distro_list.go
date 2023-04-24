package linux

// ディストリの一覧
// 上の方が優先度が高い
var DistroList []*Linux = []*Linux{
	Arch,
	Debian,
	Ubuntu,
	RHEL,
	CentOS,
	Other,
}
