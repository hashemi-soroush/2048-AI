pub mod board;

use crate::player::Player;

pub struct GameEngine<P: Player> {
    board: board::Board,
    player: P,
    score: u32,
}

pub enum MoveDirection { Up, Down, Left, Right }

impl<P: Player> GameEngine<P> {
    pub fn new(size: u32, player: P) -> GameEngine<P> {
        GameEngine {
            board: board::Board::new(size),
            player: player,
            score: 0,
        }
    }

    pub fn play(&mut self) -> u32 {
        while self.board.can_move() {
            let player_move = self.player.play(&self.board);
            self.score += self.board.move_(&player_move);
            self.board.put_randomly(1, &[2, 4]);
        }
        return self.score;
    }
}

impl<P: Player> std::string::ToString for GameEngine<P> {
    fn to_string(&self) -> String {
        let mut board_str = self.board.to_string()
            .lines()
            .map(|line| format!("\t{}\n", line))
            .reduce(|cur, line| cur + &line);
        
        match board_str {
            Some(board_str) => "Game Engine {\n".to_string() + &board_str + "}",
            None => panic!("panic"),
        }
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