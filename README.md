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

or for momentary stop so you don't need to create a new container

```bash
docker-compose stop
```

### Authentication (OAuth2)

---

### --- Sign In ---

- Method: `GET`
- Endpoint: `/oauth/login/google`
- Header:

  - Content-Type: `text/html`

- Response:

  - redirect -> `<FE_URL>?access_token=<access_token>&expires_in=<expires_in>&refresh_token=<refresh_token>`

- Error Response:
  - redirect -> `<FE_URL>?err=<err_msg>`

### --- Refresh Access Token ---

- Method: `POST`
- Endpoint: `/oauth/refresh/`
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
    "expires_in": 1676801048,
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
  - size: `20`
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
        {
          "image_id": "jkl",
          "url": "product_image"
        },
        {
          "image_id": "ijkl",
          "url": "image2"
        }
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
      "images": [
        {
          "image_id": "jkl",
          "url": "product_image"
        },
        {
          "image_id": "ijkl",
          "url": "image2"
        }
      ],
      "description": "Nama dada berbahan PVC Warna Hitam digravir dengan mesin Laser Uk. 2x8 cm pakai Peniti. Bisa juga pakai magnet atau paku",
      "min_order": 1,
      "created_at": "2023-01-18T12:03:56.595459Z",
      "updated_at": "2023-01-18T12:03:56.595459Z"
    }
  ],
  "meta": {
    "current_page": 1,
    "total_page": 10,
    "total_data": 200,
    "next_page": 2,
    "previous_page": 1
  }
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

### --- Product Detail ---

- Method: `GET`
- Endpoint: `/api/products/{product_id}`
- Header:

  - Content-Type: `application/json`

- Response:

```json
{
  "code": 200,
  "status": "OK",
  "error": null,
  "data": {
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
      {
        "image_id": "jkl",
        "url": "product_image"
      },
      {
        "image_id": "ijkl",
        "url": "image2"
      }
    ],
    "description": "Menggunakan kertas  NCR, ukuran 1/8 Folio 1 blok 1 Ply isi 100 lbr, 1 blok 2-4 Ply isi 50 set, Urutan warna pertama selalu putih",
    "min_order": 24,
    "created_at": "2023-01-18T12:03:56.595459Z",
    "updated_at": "2023-01-18T12:03:56.595459Z"
  }
}
```

- Error Response

```json
{
  "code": 400,
  "status": "BAD_REQUEST",
  "error": {
    "message": "record not found"
  },
  "data": null
}
```
