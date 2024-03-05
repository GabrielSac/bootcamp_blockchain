use tendermint_abci::Application;

#[derive(Clone, Default)]
pub struct EchoApp;
impl EchoApp {
    pub fn new() -> Self {
        Self {}
    }
}

impl Application for EchoApp {}
