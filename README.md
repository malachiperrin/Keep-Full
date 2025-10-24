# 🍽️ Inventory Management API

This API powers the inventory system for deli and quick food spots. Below are the available endpoints and data structures.

---

## 📦 Product Endpoints

### `GET /products`
Fetches all products in the inventory.
**Returns:** Array of `Product` objects.

### `GET /products/{id}`
Fetches a single product by its ID.
**Returns:** A `Product` object.

---

## 🏪 Vendor Endpoints

### `GET /vendors`
Fetches all vendors supplying products.
**Returns:** Array of `Vendor` objects.

---

## 🧬 Data Types

### Product

| Field                | Type                | Description                    |
|---------------------|---------------------|--------------------------------|
| `ProductID`          | `int`               | Unique identifier              |
| `ProductName`        | `string`            | Name of the product            |
| `ProductDependencies`| `[]string (nullable)` | Optional dependencies         |
| `VendorID`           | `int`               | Linked vendor ID               |

### Vendor

| Field           | Type     | Description           |
|----------------|----------|-----------------------|
| `VendorID`      | `int`    | Unique identifier     |
| `VendorName`    | `string` | Name of the vendor    |
| `VendorContact` | `string` | Contact information   |

---

_Built with ❤️ in Go. Last updated: October 2025_
