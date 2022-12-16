pub mod random;

use crate::game_engine::MoveDirection;
use crate::game_engine::board::Board;

pub trait Player {
    fn play(&self, board: &Board) -> MoveDirection;
}