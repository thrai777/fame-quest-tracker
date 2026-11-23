# web

React + TypeScript + Vite のフロントエンドです。

## 開発

前提: Go API が `http://localhost:8080` で起動していること（リポジトリ直下で `go run ./cmd/server`）。

```bash
npm install
npm run dev   # http://localhost:5173 （/api と /health は Vite の proxy で API へ転送）
npm run build # 型チェック + 本番ビルド
```
