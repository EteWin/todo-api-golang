# TODO API (Go + Clean Architecture)

## ■ 概要

このリポジトリは、Go + GORM + Gin + Docker を用いた
シンプルなTODO APIです。

クリーンアーキテクチャに基づき、以下の構造で実装しています。
（後にフロントエンド部分を実装予定です）

## ■ 使用技術

- Go
- Gin
- GORM
- PostgreSQL
- Docker / Docker Compose

## ■ ディレクトリ構成

```
・
├── backend/
│   ├── Dockerfile
│   ├── cmd/
│   │   └── main.go
│   ├── internal/
│   │   ├── domain/          # ビジネスロジック
│   │   ├── usecase/         # ユースケース（処理の流れ）
│   │   ├── interface/       # 外部との接点（HTTP）
│   │   └── infrastructure/  # DBなどの実装
│          └── persistence/  # Usecaseの保存係
│
├── docker-compose.yml
├── .env
├── .env.example
└── README.md
```

## ■ アーキテクチャ

Domain ← Usecase ← Interface ← Infrastructure

### 各層の役割

- Domain
  - ビジネスルールを定義
  - 外部（DB・FW）に依存しない
- Usecase
- アプリケーションの処理を定義
- Domainを利用する
- Interface
- HTTPリクエスト/レスポンスの変換
- Usecaseを呼び出す
- Infrastructure
- DBや外部サービスの実装
- GORMなどを使用

## ■ 起動方法

① リポジトリをクローン

- git clone <repository_url>
- cd <project_name>

② Docker起動

- docker compose build
- docker compose up

③ API確認

- http://localhost:8080

## ■ 補足

- IDはUUIDで管理されています
- DBはPostgreSQLを使用しています
- マイグレーションはGORMのAutoMigrateを使用しています

## ■ 開発メモ

- クリーンアーキテクチャは「依存関係の方向」が重要
- フォルダ構成よりも責務分離を優先

### ビジネスロジックの流れ

1. HTTP リクエスト
2. Handler
3. Usecase
4. Domain
5. Repository
6. Infrastructure

### メモ書き

- Domain（どのように振る舞うかだけを意識すると良さそう）
  - ビジネスルールそのもの
  - エンティティ、値オブジェクト
  - 外部技術を一切知らない

- Usecase（何を行うのか定義する）
  - ユーザーの操作単位の処理
  - Domainを組み合わせる
  - Repositoryインターフェースを用いる

- Interface（外界との橋渡し）
  - HTTPやCLIの受け口
  - Usecaseの呼び出しのみ

- Infrastructure（どのように実現したいか）
  - DB接続
  - ORM
  - 外部API
