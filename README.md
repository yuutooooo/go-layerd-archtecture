# go-layerd-archtecture

# Go 4層アーキテクチャ（Clean Architecture）

## **ディレクトリ構成**

```
.
├── README.md
├── cmd
│   └── main.go
├── domain
│   ├── model
│   ├── repository # interfaceで実装は行わない
│   ├── service
│   └── usecase
├── handler
├── infrastructure
│   └── repository # ここで実際のリポジトリ層の実装を行う
└── router
```

---

## **アーキテクチャの流れ**
本アーキテクチャでは、**`handler` → `usecase` → `service` → `repository`** の順で処理が流れる。

### **1. `handler`（プレゼンテーション層）**
**役割**
- クライアントからの HTTP リクエストを受け取り、適切な `usecase` を呼び出す。
- リクエストデータのバリデーションを行う。
- `usecase` からのレスポンスを整形し、クライアントに返す。

**依存関係**
- `usecase` に依存
- `usecase` のみを呼び出し、`service` や `repository` には依存しない

---

### **2. `usecase`（ユースケース層）**
**役割**
- アプリケーションのワークフローを管理する。
- 必要に応じて `service` を呼び出し、ドメインサービスを実行する。
- `repository` を利用してデータの取得・保存を行う。

**依存関係**
- `service` に依存（ドメインサービスの実行）
- `repository` に依存（データの取得・保存）
- `handler` に依存される（リクエストの入口）

---

### **3. `service`（ドメインサービス層）**
**役割**
- ドメインに関する **再利用可能なドメインサービス** を実装する。
- `usecase` から呼び出され、`repository` には直接依存しない。
- 例：パスワードのハッシュ化、バリデーション、データ変換。

**依存関係**
- `usecase` に依存される（ユースケースの中で利用される）
- 他の `usecase` でも再利用可能

---

### **4. `repository`（データアクセス層）**
**役割**
- `domain/repository/` にはリポジトリの **インターフェース** を定義する。
- `infrastructure/repository/` では、そのインターフェースの **実装** を行う。
- データの永続化を担当（データの保存・取得・更新・削除）。

**依存関係**
- `usecase` に依存される（データを取得・保存するため）
- `service` には直接依存しない

---

## **依存関係のまとめ**
| 層           | 依存先                            | 依存される先 |
| ------------ | --------------------------------- | ------------ |
| `handler`    | `usecase`                         | クライアント |
| `usecase`    | `service`, `repository`           | `handler`    |
| `service`    | なし (`usecase` からのみ呼ばれる) | `usecase`    |
| `repository` | なし (`usecase` からのみ呼ばれる) | `usecase`    |

---

## **処理の流れ**
### **例: ユーザー登録の流れ**
1. **`handler` がクライアントのリクエストを受け取る**
   - `/users/register` にリクエストが送られる
   - JSON のリクエストデータをパースする
   - `usecase` にデータを渡す

2. **`usecase` がワークフローを管理**
   - `service` を呼び出し、パスワードをハッシュ化する
   - `repository` を呼び出し、データベースに保存する

3. **`service` でドメインサービスを処理**
   - `HashPassword()` を実行し、パスワードを暗号化する

4. **`repository` でデータを保存**
   - `Create(user)` を実行し、DB に保存する

5. **レスポンスを返す**
   - `usecase` → `handler` に処理が戻る
   - `handler` が適切なレスポンスをクライアントに返す

---

## **4層アーキテクチャの利点**
✅ **責務分離**：各レイヤーが特定の責務を持ち、関心事が明確になる  
✅ **再利用性**：`service` のドメインサービスは `usecase` 間で共有可能  
✅ **テスト容易性**：各層をモック化して単体テストが可能  
✅ **拡張性**：`usecase` を増やして新しい機能を追加しやすい  

---

## **まとめ**
- **`handler` → `usecase` → `service` → `repository` の順で処理が流れる**
- **`usecase` がワークフローを管理し、`service` は再利用可能なロジックを提供**
- **`repository` はデータベースとのやり取りのみを担当**