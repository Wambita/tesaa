# Tesaa FinBridge
![Tesaa FinBridge Landing Page](/images/landing-page.png)


**Bridging Small Businesses and Microfinance Institutions Through Technology.**

Tesaa FinBridge is a financial technology application designed to connect small businesses with microfinance institutions (MFIs), making financial services more accessible. The platform explores blockchain integration and M-Pesa payments to support transparent, secure, and accessible microfinance services.

---

## Screenshots

Explore the platform's user interface and key features.

### 1. Landing Page

The landing page introduces Tesaa FinBridge and its mission to improve access to financial services for small businesses.

![Tesaa FinBridge Landing Page](/images/landing-page.png)

### 2. Business Onboarding

A dedicated onboarding experience for small businesses seeking access to financial services.

![Business Onboarding](/images/landing page-borrower.png)


![Business Onboarding](/images/business-onboarding.png)

### 3. Microfinance Institution Onboarding

An onboarding experience designed for microfinance institutions joining the platform.
![MFI Onboarding](/images/mfi-onboarding.png)

![MFI Onboarding](/images/landing page-lender.png)

### 4. Borrower Dashboard

A centralised dashboard where business users can access their account information and manage their borrowing activities.

![Borrower Dashboard](/images/borrower-dashboard.png)

![Borrower Profile](/images/borrower-profile.png)


### 5. Loan Application

An interface that allows businesses to submit loan applications through the platform.

![Loan Application](/images/loan-application.png)

### 6. Microfinance Institution Dashboard

A dedicated dashboard for microfinance institutions to manage their platform activities.

![MFI Dashboard](/images/mfi-dashboard.png)

![MFI Dashboard](/images/mfi-transactions.png)


---

## Technology Stack

- **Backend:** Go 1.22
- **Mock API:** JSON Server
- **Containerisation:** Docker
- **Frontend:** HTML, CSS, Javascript

---

## Prerequisites

Before running the project, ensure you have:

- Docker installed on your machine.
- Go 1.22 or later (for local development).
- Node.js and npm (for the mock JSON server).

---

## Installation & Setup

### 1. Clone the Repository

```bash
git clone https://github.com/Wambita/tesaa.git
cd tesaa
```

### 2. Build and Run with Docker

The project uses a multi-stage Dockerfile to run the Go backend and JSON mock server.

Build the Docker image:

```bash
docker build -t tesaa-finbridge .
```

Run the container:

```bash
docker run -p 8080:8080 -p 3000:3000 tesaa-finbridge
```

### 3. Local Development

Alternatively, you can run the services separately.

**Start the JSON Server:**

```bash
json-server --watch ./api/database/users.json --port 3000
```

**Start the Go Application:**

```bash
go run .
```

---

## Project Structure

```text
tesaa/
├── api/          # Mock database and JSON schemas
├── cmd/          # Application entry points (if applicable)
├── frontend/     # Frontend components and design assets
├── Dockerfile    # Docker configuration
└── README.md
```

---

## Contributing

Contributions are welcome! If you'd like to improve the project:

1. Fork the repository.
2. Create a feature branch:

   ```bash
   git checkout -b feature/AmazingFeature
   ```

3. Commit your changes:

   ```bash
   git commit -m "Add some AmazingFeature"
   ```

4. Push your branch:

   ```bash
   git push origin feature/AmazingFeature
   ```

5. Open a Pull Request.

---

## License

This project is distributed under the MIT License. See the `LICENSE` file for more information.

---

**Tesaa FinBridge** - Empowering small businesses through secure, transparent, and accessible financial technology.

