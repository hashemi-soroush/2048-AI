pub mod board;

#[derive(Debug)]
pub struct GameEngine {
    board: board::Board,
}

impl GameEngine {
    pub fn new() -> GameEngine {
        let game_engine = GameEngine {
            board: board::Board::new(),
        };
        return game_engine
    }
}

impl std::string::ToString for GameEngine {
    fn to_string(&self) -> String {
        let mut res = String::from("Game Engine {\n");

        let board_string = self.board.to_string();
        for line in board_string.lines() {
            res.push_str("\t");
            res.push_str(line);
            res.push_str("\n");
        }

        res.push_str("}");
        return res;
    }
}