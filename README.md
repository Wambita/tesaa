# Tesaa FinBridge

**Tesaa FinBridge** is a blockchain-based financial ecosystem designed to bridge the gap between microfinance institutions (MFIs) and small business owners. By leveraging blockchain technology and the Safaricom API (M-Pesa), the platform ensures transparency, enhances transaction security, and expands access to capital for underserved entrepreneurs.

---

## Features

* **Blockchain Transparency:** Immutable ledger for recording loans and repayments, reducing fraud and building trust between lenders and borrowers.
* **M-Pesa Integration:** Seamless mobile money transactions via Safaricom API for instant loan disbursement and repayments.
* **User-Centric Design:** A highly accessible frontend designed for diverse users, from professional credit officers to small-scale traders.
* **Microfinance Connectivity:** Dedicated portals for MFIs to manage portfolios and for businesses to track their credit health.
* **Automated Verification:** Secure data handling to speed up the lending approval process.

---

## Tech Stack

* **Backend:** Go (Golang 1.22)
* **Blockchain:** Distributed Ledger Technology for secure transaction history.
* **Mobile Payments:** Safaricom API (M-Pesa).
* **Frontend:** Focused on usability and seamless UI/UX.
* **Mock Database:** JSON Server (for development/testing).
* **Containerization:** Docker.

---

## Getting Started

### Prerequisites

* **Docker** installed on your machine.
* **Go 1.22** (if running locally without Docker).
* **Node.js** (for the mock JSON server if running locally).

### Installation & Setup

1. **Clone the repository:**
```bash
git clone https://github.com/Wambita/tesaa.git
cd tesaa

```


2. **Build and Run with Docker:**
The project uses a multi-stage Dockerfile to orchestrate the Go backend and the JSON mock server.
```bash
# Build the image
docker build -t tesaa-finbridge .

# Run the container
docker run -p 8080:8080 -p 3000:3000 tesaa-finbridge

```


3. **Local Development (Manual):**
If you prefer to run the services separately:
* **Start the JSON Server:**
```bash
json-server --watch ./api/database/users.json --port 3000

```


* **Start the Go Application:**
```bash
go run .

```





---

## Project Structure

* `/api`: Contains database mocks and JSON schemas.
* `/cmd` or root: Go application entry points.
* `/frontend`: UI components and design assets.

---

## Contributing

Contributions make the open-source community an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git checkout origin feature/AmazingFeature`)
5. Open a Pull Request

---

## License

Distributed under the **MIT License**. See `LICENSE` for more information.

---

**Tesaa FinBridge** — Empowering small businesses through secure, transparent, and accessible financial technology.
