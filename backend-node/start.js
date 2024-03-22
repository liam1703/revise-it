import createServer from "./server";

async function start() {

    const app = await createServer();
    app.listen(4112, () =>
    console.log(`***** Server running at port ${port} *****`)
    );
}

start();