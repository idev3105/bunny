# Bunny

### A Template Starter for Go Web Applications

---

## Introduction

As a developer, you understand the challenges of setting up a new project with all the necessary tools and a clean, tidy architecture. That's why Bunny was born.

---

## Features

- **Web Framework**: [Echo](https://echo.labstack.com/)
- **SQL Client**: [sqlc](https://github.com/sqlc-dev/sqlc)
- **Build Tool**: [GNU Make](https://www.gnu.org/software/make/manual/make.html#Rule-Introduction)
- **Databases**: MySQL, PostgreSQL, SQLite, MongoDB
- **Caching**: [Redis](https://redis.io/)
- **Migration**: [Go-Migrate](https://github.com/golang-migrate/migrate)
- **Pub/Sub**: [Kafka](https://kafka.apache.org/)
- **Live Reload**: [Air](https://github.com/cosmtrek/air)

---

## Getting Started

1. **Clone the repository**:

   ```sh
   git clone https://github.com/yourusername/bunny.git
   cd bunny
   ```

2. **Install dependencies**:

   ```sh
   make install
   ```

3. **Run the server**:

   ```sh
   make serve-dev
   ```

4. **Run the consumer**:

   ```sh
   make consumer name=example
   ```

---

## Contributing

We welcome contributions! Please see our [contributing guidelines](CONTRIBUTING.md) for more details.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Contact

For any inquiries, please contact [yourname@domain.com](mailto:yourname@domain.com).

---

## Acknowledgements

Special thanks to all the contributors and the open-source community for their invaluable support.
