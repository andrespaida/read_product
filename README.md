# Read Product Microservice

This microservice retrieves product information from the MongoDB database. It is part of the ToyShop platform and allows users to view all available products.

## Technologies Used

- Go (Golang)
- Gin (web framework)
- MongoDB
- Docker
- GitHub Actions

## Getting Started

### Prerequisites

- Go >= 1.18
- MongoDB
- Git

### Installation

```bash
git clone https://github.com/andrespaida/read_product.git
cd read_product
go mod tidy
```

### Environment Variables

Create a `.env` file in the root directory with the following content:

```env
PORT=4001
MONGO_URI=mongodb://your_mongo_host:27017
DB_NAME=toyshop_db
COLLECTION_NAME=products
```

### Running the Service

```bash
go run main.go
```

The service will be running at `http://localhost:4001`.

## Available Endpoint

### GET `/products`

Retrieves a list of all products in the database.

#### Example Response:

```json
[
  {
    "id": "60f5cbb2a2e3f0a001d4a7d9",
    "name": "Toy Car",
    "description": "Red racing toy car",
    "price": 19.99,
    "stock": 50,
    "category": "Vehicles",
    "image_url": "http://localhost:4002/uploads/toycar.jpg"
  }
]
```

## Docker

To build and run the service using Docker:

```bash
docker build -t read-product .
docker run -p 4001:4001 --env-file .env read-product
```

## GitHub Actions Deployment

This project includes a GitHub Actions workflow for automatic deployment to an EC2 instance. Configure the following secrets in your GitHub repository:

- `EC2_HOST`
- `EC2_USERNAME`
- `EC2_KEY`
- `EC2_PORT`

## License

This project is licensed under the MIT License.