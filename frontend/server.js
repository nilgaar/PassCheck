const express = require("express");
const path = require("path");
const PORT = process.env.PORT || 3000;
const RateLimit = require("express-rate-limit");

const app = express();

const limiter = new RateLimit({
  windowMs: 1 * 60 * 1000,
  max: 50,
});

app.use(limiter);

app.use(express.static(path.join(__dirname, "dist")));

app.get("*", (req, res) => {
  res.sendFile(path.resolve(__dirname, "dist", "index.html"));
});

app.listen(PORT, () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});
