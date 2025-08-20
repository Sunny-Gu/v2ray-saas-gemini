# API Documentation

This document provides a detailed description of the API endpoints for the V2Ray SaaS platform.

**Base URL**: `http://localhost:<PORT>`

- **Portal API Port**: `8080`
- **Admin API Port**: `8081`

---

## Portal API (`portal-api`)

All portal API endpoints are prefixed with `/api/v1`.

### Authentication

- **Public Routes**: Accessible by anyone.
- **Authenticated Routes**: Require a `Bearer Token` in the `Authorization` header.

### User Management (`/user`)

#### **`POST /user/register`**
- **Description**: Registers a new user.
- **Auth**: Public
- **Request Body**:
  ```json
  {
    "email": "user@example.com",
    "password": "password123"
  }
  ```
- **Success Response (201)**:
  ```json
  {
    "message": "User registered successfully",
    "user_id": 1
  }
  ```

#### **`POST /user/login`**
- **Description**: Logs in a user and returns a JWT.
- **Auth**: Public
- **Request Body**:
  ```json
  {
    "email": "user@example.com",
    "password": "password123"
  }
  ```
- **Success Response (200)**:
  ```json
  {
    "message": "Login successful",
    "token": "ey..."
  }
  ```

#### **`POST /user/request-password-reset`**
- **Description**: Initiates the password reset process for a user.
- **Auth**: Public
- **Request Body**:
  ```json
  {
    "email": "user@example.com"
  }
  ```
- **Success Response (200)**:
  ```json
  {
    "message": "If an account with that email exists, a password reset link has been sent."
  }
  ```

#### **`POST /user/reset-password`**
- **Description**: Resets the user's password using a valid token.
- **Auth**: Public
- **Request Body**:
  ```json
  {
    "token": "valid-reset-token",
    "new_password": "newStrongPassword123"
  }
  ```
- **Success Response (200)**:
  ```json
  {
    "message": "Password has been reset successfully."
  }
  ```

### Profile (`/profile`)

#### **`GET /profile/`**
- **Description**: Retrieves the authenticated user's profile information.
- **Auth**: Required
- **Success Response (200)**:
  ```json
  {
    "email": "user@example.com",
    "balance": 100.50,
    "plan_name": "Basic Plan",
    "plan_expiry_date": "2025-12-31",
    "used_traffic_gb": 50.5,
    "total_traffic_gb": 200.0
  }
  ```

#### **`GET /profile/history`**
- **Description**: Retrieves the user's consumption history (subscription orders).
- **Auth**: Required
- **Success Response (200)**:
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

### Store & Subscription (`/store`, `/subscription`)

#### **`GET /store/plans`**
- **Description**: Lists all available service plans.
- **Auth**: Public
- **Success Response (200)**:
  ```json
  [
    {
      "ID": 1,
      "Name": "Basic Plan",
      "Description": "A great starting plan.",
      "Price": 50.0,
      "TrafficGB": 200.0,
      "DurationDays": 30,
      "IsActive": true,
      "IsFeatured": true
    }
  ]
  ```

#### **`POST /store/purchase`**
- **Description**: Purchases a plan using the user's balance.
- **Auth**: Required
- **Request Body**:
  ```json
  {
    "plan_id": 1
  }
  ```
- **Success Response (201)**:
  ```json
  {
    "message": "Plan purchased successfully",
    "subscription": { ... }
  }
  ```

#### **`GET /subscription/`**
- **Description**: Retrieves the user's active subscription details.
- **Auth**: Required
- **Success Response (200)**:
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
- **Description**: Resets the user's subscription link.
- **Auth**: Required
- **Success Response (200)**:
  ```json
  {
    "message": "Subscription link reset successfully",
    "subscription": { ... }
  }
  ```

### Recharge (`/recharge`)

#### **`GET /store/recharge-presets`**
- **Description**: Lists all available recharge options.
- **Auth**: Public
- **Success Response (200)**:
  ```json
  [
    {
      "ID": 1,
      "AmountCNY": 50.0,
      "Description": "50 CNY Pack",
      "IsEnabled": true
    }
  ]
  ```

#### **`POST /recharge/create-order`**
- **Description**: Creates a new USDT recharge order.
- **Auth**: Required
- **Request Body**:
  ```json
  {
    "preset_id": 1
  }
  ```
- **Success Response (201)**:
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
- **Description**: Redeems a coupon code.
- **Auth**: Required
- **Request Body**:
  ```json
  {
    "code": "VALID-COUPON-CODE"
  }
  ```
- **Success Response (200)**:
  ```json
  {
    "message": "Coupon redeemed successfully"
  }
  ```

#### **`POST /recharge/generate-coupon`**
- **Description**: Generates a new coupon from the user's balance.
- **Auth**: Required
- **Request Body**:
  ```json
  {
    "value": 20.0
  }
  ```
- **Success Response (201)**:
  ```json
  {
    "ID": 1,
    "Code": "U-unique-uuid-string",
    "Value": 20.0,
    "Status": "active"
  }
  ```

### Other Public Endpoints

- **`GET /nodes/status`**: Lists the status of all server nodes.
- **`GET /announcements`**: Lists all active announcements.
- **`GET /help-documents`**: Lists all help documents.

---

## Admin API (`admin-api`)

All admin API endpoints are prefixed with `/api/v1/admin`. All endpoints (except `/login`) require a `Bearer Token` from an admin login.

### Authentication

#### **`POST /login`**
- **Description**: Logs in an administrator.
- **Request Body**:
  ```json
  {
    "username": "admin",
    "password": "adminpassword"
  }
  ```
- **Success Response (200)**:
  ```json
  {
    "token": "ey..."
  }
  ```

### Node Management (`/nodes`)

- **`POST /`**: Creates a new node.
- **`GET /`**: Lists all nodes.
- **`GET /:id`**: Gets a single node by ID.
- **`PUT /:id`**: Updates a node.
- **`DELETE /:id`**: Deletes a node.

### User Management (`/users`)

- **`GET /`**: Lists users with pagination and search (`?email=...&page=1&page_size=20`).
- **`PATCH /:id/status`**: Updates a user's status (e.g., `{"status": "locked"}`).

### Plan Management (`/plans`)

- **`POST /`**: Creates a new service plan.
- **`GET /`**: Lists all plans.
- **`GET /:id`**: Gets a single plan by ID.
- **`PUT /:id`**: Updates a plan.
- **`DELETE /:id`**: Deletes a plan.

### Order Management (`/orders`)

- **`GET /recharge`**: Lists recharge orders with pagination.
- **`GET /subscription`**: Lists subscription orders with pagination.

### Content Management (`/content`)

- **`POST /announcements`**: Creates a new announcement.
- **`PUT /announcements/:id`**: Updates an announcement.
- **`DELETE /announcements/:id`**: Deletes an announcement.
- **`POST /help-documents`**: Creates a new help document.
- **`PUT /help-documents/:id`**: Updates a help document.
- **`DELETE /help-documents/:id`**: Deletes a help document.

### Audit (`/audit`)

- **`GET /financial-overview`**: Gets a financial overview of the platform.
