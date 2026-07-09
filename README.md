# 💖 Lovely Valentine

Lovely Valentine is a cute web application designed to send and receive special, password-protected messages for your loved ones. It's a sweet way to express your feelings digitally!

## ✨ Features

- **Password Protection:** Ensure only the right person can read your message.
- **Custom Messages:** Create and store personalized Valentine messages.
- **Cute UI:** Adorable design with various backgrounds and animations.
- **Easy Setup:** Powered by Go for fast and efficient performance.

## 🛠️ Tech Stack

- **Language:** [Go (Golang)](https://go.dev/)
- **Frontend:** HTML, CSS
- **Configuration:** `.env` for environment variables

## 🚀 Getting Started

### 1. Clone the Repository
```bash
git clone https://github.com/tudemaha/lovely-valentine.git
cd lovely-valentine
```

### 2. Environment Setup
Copy the example environment file and fill in your MongoDB credentials:
```bash
cp .env.example .env
```
Then, edit the `.env` file and provide the following details:
- `MONGO_USER`: Your MongoDB username
- `MONGO_PASS`: Your MongoDB password
- `MONGO_HOST`: Your MongoDB host/connection string
- `MONGO_DATABASE`: The name of your database
- `MONGO_COLLECTION`: The name of your collection

### 3. Run the Application
```bash
go run main.go
```
The application will be available at `http://localhost:3000`.

## 📂 Project Structure

```text
lovely-valentine/
├── controller/       # HTTP handlers and request logic
├── model/            # Data models and database operations
├── view/             # HTML templates for the UI
├── public/           # Static assets (images, CSS, favicon)
├── utils/            # Utility functions (e.g., hashing)
└── main.go           # Application entry point
```

## 📜 License

This project is licensed under the **MIT License**. See the [LICENSE](LICENSE) file for details.
