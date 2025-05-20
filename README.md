# 📘 Microservices Summary: Customer Authentication System

This microservice is designed to handle **user account management** and **authentication** for customers using **MongoDB** for data storage and **JWT (JSON Web Tokens)** for secure session handling. The service is built in Go and structured in a modular fashion for easy scalability and maintenance.

---

### 🔑 Core Features

1. **Create Customer Account**
    - Registers a new user with email uniqueness validation.

2. **Customer Login**
    - Authenticates a user using email/password and returns a JWT token for session tracking.

3. **Reset Password (Authenticated)**
    - Allows users to update their password using the old password (requires valid token).

4. **Generate Token Using Email**
    - Generates a JWT token by verifying the customer's email.

5. **Update Password via Token**
    - Enables password update using a valid JWT token (email is extracted from the token).

---

### ⚙️ Technologies Used

- **Language:** Go
- **Database:** MongoDB
- **Authentication:** JWT
- **Environment Management:** `.env` file

---

### 📄 Customer Data Schema

The customer document includes fields like `customer_id`, `first_name`, `last_name`, `email`, `password`, `phone`, `DOB`, `gender`, and timestamps (`created_at`, `updated_at`).

---

### 🔌 API Design Highlights

- All endpoints use `POST` method.
- Input validation and proper error handling are implemented.
- JWT-based authorization ensures secure access.
- Structured and consistent API responses (`success`, `message`, `data`).

---

# Customer Authentication Microservice

This microservice handles customer authentication and password management using **MongoDB** and **JWT** in a modular Go application.

## 🔧 Features

1. Create Customer Account
2. Customer Login (JWT Token generation)
3. Reset Password (requires old password)
4. Create JWT Token using Email
5. Update Password via JWT Token

## 🗃️ Tech Stack

- **Database:** MongoDB
- **Authentication:** JWT
- **Environment Config:** `.env` file

## 📂 Customer Collection Structure

```json
{
  "_id": { "$oid": "682894dba7245b81c235dd91" },
  "customer_id": "2104356",
  "customer_first_name": "Amanda",
  "customer_last_name": "Harris",
  "customer_email": "amanda.harris77@example.com",
  "customer_password": "This is password string",
  "customer_phone": "685-588-4720",
  "customer_dob": "1968-12-04T13:40:26",
  "customer_gender": "F",
  "created_at": "2025-05-17T13:40:26",
  "updated_at": "2025-05-19T02:47:48.503Z"
}
```

## 📁 Environment Variables (`.env`)

```env
# JWT Config
JWT_SECRET=JWT_SECRET
JWT_EXPIRY_HOURS=360
JWT_ALGORITHM=HS256

# MongoDB Config
MONGO_URI=MONGODB_CONNECTION_STRING
MONGO_DB=MYSHOP_DB
```

## 📌 API Endpoints

### 1. 🔐 Create Customer Account

- **URL:** `http://localhost:8080/api/createaccount`
- **Method:** `POST`
- **Payload:**

```json
{
  "customer_first_name": "Jessica2",
  "customer_last_name": "Jackson2",
  "customer_email": "jessica.jackson762@example.com",
  "customer_password": "P@ssw0rd!",
  "customer_phone": "984-513-6972",
  "customer_dob": "1978-09-10",
  "customer_gender": "F"
}
```

- **Success Response:**

```json
{
  "data": {
    "customer_id": "1747572480496770990"
  },
  "message": "Account created successfully",
  "success": true
}
```

- **Error (Email Exists):**

```json
{
  "message": "Failed to create account: email already exists",
  "success": false
}
```

---

### 2. 🔑 Customer Login

- **URL:** `http://localhost:8080/api/customerlogin`
- **Method:** `POST`
- **Payload:**

```json
{
  "email": "amanda.harris77@example.com",
  "password": "Temp@123"
}
```

- **Success Response:**

```json
{
  "data": {
    "token": "THIS IS A JWT TOKEN"
  },
  "message": "Login successful",
  "success": true
}
```

- **Failure Response:**

```json
{
  "message": "Login failed: invalid email / password",
  "success": false
}
```

---

### 3. 🔁 Reset Password

- **URL:** `http://localhost:8080/api/resetpassword`
- **Method:** `POST`
- **Payload:**

```json
{
  "email": "amanda.harris77@example.com",
  "old_password": "Test@123",
  "new_password": "Temp@123"
}
```

- **Success Response:**

```json
{
  "data": null,
  "message": "Password updated successfully",
  "success": true
}
```

- **Error Response:**

```json
{
  "message": "Failed to reset password: invalid old password",
  "success": false
}
```

---

### 4. 🛠️ Create Token Using Email

- **URL:** `http://localhost:8080/api/createtoken`
- **Method:** `POST`
- **Payload:**

```json
{
  "email": "amanda.harris77@example.com"
}
```

- **Response:**

```json
{
  "data": {
    "token": "THIS IS A JWT TOKEN"
  },
  "message": "Token generated successfully",
  "success": true
}
```

---

### 5. 🔄 Update Password via Token

- **URL:** `http://localhost:8080/api/updatepwd`
- **Method:** `POST`
- **Payload:**

```json
{
  "token": "THIS IS A JWT TOKEN",
  "new_password": "Test@123"
}
```

- **Success Response:**

```json
{
  "data": null,
  "message": "Password has been reset successfully",
  "success": true
}
```

- **Error Response:**

```json
{
  "message": "Invalid or expired token",
  "success": false
}
```

---

## 🚀 Getting Started

1. Set up your MongoDB database and collection (`customerinfo`) based on the sample schema above.
2. Create a `.env` file in the root of your project using the structure provided.
3. Run the Go service.
4. Test the endpoints using tools like Postman or curl.

---

## 📬 Contact

For issues or contributions, feel free to open an issue or submit a pull request.