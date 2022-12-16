pub mod board;

pub struct GameEngine {
    pub board: board::Board,
}

pub enum MoveDirection { Up, Down, Left, Right }

impl GameEngine {
    pub fn new() -> GameEngine {
        GameEngine {
            board: board::Board::new(),
        }
    }
}

impl std::string::ToString for GameEngine {
    fn to_string(&self) -> String {
        let mut res = String::from("Game Engine {\n");

        let board_string = self.board.to_string();
        for line in board_string.lines() {
            res.push_str(&format!("\t{line}\n", line=line));
        }

        res.push_str("}");
        return res;
    }
}

impl std::string::ToString for MoveDirection {
    fn to_string(&self) -> String {
        match self {
            MoveDirection::Up => String::from("Up"),
            MoveDirection::Down => String::from("Down"),
            MoveDirection::Right => String::from("Right"),
            MoveDirection::Left => String::from("Left"),
        }
    }
}