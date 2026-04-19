# ジャッジメントダイス

Ebitengine で実装したサイコロゲーム。遊戯王の「ジャッジメントダイス」と、動画配信者「タンクトップ小隊」を題材にした 3 つのモードを収録。

## モード

1. **遊戯王ジャッジメントダイス** — サイコロを振り、出目 1〜6 に対応する効果を表示。
2. **タンクトップ小隊のジャッジメントダイス** — CPU と同時にサイコロを振り、相手より小さい目を出せば勝利。
3. **タンクトップ小隊ジェンガモード** — 1 ならこちらのターン、2〜6 なら相手のターンを判定。

## 動作環境

- Go 1.25.1 以上
- [Ebitengine v2](https://ebitengine.org/)

## 起動方法

```sh
go run .
```

## 操作

| シーン | キー | 動作 |
| --- | --- | --- |
| タイトル | `1` / `2` / `3` | 各モードへ遷移 |
| 共通 | `ESC` | タイトルへ戻る |
| 効果モード | `SPACE` | ダイスを振る |
| 効果モード | `R` | リセット |
| 対戦 / ジェンガ | `SPACE` | 開始 / 次の勝負 |

## ディレクトリ構成

```
.
├── main.go              # エントリポイント
└── internal/game/       # ゲームロジック（Dice / Duel / Jenga / Effects）
```

## Web 版

- 公開 URL: https://rin2yh.github.io/judgement-dice/
- `main` へ push すると GitHub Actions が wasm ビルド＆ GitHub Pages へ自動デプロイする。
- ローカル検証:

```sh
GOOS=js GOARCH=wasm go build -o web/main.wasm ./
cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" web/
(cd web && python3 -m http.server 8080)
# ブラウザで http://localhost:8080/ を開く
```

## テスト

```sh
go test ./...
```
