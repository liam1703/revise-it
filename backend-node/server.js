import express from "express";
import bodyParse from "body-parser";
// import cors from "cors";


async function createServer() {
    const app = express();
    app.use(bodyParse.json());
    return app
}

export default createServer;