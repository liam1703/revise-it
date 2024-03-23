import createServer from "./server";
import connectDatabase from "./db";

async function start() {
    connectDatabase()
    const app = await createServer();
    app.listen(4112, () =>
    console.log(`***** Server running at port ${4112} *****`)
    );
}

start();