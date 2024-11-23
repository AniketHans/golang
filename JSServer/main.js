import express from "express";

const app = express();
const PORT = 8000;

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.get("/", (req, res) => {
  res.status(200).json({ message: "Hello from JS Server" });
});

app.post("/post", (req, res) => {
  let val = req.body;
  res.status(200).send(val);
});

app.post("/postform", (req, res) => {
  res.json(req.body);
});

app.listen(PORT, () => {
  console.log(`Server listening at PORT ${PORT}`);
});
