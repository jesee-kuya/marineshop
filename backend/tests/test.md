# API Tests

## Health Check

```bash
curl http://localhost:8080/health
```

---

## Auth

### Sign Up

```bash
curl -X POST http://localhost:8080/api/v1/auth/signup \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "testuser@example.com",
    "password": "password123"
  }'
```

### Login

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "testuser@example.com",
    "password": "password123"
  }'
```

### Change Password

Requires a valid JWT from login/signup in the `Authorization` header.

```bash
curl -X POST http://localhost:8080/api/v1/auth/change-password \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{
    "old_password": "password123",
    "new_password": "newpassword456"
  }'
```

### Reset Password

```bash
curl -X POST http://localhost:8080/api/v1/auth/reset-password \
  -H "Content-Type: application/json" \
  -d '{
    "email": "testuser@example.com",
    "new_password": "newpassword456"
  }'
```

---

## Seller

All seller endpoints require a valid JWT in the `Authorization` header.

### Collect KYC

Submits personal identity verification. Must be done before SetUpShop or SetUpPayment.

```bash
curl -X POST http://localhost:8080/api/v1/seller/kyc \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{
    "full_name": "John Doe",
    "phone_number": "+254712345678",
    "national_id": "12345678",
    "national_id_document": "https://storage.example.com/docs/national_id.jpg",
    "selfie": "https://storage.example.com/docs/selfie.jpg",
    "location": "Nairobi, Kenya"
  }'
```

**Errors**
- `409 Conflict` — KYC already submitted for this account.

### Set Up Shop

Submits business verification. Requires personal KYC to be submitted first.

```bash
curl -X POST http://localhost:8080/api/v1/seller/shop \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{
    "business_name": "Acme Marine Supplies",
    "document_type": "permit",
    "document": "https://storage.example.com/docs/business_permit.pdf"
  }'
```

`document_type` must be one of: `permit`, `certificate`, `incorporation_letter`.

**Errors**
- `422 Unprocessable Entity` — personal KYC not submitted yet.
- `409 Conflict` — business KYC already submitted.

### Set Up Payment — M-Pesa

Registers a payout account. Requires personal KYC to be submitted first.
A seller can register multiple accounts (one or more per wallet type).

```bash
curl -X POST http://localhost:8080/api/v1/seller/payment \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{
    "wallet_type": "mpesa",
    "account_name": "John Doe",
    "phone_number": "+254712345678",
    "is_default": true
  }'
```

### Set Up Payment — Airtel Money

```bash
curl -X POST http://localhost:8080/api/v1/seller/payment \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{
    "wallet_type": "airtel_money",
    "account_name": "John Doe",
    "phone_number": "+254733345678",
    "is_default": false
  }'
```

### Set Up Payment — Bank

```bash
curl -X POST http://localhost:8080/api/v1/seller/payment \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{
    "wallet_type": "bank",
    "account_name": "John Doe",
    "bank_name": "Equity Bank",
    "bank_code": "068",
    "account_number": "0123456789",
    "is_default": false
  }'
```

### Set Up Payment — Crypto

```bash
curl -X POST http://localhost:8080/api/v1/seller/payment \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{
    "wallet_type": "crypto",
    "account_name": "John Doe",
    "crypto_address": "0xAbCdEf1234567890AbCdEf1234567890AbCdEf12",
    "crypto_network": "ethereum",
    "is_default": false
  }'
```

`wallet_type` must be one of: `mpesa`, `airtel_money`, `bank`, `crypto`.

**Errors**
- `422 Unprocessable Entity` — personal KYC not submitted yet.
- `400 Bad Request` — missing required fields for the chosen wallet type.
