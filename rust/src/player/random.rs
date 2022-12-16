use rand::Rng;
use crate::game_engine::board::Board;
use crate::game_engine::MoveDirection;

pub struct RandomPlayer {}

impl super::Player for RandomPlayer {
	fn play(&self, _board: &Board) -> MoveDirection {
		let mut rng = rand::thread_rng();
		match rng.gen::<usize>() % 4 {
			0 => MoveDirection::Up,
			1 => MoveDirection::Down,
			2 => MoveDirection::Right,
			3 => MoveDirection::Left,
			_ => panic!("it shouldn't happen"),
		}
	}
}