# User Day Randomizer

A Go project that generates a random schedule of users across weekdays and presents the results in a structured format.

This project is designed to demonstrate:
- Random data generation
- Business rule implementation
- Clean and readable code structure
- Ordered and sorted output

---

## ✨ Features

- Randomly assigns users to weekdays. (Mon–Fri)
- Dynamic capacity per day based on the number of users.
- Days are always displayed in fixed order. (Mon → Fri)
- Implement error handling when user count exceeds the allowed capacity.

---

## 🧠 Business Rules

1. If the number of users is less than or equal to the number of days, each user is assigned a unique day.
2. If the number of users exceeds the number of days, multiple users may share the same day.
3. The maximum number of users per day is calculated dynamically based on the formula.

---

## 🚀 How to Run

### Option 1: Run with Go
```bash
go run main.go
```

### Option 2: Run with Docker
Build the Docker image
```bash
docker build -t user-day-randomizer:dev .
```
Run the container
```bash
docker run -d --name user-day-randomizer -v $(pwd):/app user-day-randomizer:dev
```

---

## 📊 Example Output
```bash
User Schedule
-------------
Mon: Alice, Ivy, Oscar
Tue: Bob, Emma, Mike
Wed: Charlie, Frank, Paul
Thu: David, Grace, Lily
Fri: Henry, Jack, Kevin, Nina
```

---

## 📄 License
MIT License
