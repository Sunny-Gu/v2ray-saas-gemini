# API 文档

本文档详细描述了 V2Ray SaaS 平台的 API 接口。

**基础 URL**: `http://localhost:<PORT>`

- **门户 API 端口**: `8080`
- **后台管理 API 端口**: `8081`

---

## 门户 API (`portal-api`)

所有门户 API 接口均以 `/api/v1` 为前缀。

### 认证

- **公开路由**: 无需认证即可访问。
- **认证路由**: 需要在 `Authorization` 请求头中提供 `Bearer Token`。

### 用户管理 (`/user`)

#### **`POST /user/register`**
- **描述**: 注册一个新用户。
- **认证**: 公开
- **请求体**:
  ```json
  {
    "email": "user@example.com",
    "password": "password123"
  }
  ```
- **成功响应 (201)**:
  ```json
  {
    "message": "User registered successfully",
    "user_id": 1
  }
  ```

#### **`POST /user/login`**
- **描述**: 用户登录并返回 JWT。
- **认证**: 公开
- **请求体**:
  ```json
  {
    "email": "user@example.com",
    "password": "password123"
  }
  ```
- **成功响应 (200)**:
  ```json
  {
    "message": "Login successful",
    "token": "ey..."
  }
  ```

#### **`POST /user/request-password-reset`**
- **描述**: 为用户发起密码重置流程。
- **认证**: 公开
- **请求体**:
  ```json
  {
    "email": "user@example.com"
  }
  ```
- **成功响应 (200)**:
  ```json
  {
    "message": "If an account with that email exists, a password reset link has been sent."
  }
  ```

#### **`POST /user/reset-password`**
- **描述**: 使用有效的令牌重置用户密码。
- **认证**: 公开
- **请求体**:
  ```json
  {
    "token": "valid-reset-token",
    "new_password": "newStrongPassword123"
  }
  ```
- **成功响应 (200)**:
  ```json
  {
    "message": "Password has been reset successfully."
  }
  ```

### 个人中心 (`/profile`)

#### **`GET /profile/`**
- **描述**: 获取当前认证用户的个人信息。
- **认证**: 需要
- **成功响应 (200)**:
  ```json
  {
    "email": "user@example.com",
    "balance": 100.50,
    "plan_name": "基础套餐",
    "plan_expiry_date": "2025-12-31",
    "used_traffic_gb": 50.5,
    "total_traffic_gb": 200.0
  }
  ```

#### **`GET /profile/history`**
- **描述**: 获取用户的消费历史（订阅订单）。
- **认证**: 需要
- **成功响应 (200)**:
  ```json
  [
    {
      "ID": 1,
      "CreatedAt": "...",
      "UserID": 1,
      "PlanID": 1,
      "Amount": 50.0,
      "FinalAmount": 50.0,
      "Status": "completed"
    }
  ]
  ```

### 商店与订阅 (`/store`, `/subscription`)

#### **`GET /store/plans`**
- **描述**: 列出所有可购买的套餐计划。
- **认证**: 公开
- **成功响应 (200)**:
  ```json
  [
    {
      "ID": 1,
      "Name": "基础套餐",
      "Description": "一个非常棒的入门套餐。",
      "Price": 50.0,
      "TrafficGB": 200.0,
      "DurationDays": 30,
      "IsActive": true,
      "IsFeatured": true
    }
  ]
  ```

#### **`POST /store/purchase`**
- **描述**: 使用账户余额购买一个套餐。
- **认证**: 需要
- **请求体**:
  ```json
  {
    "plan_id": 1
  }
  ```
- **成功响应 (201)**:
  ```json
  {
    "message": "Plan purchased successfully",
    "subscription": { ... }
  }
  ```

#### **`GET /subscription/`**
- **描述**: 获取用户当前有效的订阅详情。
- **认证**: 需要
- **成功响应 (200)**:
  ```json
  {
    "ID": 1,
    "UserID": 1,
    "PlanID": 1,
    "ExpiredAt": "...",
    "TotalTrafficGB": 200.0,
    "UsedTrafficGB": 50.5,
    "SubscriptionURL": "unique-uuid-string"
  }
  ```

#### **`POST /subscription/reset-link`**
- **描述**: 重置用户的订阅链接。
- **认证**: 需要
- **成功响应 (200)**:
  ```json
  {
    "message": "Subscription link reset successfully",
    "subscription": { ... }
  }
  ```

### 充值 (`/recharge`)

#### **`GET /store/recharge-presets`**
- **描述**: 列出所有可用的充值选项。
- **认证**: 公开
- **成功响应 (200)**:
  ```json
  [
    {
      "ID": 1,
      "AmountCNY": 50.0,
      "Description": "50元充值包",
      "IsEnabled": true
    }
  ]
  ```

#### **`POST /recharge/create-order`**
- **描述**: 创建一个新的 USDT 充值订单。
- **认证**: 需要
- **请求体**:
  ```json
  {
    "preset_id": 1
  }
  ```
- **成功响应 (201)**:
  ```json
  {
    "ID": 1,
    "UserID": 1,
    "AmountCNY": 50.0,
    "PaidUSDT": 5.56,
    "PaymentAddress": "T...",
    "Status": "pending"
  }
  ```

#### **`POST /recharge/redeem`**
- **描述**: 兑换一个充值码。
- **认证**: 需要
- **请求体**:
  ```json
  {
    "code": "VALID-COUPON-CODE"
  }
  ```
- **成功响应 (200)**:
  ```json
  {
    "message": "Coupon redeemed successfully"
  }
  ```

#### **`POST /recharge/generate-coupon`**
- **描述**: 从用户余额生成一个新的充值码。
- **认证**: 需要
- **请求体**:
  ```json
  {
    "value": 20.0
  }
  ```
- **成功响应 (201)**:
  ```json
  {
    "ID": 1,
    "Code": "U-unique-uuid-string",
    "Value": 20.0,
    "Status": "active"
  }
  ```

### 其他公开接口

- **`GET /nodes/status`**: 列出所有服务器节点的状态。
- **`GET /announcements`**: 列出所有有效的公告。
- **`GET /help-documents`**: 列出所有帮助文档。

---

## 后台管理 API (`admin-api`)

所有后台管理 API 接口均以 `/api/v1/admin` 为前缀。除 `/login` 外，所有接口都需要管理员登录后获得的 `Bearer Token`。

### 认证

#### **`POST /login`**
- **描述**: 管理员登录。
- **请求体**:
  ```json
  {
    "username": "admin",
    "password": "adminpassword"
  }
  ```
- **成功响应 (200)**:
  ```json
  {
    "token": "ey..."
  }
  ```

### 节点管理 (`/nodes`)

- **`POST /`**: 创建一个新节点。
- **`GET /`**: 列出所有节点。
- **`GET /:id`**: 获取单个节点详情。
- **`PUT /:id`**: 更新一个节点。
- **`DELETE /:id`**: 删除一个节点。

### 用户管理 (`/users`)

- **`GET /`**: 列出用户，支持分页和搜索 (`?email=...&page=1&page_size=20`)。
- **`PATCH /:id/status`**: 更新用户状态 (例如, `{"status": "locked"}`)。

### 套餐管理 (`/plans`)

- **`POST /`**: 创建一个新套餐。
- **`GET /`**: 列出所有套餐。
- **`GET /:id`**: 获取单个套餐详情。
- **`PUT /:id`**: 更新一个套餐。
- **`DELETE /:id`**: 删除一个套餐。

### 订单管理 (`/orders`)

- **`GET /recharge`**: 分页列出充值订单。
- **`GET /subscription`**: 分页列出订阅订单。

### 内容管理 (`/content`)

- **`POST /announcements`**: 创建新公告。
- **`PUT /announcements/:id`**: 更新公告。
- **`DELETE /announcements/:id`**: 删除公告。
- **`POST /help-documents`**: 创建新帮助文档。
- **`PUT /help-documents/:id`**: 更新帮助文档。
- **`DELETE /help-documents/:id`**: 删除帮助文档。

### 资金审计 (`/audit`)

- **`GET /financial-overview`**: 获取平台资金总览。