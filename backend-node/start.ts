import createServer from "./server";
import connectDatabase from "./db";

async function start() {
    await connectDatabase()
    const app = await createServer();
    app.listen(4112, () =>
    console.log(`***** Server running at port ${4112} *****`)
    );
}

start();