mod echo;

use echo::EchoApp;
use tendermint_abci::ServerBuilder;

fn main() {
    let read_buf_size = 1048576;
    let host = "localhost";
    let port = 26658;

    let addr = format!("{}:{}", host, port);

    let app = EchoApp::new();
    let server = ServerBuilder::new(read_buf_size).bind(addr, app).unwrap();

    server.listen().unwrap();
}
