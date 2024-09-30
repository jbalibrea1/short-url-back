# API Backends in Different Languages

This is a simple API for shortening URLs, demonstrating the implementation in both GO and Express (Node.js). URL shortening is a convenient way to condense long URLs into manageable links.

## Table of Contents

- [Technologies Used](#technologies-used)
- [Endpoints](#endpoints)

## Technologies Used

- Express - Node.js
- GO
- MongoDB Atlas for cloud-hosted database
  To run the projects, create a `.env` file in each project directory with the MongoDB URI configured as shown in `.env.example`.

[GO](back-go/README.md)

[Express - Node.js](api-express/README.md)

## Endpoints

<details>
<summary>GET /shorturl | to get all shortened URLs</summary>

```json
[
  {
    "id": "<id>",
    "url": "<original-url>",
    "title": "<title>",
    "logo": "<logo>",
    "description": "<description>",
    "shortURL": "<short-url>",
    "totalClicks": <total-clicks>,
    "createdAt": "<created-at>"
  },
  ...
]
```

</details>

<details>
<summary>POST /shorturl | to shorten a URL</summary>

```json
- Request
{
  "url": "<original-url>"
}

- Response
{
  "id": "<id>",
  "url": "<original-url>",
  "title": "<title>",
  "logo": "<logo>",
  "description": "<description>",
  "shortURL": "<short-url>",
  "totalClicks": 0,
  "createdAt": "<created-at>"
}
```

</details>

<details>
<summary>GET /shorturl/:shorturl | to get info from a specific shorten url</summary>

```json
  {
    "id": "<id>",
    "url": "<original-url>",
    "title": "<title>",
    "logo": "<logo>",
    "description": "<description>",
    "shortURL": "<short-url>",
    "totalClicks": <total-clicks>,
    "createdAt": "<created-at>"
  }
```

</details>
