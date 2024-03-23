import dotenv from "dotenv";

dotenv.config();
var Mongoose = require("mongoose").Mongoose;

export const database = new Mongoose();


export default async function connectDatabase() {
  const DbUrl = process.env.DATABASE_URL;
  await database.connect(DbUrl, {
    useNewUrlParser: true,
    useUnifiedTopology: true,
    useFindAndModify: false,
  });
  console.log("***** Database connected *****");
}
