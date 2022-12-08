use super::MoveDirection;
use rand::Rng;


#[derive(Debug)]
pub struct Board([[u32; 4]; 4]);

impl Board {
    pub fn new() -> Self {
        let mut board = Board{0: [[0u32;4];4]};
        
        let mut rng = rand::thread_rng();
        let mut seed_count = 2;
        while seed_count > 0 {
            let mut i: usize = rng.gen(); i = i % board.0.len();
            let mut j: usize = rng.gen(); j = j % board.0[0].len();
            if board.0[i][j] > 0 {
                continue;
            }

            let mut val: u32 = rng.gen(); val = val % 2 + 1;
            board.0[i][j] = val;
            seed_count -= 1;
        }

        return board;
    }

    pub fn move_(&mut self, direction: &MoveDirection) -> u32 {
        match direction {
            MoveDirection::Up => self.move_up(),
            MoveDirection::Down => self.move_down(),
            MoveDirection::Left => self.move_left(),
            MoveDirection::Right => self.move_right(),
        }
    }

    fn move_up(&mut self) -> u32 {
        let mut score: u32 = 0;
        for j in 0..self.0.len() {
            let mut cur = 0;
            for i in 0..self.0.len() {
                if i == cur || self.0[i][j] == 0 {
                    continue;
                }

                if self.0[cur][j] == 0 {
                    self.0[cur][j] = self.0[i][j];
                    self.0[i][j] = 0;
                } else if self.0[cur][j] == self.0[i][j] {
                    self.0[cur][j] *= 2;
                    self.0[i][j] = 0;
                    score += self.0[cur][j];
                    cur += 1;

                } else {
                    cur += 1;
                    let temp = self.0[i][j];
                    self.0[i][j] = 0;
                    self.0[cur][j] = temp;
                }
            }
        }
        return score;
    }

    fn move_down(&mut self) -> u32 {
        let mut score: u32 = 0;
        for j in 0..self.0[0].len() {
            let mut cur = self.0.len()-1;
            for i in (0..self.0.len()).rev() {
                if cur == i || self.0[i][j] == 0 {
                    continue;
                }

                if self.0[cur][j] == 0 {
                    self.0[cur][j] = self.0[i][j];
                    self.0[i][j] = 0;
                } else if self.0[i][j] == self.0[cur][j] {
                    self.0[cur][j] *= 2;
                    self.0[i][j] = 0;
                    score += self.0[cur][j];
                    cur -= 1;
                } else {
                    cur -= 1;
                    let temp = self.0[i][j];
                    self.0[i][j] = 0;
                    self.0[cur][j] = temp;
                }
            }
        }
        return score;
    }

    fn move_left(&mut self) -> u32 {
        let mut score: u32 = 0;
        for i in 0..self.0.len() {
            let mut cur = 0;
            for j in 0..self.0[i].len() {
                if self.0[i][j] == 0 || cur == j {
                    continue;
                }

                if self.0[i][cur] == 0 {
                    self.0[i][cur] = self.0[i][j];
                    self.0[i][j] = 0;
                } else if self.0[i][j] == self.0[i][cur] {
                    self.0[i][cur] *= 2;
                    self.0[i][j] = 0;
                    score += self.0[i][cur];
                    cur += 1;
                } else {
                    cur += 1;
                    let temp = self.0[i][j];
                    self.0[i][j] = 0;
                    self.0[i][cur] = temp;
                }
            }
        }
        return score;
    }
    
    fn move_right(&mut self) -> u32 {
        let mut score = 0;
        for i in 0..self.0.len() {
            let mut cur = self.0[0].len() - 1;
            for j in (0..self.0[0].len()).rev() {
                if self.0[i][j] == 0 || cur == j {
                    continue;
                }

                if self.0[i][cur] == 0 {
                    self.0[i][cur] = self.0[i][j];
                    self.0[i][j] = 0;
                } else if self.0[i][j] == self.0[i][cur] {
                    self.0[i][cur] *= 2;
                    self.0[i][j] = 0;
                    score += self.0[i][cur];
                    cur -= 1;
                } else {
                    cur -= 1;
                    let temp = self.0[i][j];
                    self.0[i][j] = 0;
                    self.0[i][cur] = temp;
                }
            }
        }
        return score;
    }
}

impl std::string::ToString for Board {
    fn to_string(&self) -> String {
        let mut res = String::from("Board {\n");
        for i in 0..self.0.len() {
            res += "\t";
            for j in 0..self.0[i].len() {
                res.push_str("\t");
                res.push_str(&self.0[i][j].to_string());
            }
            res.push_str("\n");
        }
        res.push_str("}");
        return res;
    }
}

impl std::cmp::PartialEq for Board {
    fn eq(&self, other: &Self) -> bool {
        if self.0.len() != other.0.len() { return false; }
        for i in 0..self.0.len() {
            if self.0[i].len() != other.0[i].len() { return false; }

            for j in 0..self.0[i].len() {
                if self.0[i][j] != other.0[i][j] { return false; }
            }
        }
        return true;
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn board_moves() {
        struct MoveSample {
            start: super::Board,
            end: super::Board,
            direction: super::MoveDirection
        }

        let mut samples = [
            MoveSample{
                start:  super::Board{0: [[0, 2, 2, 2], [0, 0, 2, 0], [0, 0, 0, 2], [0, 0, 0, 0]]},
                end:    super::Board{0: [[0, 2, 4, 4], [0, 0, 0, 0], [0, 0, 0, 0], [0, 0, 0, 0]]},
                direction: super::MoveDirection::Up,
            },
            MoveSample{
                start:  super::Board{0: [[0, 2, 0, 0], [0, 0, 2, 2], [2, 0, 2, 2], [0, 2, 0, 0]]},
                end:    super::Board{0: [[2, 4, 4, 4], [0, 0, 0, 0], [0, 0, 0, 0], [0, 0, 0, 0]]},
                direction: super::MoveDirection::Up,
            },
            MoveSample{
                start:  super::Board{0: [[2, 0, 0, 4], [2, 2, 4, 0], [2, 2, 2, 2], [2, 4, 2, 2]]},
                end:    super::Board{0: [[4, 4, 4, 4], [4, 4, 4, 4], [0, 0, 0, 0], [0, 0, 0, 0]]},
                direction: super::MoveDirection::Up,
            },
            MoveSample{
                start:  super::Board{0: [[0, 2, 2, 2], [0, 0, 2, 0], [0, 0, 0, 2], [0, 0, 0, 0]]},
                end:    super::Board{0: [[0, 0, 0, 0], [0, 0, 0, 0], [0, 0, 0, 0], [0, 2, 4, 4]]},
                direction: super::MoveDirection::Down,
            },
            MoveSample{
                start:  super::Board{0: [[0, 2, 0, 0], [0, 0, 2, 2], [2, 0, 2, 2], [0, 2, 0, 2]]},
                end:    super::Board{0: [[0, 0, 0, 0], [0, 0, 0, 0], [0, 0, 0, 2], [2, 4, 4, 4]]},
                direction: super::MoveDirection::Down,
            },
            MoveSample{
                start:  super::Board{0: [[2, 0, 0, 4], [2, 2, 4, 0], [2, 2, 2, 2], [2, 4, 2, 2]]},
                end:    super::Board{0: [[0, 0, 0, 0], [0, 0, 0, 0], [4, 4, 4, 4], [4, 4, 4, 4]]},
                direction: super::MoveDirection::Down,
            },
            MoveSample{
                start:  super::Board{0: [[0, 0, 0, 0], [2, 0, 0, 0], [2, 2, 0, 0], [2, 0, 2, 0]]},
                end:    super::Board{0: [[0, 0, 0, 0], [0, 0, 0, 2], [0, 0, 0, 4], [0, 0, 0, 4]]},
                direction: super::MoveDirection::Right,
            },
            MoveSample{
                start:  super::Board{0: [[0, 0, 2, 0], [2, 0, 0, 2], [0, 2, 2, 0], [0, 2, 2, 2]]},
                end:    super::Board{0: [[0, 0, 0, 2], [0, 0, 0, 4], [0, 0, 0, 4], [0, 0, 2, 4]]},
                direction: super::MoveDirection::Right,
            },
            MoveSample{
                start:  super::Board{0: [[2, 2, 2, 2], [0, 2, 2, 4], [0, 4, 2, 2], [4, 0, 2, 2]]},
                end:    super::Board{0: [[0, 0, 4, 4], [0, 0, 4, 4], [0, 0, 4, 4], [0, 0, 4, 4]]},
                direction: super::MoveDirection::Right,
            },
            MoveSample{
                start:  super::Board{0: [[0, 0, 0, 0], [2, 0, 0, 0], [2, 2, 0, 0], [2, 0, 2, 0]]},
                end:    super::Board{0: [[0, 0, 0, 0], [2, 0, 0, 0], [4, 0, 0, 0], [4, 0, 0, 0]]},
                direction: super::MoveDirection::Left,
            },
            MoveSample{
                start:  super::Board{0: [[0, 0, 2, 0], [2, 0, 0, 2], [0, 2, 2, 0], [0, 2, 2, 2]]},
                end:    super::Board{0: [[2, 0, 0, 0], [4, 0, 0, 0], [4, 0, 0, 0], [4, 2, 0, 0]]},
                direction: super::MoveDirection::Left,
            },
            MoveSample{
                start:  super::Board{0: [[2, 2, 2, 2], [0, 2, 2, 4], [0, 4, 2, 2], [4, 0, 2, 2]]},
                end:    super::Board{0: [[4, 4, 0, 0], [4, 4, 0, 0], [4, 4, 0, 0], [4, 4, 0, 0]]},
                direction: super::MoveDirection::Left,
            }
        ];

        for sample in (&mut samples).into_iter() {
            sample.start.move_(&sample.direction);
            assert_eq!(sample.start, sample.end, "{}", sample.direction.to_string());
        }
    }
}