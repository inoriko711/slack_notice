# slack_notice
- このリポジトリは[Slack Block Kit](https://api.slack.com/block-kit)の要件を満たす構造体を提供します

# 構成
- types.go
  - [Blocks](https://api.slack.com/reference/block-kit/blocks)と[Block Elements](https://api.slack.com/reference/block-kit/block-elements)、[Composition objects](https://api.slack.com/reference/block-kit/composition-objects)の要件を満たす構造体を提供しております

- sample
  - 構造体の使用例になります
  - sample/sample1.go はてなブログで書かれている「癒し画像を投稿する」の実装例になります
  - sample/sample2.go Multi-select menu elementの実装例になります

- sample/data
  - data.json
  - sample1/sample1.goで使用している、投稿内容を管理するファイルです

# その他
- このリポジトリははてなブログと連携しております
- https://inoriko.hatenablog.com/entry/2020/12/08/090000
