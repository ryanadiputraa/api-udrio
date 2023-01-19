# UD Rio Backend

Backend service for UD Rio (prototype) website.

## Get Started

---

You need to have `docker-compose` to simplify development, config your `.env` file and run the following command to setup your local db & run dev server

```bash
docker-compose up -d
```

it has `Hot Reload` setup, so you don't need to restart server to make changes in development.
To stop dev server simply run

```bash
docker-compose down
```

### Authentication

---

### --- Sign Up ---

- Method: `POST`
- Endpoint: `/auth/register/`
- Header:
  - Content-Type: `application/json`
- Body:

```json
{
  "username": "username",
  "email": "user@mai.com",
  "password": "password"
}
```

- Response:

```json
{
  "code": 201,
  "status": "CREATED",
  "error": null,
  "data": null
}
```

- Error Response

```json
{
  "code": 400,
  "status": "BAD_REQUEST",
  "error": {
    "message": "invalid param",
    "username": ["must not be null", "must be greater than 3 character"],
    "email": ["invalid email format"],
    "passowrd": ["must not be null", "must be greater than 3 character"]
  },
  "data": null
}
```

### --- Sign In ---

- Method: `POST`
- Endpoint: `/auth/login/`
- Header:
  - Content-Type: `application/json`
- Body:

```json
{
  "email": "user@mai.com",
  "password": "password"
}
```

- Response:

```json
{
  "code": 200,
  "status": "OK",
  "error": null,
  "data": {
    "access_token": "MTQ0NjJkZmQ5OTM2NDE1ZTZjNGZmZjI3",
    "expires_in": 3600,
    "refresh_token": "IwOGYzYTlmM2YxOTQ5MGE3YmNmMDFkNTVk"
  }
}
```

- Error Response

```json
{
  "code": 400,
  "status": "BAD_REQUEST",
  "error": {
    "message": "invalid password"
  },
  "data": null
}
```

### --- Refresh Access Token ---

- Method: `POST`
- Endpoint: `/auth/refresh/`
- Header:

  - Content-Type: `application/json`
  - Authorization: `Bearer <refresh_token>`

- Response:

```json
{
  "code": 200,
  "status": "OK",
  "error": null,
  "data": {
    "access_token": "MTQ0NjJkZmQ5OTM2NDE1ZTZjNGZmZjI3",
    "expires_in": 3600,
    "refresh_token": "IwOGYzYTlmM2YxOTQ5MGE3YmNmMDFkNTVk"
  }
}
```

- Error Response

```json
{
  "code": 400,
  "status": "BAD_REQUEST",
  "error": {
    "message": "invalid refresh token header"
  },
  "data": null
}
```

### Users

---

### --- Change password ---

- Method: `PUT`
- Endpoint: `/api/users/`
- Header:
  - Content-Type: `application/json`
  - Authorization: `Bearer <access_token>`
- Body:

```json
{
  "password": "password",
  "new_password": "new_password"
}
```

- Response:

```json
{
  "code": 200,
  "status": "OK",
  "error": null,
  "data": null
}
```

- Error Response

```json
{
  "code": 400,
  "status": "BAD_REQUEST",
  "error": {
    "message": "invalid param",
    "password": "invalid old password"
  },
  "data": null
}
```

### Products

---

### --- List of categories ---

- Method: `GET`
- Endpoint: `/api/categories/`
- Header:

  - Content-Type: `application/json`

- Response:

```json
{
  "code": 200,
  "status": "OK",
  "error": null,
  "data": [
    {
      "id": 1,
      "category": "Spanduk & Banner",
      "icon": "https://domain.com/ico.png"
    },
    {
      "id": 2,
      "category": "Cetak Kalender & Kartu Nama",
      "icon": "https://domain.com/ico.png"
    },
    {
      "id": 3,
      "category": "Undangan, Karcis/Tiket & Foto",
      "icon": "https://domain.com/ico.png"
    },
    {
      "id": 4,
      "category": "Perkantoran",
      "icon": "https://domain.com/ico.png"
    }
  ]
}
```

- Error Response

```json
{
  "code": 404,
  "status": "NOT_FOUND",
  "error": {
    "message": "categories empty"
  },
  "data": null
}
```

### --- List of products ---

- Method: `GET`
- Endpoint: `/api/products/`
- Params:
  - page: `1`
  - category_id: `1` (optional)
- Header:

  - Content-Type: `application/json`

- Response:

```json
{
  "code": 200,
  "status": "OK",
  "error": null,
  "data": [
    {
      "id": "MTQ0NjJkZmQ5OTM2NDE1ZTZjNGZmZjI3",
      "product_name": "Kop Surat",
      "product_category": {
        "category_id": 1,
        "category": "ATK",
        "icon": "https://domain.com/image.png"
      },
      "price": 7500,
      "available": true,
      "images": [
        "https://domain.com/image.png",
        "https://domain.com/image2.png"
      ],
      "description": "Menggunakan kertas  NCR, ukuran 1/8 Folio 1 blok 1 Ply isi 100 lbr, 1 blok 2-4 Ply isi 50 set, Urutan warna pertama selalu putih",
      "min_order": 24,
      "created_at": "2023-01-18T12:03:56.595459Z",
      "updated_at": "2023-01-18T12:03:56.595459Z"
    },
    {
      "id": "MTQ0NjJkZmQ5OTM2NDE1ZTZjNGZmZjI6",
      "product_name": "Nama Dada / Namepin Gravir Hitam Standar Uk. 2x8cm",
      "product_category": {
        "category_id": 2,
        "category": "Aksesoris",
        "icon": "https://domain.com/image.png"
      },
      "price": 30000,
      "available": false,
      "images": ["https://domain.com/image.png"],
      "description": "Nama dada berbahan PVC Warna Hitam digravir dengan mesin Laser Uk. 2x8 cm pakai Peniti. Bisa juga pakai magnet atau paku",
      "min_order": 1,
      "created_at": "2023-01-18T12:03:56.595459Z",
      "updated_at": "2023-01-18T12:03:56.595459Z"
    }
  ]
}
```

- Error Response

```json
{
  "code": 400,
  "status": "BAD_REQUEST",
  "error": {
    "message": "invalid param",
    "category": "invaild category"
  },
  "data": null
}
```
