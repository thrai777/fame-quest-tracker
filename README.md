# fame-quest-tracker

『ファイアーエムブレム 万紫千紅』の「人助け」クエストと名声ポイントを管理する、個人用のWebアプリです。

名声レベルは周回に引き継がれないため、ルートごとに「期日までに・あと何ポイント必要か」を把握するのが地味に大変です。このアプリはその管理を肩代わりします。

## 機能（予定）

- クエスト一覧：名前・名声pt・ゲーム内期日・状態（未受注／進行中／完了／失敗）
- 納品系クエストの進捗カウンタ（例：野菜 16/30）
- 名声レベル計算：現在のptから現在レベルと次レベルまでの残りptを表示
- 期日アラート：現在のゲーム内日付と残りターンから、危ない順に並び替え
- データの保存とJSONエクスポート／インポート

## 技術構成

| 層 | 技術 |
|---|---|
| API | Go (net/http, 標準ライブラリ中心) |
| フロントエンド | React + TypeScript + Vite（`web/`、実装予定） |
| インフラ | Cloudflare Pages（フロント）＋ Cloudflare Tunnel（API） |
| CI | GitHub Actions（go vet / go test） |

## セットアップ

```bash
# APIサーバー起動（デフォルト :8080）
go run ./cmd/server

# 動作確認
curl http://localhost:8080/health
curl http://localhost:8080/api/quests
```

## API

| メソッド | パス | 説明 |
|---|---|---|
| GET | `/health` | ヘルスチェック |
| GET | `/api/quests` | クエスト一覧 |

## テスト

```bash
go vet ./...
go test ./...
```

## ロードマップ

- [x] v0.0 プロジェクト雛形（ヘルスチェック＋クエスト一覧API）
- [ ] クエストの状態変更・追加API
- [ ] Reactフロント（一覧画面）
- [ ] 納品カウンタと名声レベル計算
- [ ] JSONファイルへの永続化とエクスポート／インポート
- [ ] Cloudflare Pages公開（v0.1）
- [ ] 独自ドメイン＋Tunnel＋Access（v0.2）

## ライセンス

MIT
