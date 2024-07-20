const express = require("express");
const path = require("path");
const PORT = process.env.PORT || 3000;
let RateLimit = require("express-rate-limit");

const limiter = new RateLimit({
  windowMs: 1 * 60 * 1000,
  max: 50,
});

express().use(express.static(path.join(__dirname, "dist")), limiter);

express().get("*", (req, res) => {
  res.sendFile(path.resolve(__dirname, "dist", "index.html"));
});

express().listen(PORT, () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});
