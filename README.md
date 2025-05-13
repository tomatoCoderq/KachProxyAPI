# KachProxyAPI

A lightweight proxy API for the official [teatrkachalov.ru](https://teatrkachalov.ru/) website. This service scrapes and exposes public data about plays, actors, and musicians in a clean, JSON-based format.

---

## Features

- Get all upcoming **plays** with filtering options
- Fetch details about a **specific play**
- Get a list of all **actors**
- Get details for a **specific actor**
- Get a list of all **musicians**

---

## API Endpoints

### Plays

#### `GET /plays`

Returns a list of upcoming plays with optional filters.

**Query Parameters**:

| Name        | Type   | Description                                        |
|-------------|--------|----------------------------------------------------|
| `scene`     | string | Filter by season (`Основная сцена`, `Малая сцена`) |
| `author`    | string | Filter by surname (Case sensetive)                 |
| `name`      | string | Filter by name of the play                         |
| `age_rating`| string | `6`, `12`, `16`, or `18` (+ omitted)               |
| `month`     | string | Filter by month (from 1 to 12)                     |

#### `GET /plays/{id}`

Returns a JSON object with details about a specific play.

---

### 🎭 Actors

#### `GET /actors`

Returns a list of all actors from the theater.

#### `GET /actors/{id}`

Returns details for a specific actor.

---

### 🎵 Musicians

#### `GET /musicians`

Returns a list of all musicians from the theater.

---

## Starting project

### 🔧 Prerequisites

- Go 1.20 or higher

### 🛠 Install & Run

1. **Clone the repository:**

```bash
git clone https://github.com/your-username/KachProxyAPI.git
cd KachProxyAPI
```

2. **Add all required files**
```bash
touch theat.toml
```

