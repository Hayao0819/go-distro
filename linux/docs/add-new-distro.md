## 新しいディストリビューションを追加するには

このモジュールでは`Linux`構造体を用いることで新しいディストリビューションを定義できます。

### 1.構造体を定義する

以下を参考にして構造体を定義してください。このモジュールでは`ostype.V`は`Version`として実装されています。

```go
type Linux struct {
	id      string          // /etc/os-releaseのid
	name    string          // 小文字+大文字 空白あり
	verfunc func() ostype.V // バージョン情報を取得する関数
	require func() bool     // このディストリビューションであるかを判定する関数
}

type Version struct {
	id       string // 空白を含まない小文字の文字列
	codename string // 空白を含む表示用の文字列
}
```

#### 例

[Arch Linux](/linux/distro_arch.go)を参考にしてください。


### 2.一覧に追加する

[ディストリビューション一覧](/linux/distro_list.go)に追加してください。
