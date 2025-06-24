# NucleoOAuth

A minimal OAuth2 implementation with a clear separation between the **Authorization Server** and the **Resource Server**. Built for educational purposes and practical integration scenarios.

## 📚 Overview

**NucleoOAuth** demonstrates how to structure an OAuth2-based authentication flow with clean boundaries between:

- **Authorization Server**: Handles user authentication, authorization, token issuance, and consent.
- **Resource Server**: Exposes protected APIs and verifies access tokens.
- **Client** (optional example): A sample consumer application that initiates the OAuth flow.

---

## 🏗️ Architecture

```
+-----------+                                          +--------------------+
|           |--(1) Authorization Request-------------->| Authorization      |
|           |                                          |     Server         |
|           |<--(2) Authorization Code-----------------| (Login & Consent)  |
|           |                                          +--------------------+
|           |
|           |--(3) Token Request----------------------->|
|   Client  |                                           |
|           |<--(4) Access Token------------------------|
|           |
|           |--(5) API Request with Token------------->| Resource Server    |
|           |                                           | (Validates Token)  |
|           |<--(6) Protected Resource------------------|                    |
+-----------+                                          +--------------------+
```

---

## 🚀 Features

- ✅ OAuth2 Authorization Code Flow
- ✅ Token Introspection
- ✅ Modular and Extensible
- ✅ Clear server separation for real-world scenarios

---

## 📁 Repository Structure

```
NucleoOAuth/
│
├── server/
|    ├── resource/      # Resource Server
|    │   ├── main.go
|    │   └── ...
|    │
|    ├── authorization/      # Authorization Server
|        ├── main.go
│        └── ...
│          
├── client/       # client svelte app
│   └── ...
│
└── README.md             # You're here
```

---

## ⚙️ Setup & Running

### Prerequisites

- Go >= 1.20
- Redis (for token storage or session state)
- Environment Variables setup (see `.env.example` if available)

## 🔐 OAuth Flow

1. Client redirects user to Authorization Server
2. User authenticates and gives consent
3. Client receives auth code
4. Client exchanges code for token
5. Client accesses protected resource with token
6. Resource Server validates token with Authorization Server (or locally)

---

## 🌐 Endpoints

### Authorization Server

- `GET /authorize` – Initiates login + consent
- `POST /login/{id}` – Login Route
- `POST /register/{id}` – Register Route
- `POST /introspect` – Token validation
- `GET /newAccess` – Gives new Access token and refresh's refresh token
  
  ### Resource Server

- `GET /api/acess` – Protected resource requiring valid token

---


## 📄 License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

---

## 📬 Contact

Made with ❤️ by [@NucleoFusion](https://github.com/NucleoFusion)
