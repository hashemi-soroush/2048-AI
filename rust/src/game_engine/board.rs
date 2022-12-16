use std::cmp::PartialEq;
use rand::Rng;
use super::MoveDirection;


#[derive(Debug, PartialEq, Clone)]
pub struct Board {
    b: Vec<Vec<u32>>,
}

impl Board {
    pub fn new(size: u32) -> Self {
        let mut board = Board{
            b: Vec::<Vec<u32>>::new()
        };
        
        for _ in 0..size {
            let mut new_row = Vec::<u32>::new();
            for _ in 0..size {
                new_row.push(0u32);
            }
            board.b.push(new_row);
        }

        board.put_randomly(2, &[1, 2]);
        return board;
    }

    pub fn put_randomly(&mut self, mut count: u8, choices: &[u32]) {
        let mut rng = rand::thread_rng();
        while count > 0 {
            let mut i: usize = rng.gen(); i = i % self.b.len();
            let mut j: usize = rng.gen(); j = j % self.b[0].len();
            if self.b[i][j] > 0 {
                continue;
            }

            let mut val: u32 = rng.gen(); val = choices[val as usize % choices.len()];
            self.b[i][j] = val;
            count -= 1;
        }
    }

    pub fn can_move(&self) -> bool {
        let moves = [MoveDirection::Up, MoveDirection::Down, MoveDirection::Right, MoveDirection::Left];
        for move_ in moves {
            let mut cloned_board = self.clone();
            cloned_board.move_(&move_);
            if self != &cloned_board {
                return true;
            }
        }
        return false;
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
        for j in 0..self.b.len() {
            let mut cur = 0;
            for i in 0..self.b.len() {
                if i == cur || self.b[i][j] == 0 {
                    continue;
                }

                if self.b[cur][j] == 0 {
                    self.b[cur][j] = self.b[i][j];
                    self.b[i][j] = 0;
                } else if self.b[cur][j] == self.b[i][j] {
                    self.b[cur][j] *= 2;
                    self.b[i][j] = 0;
                    score += self.b[cur][j];
                    cur += 1;

                } else {
                    cur += 1;
                    let temp = self.b[i][j];
                    self.b[i][j] = 0;
                    self.b[cur][j] = temp;
                }
            }
        }
        return score;
    }

    fn move_down(&mut self) -> u32 {
        let mut score: u32 = 0;
        for j in 0..self.b[0].len() {
            let mut cur = self.b.len()-1;
            for i in (0..self.b.len()).rev() {
                if cur == i || self.b[i][j] == 0 {
                    continue;
                }

                if self.b[cur][j] == 0 {
                    self.b[cur][j] = self.b[i][j];
                    self.b[i][j] = 0;
                } else if self.b[i][j] == self.b[cur][j] {
                    self.b[cur][j] *= 2;
                    self.b[i][j] = 0;
                    score += self.b[cur][j];
                    cur -= 1;
                } else {
                    cur -= 1;
                    let temp = self.b[i][j];
                    self.b[i][j] = 0;
                    self.b[cur][j] = temp;
                }
            }
        }
        return score;
    }

    fn move_left(&mut self) -> u32 {
        let mut score: u32 = 0;
        for i in 0..self.b.len() {
            let mut cur = 0;
            for j in 0..self.b[i].len() {
                if self.b[i][j] == 0 || cur == j {
                    continue;
                }

                if self.b[i][cur] == 0 {
                    self.b[i][cur] = self.b[i][j];
                    self.b[i][j] = 0;
                } else if self.b[i][j] == self.b[i][cur] {
                    self.b[i][cur] *= 2;
                    self.b[i][j] = 0;
                    score += self.b[i][cur];
                    cur += 1;
                } else {
                    cur += 1;
                    let temp = self.b[i][j];
                    self.b[i][j] = 0;
                    self.b[i][cur] = temp;
                }
            }
        }
        return score;
    }
    
    fn move_right(&mut self) -> u32 {
        let mut score = 0;
        for i in 0..self.b.len() {
            let mut cur = self.b[0].len() - 1;
            for j in (0..self.b[0].len()).rev() {
                if self.b[i][j] == 0 || cur == j {
                    continue;
                }

                if self.b[i][cur] == 0 {
                    self.b[i][cur] = self.b[i][j];
                    self.b[i][j] = 0;
                } else if self.b[i][j] == self.b[i][cur] {
                    self.b[i][cur] *= 2;
                    self.b[i][j] = 0;
                    score += self.b[i][cur];
                    cur -= 1;
                } else {
                    cur -= 1;
                    let temp = self.b[i][j];
                    self.b[i][j] = 0;
                    self.b[i][cur] = temp;
                }
            }
        }
        return score;
    }
}

impl std::string::ToString for Board {
    fn to_string(&self) -> String {
        let mut res = String::from("Board {\n");
        for i in 0..self.b.len() {
            res += "\t";
            for j in 0..self.b[i].len() {
                res.push_str("\t");
                res.push_str(&self.b[i][j].to_string());
            }
            res.push_str("\n");
        }
        res.push_str("}");
        return res;
    }
}

// impl std::cmp::PartialEq for Board {
//     fn eq(&self, other: &Self) -> bool {
//         if self.b.len() != other.b.len() { return false; }
//         for i in 0..self.b.len() {
//             if self.b[i].len() != other.b[i].len() { return false; }

//             for j in 0..self.b[i].len() {
//                 if self.b[i][j] != other.b[i][j] { return false; }
//             }
//         }
//         return true;
//     }
// }

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
                start:  super::Board{b: vec![vec![0, 2, 2, 2], vec![0, 0, 2, 0], vec![0, 0, 0, 2], vec![0, 0, 0, 0]]},
                end:    super::Board{b: vec![vec![0, 2, 4, 4], vec![0, 0, 0, 0], vec![0, 0, 0, 0], vec![0, 0, 0, 0]]},
                direction: super::MoveDirection::Up,
            },
            MoveSample{
                start:  super::Board{b: vec![vec![0, 2, 0, 0], vec![0, 0, 2, 2], vec![2, 0, 2, 2], vec![0, 2, 0, 0]]},
                end:    super::Board{b: vec![vec![2, 4, 4, 4], vec![0, 0, 0, 0], vec![0, 0, 0, 0], vec![0, 0, 0, 0]]},
                direction: super::MoveDirection::Up,
            },
            MoveSample{
                start:  super::Board{b: vec![vec![2, 0, 0, 4], vec![2, 2, 4, 0], vec![2, 2, 2, 2], vec![2, 4, 2, 2]]},
                end:    super::Board{b: vec![vec![4, 4, 4, 4], vec![4, 4, 4, 4], vec![0, 0, 0, 0], vec![0, 0, 0, 0]]},
                direction: super::MoveDirection::Up,
            },
            MoveSample{
                start:  super::Board{b: vec![vec![0, 2, 2, 2], vec![0, 0, 2, 0], vec![0, 0, 0, 2], vec![0, 0, 0, 0]]},
                end:    super::Board{b: vec![vec![0, 0, 0, 0], vec![0, 0, 0, 0], vec![0, 0, 0, 0], vec![0, 2, 4, 4]]},
                direction: super::MoveDirection::Down,
            },
            MoveSample{
                start:  super::Board{b: vec![vec![0, 2, 0, 0], vec![0, 0, 2, 2], vec![2, 0, 2, 2], vec![0, 2, 0, 2]]},
                end:    super::Board{b: vec![vec![0, 0, 0, 0], vec![0, 0, 0, 0], vec![0, 0, 0, 2], vec![2, 4, 4, 4]]},
                direction: super::MoveDirection::Down,
            },
            MoveSample{
                start:  super::Board{b: vec![vec![2, 0, 0, 4], vec![2, 2, 4, 0], vec![2, 2, 2, 2], vec![2, 4, 2, 2]]},
                end:    super::Board{b: vec![vec![0, 0, 0, 0], vec![0, 0, 0, 0], vec![4, 4, 4, 4], vec![4, 4, 4, 4]]},
                direction: super::MoveDirection::Down,
            },
            MoveSample{
                start:  super::Board{b: vec![vec![0, 0, 0, 0], vec![2, 0, 0, 0], vec![2, 2, 0, 0], vec![2, 0, 2, 0]]},
                end:    super::Board{b: vec![vec![0, 0, 0, 0], vec![0, 0, 0, 2], vec![0, 0, 0, 4], vec![0, 0, 0, 4]]},
                direction: super::MoveDirection::Right,
            },
            MoveSample{
                start:  super::Board{b: vec![vec![0, 0, 2, 0], vec![2, 0, 0, 2], vec![0, 2, 2, 0], vec![0, 2, 2, 2]]},
                end:    super::Board{b: vec![vec![0, 0, 0, 2], vec![0, 0, 0, 4], vec![0, 0, 0, 4], vec![0, 0, 2, 4]]},
                direction: super::MoveDirection::Right,
            },
            MoveSample{
                start:  super::Board{b: vec![vec![2, 2, 2, 2], vec![0, 2, 2, 4], vec![0, 4, 2, 2], vec![4, 0, 2, 2]]},
                end:    super::Board{b: vec![vec![0, 0, 4, 4], vec![0, 0, 4, 4], vec![0, 0, 4, 4], vec![0, 0, 4, 4]]},
                direction: super::MoveDirection::Right,
            },
            MoveSample{
                start:  super::Board{b: vec![vec![0, 0, 0, 0], vec![2, 0, 0, 0], vec![2, 2, 0, 0], vec![2, 0, 2, 0]]},
                end:    super::Board{b: vec![vec![0, 0, 0, 0], vec![2, 0, 0, 0], vec![4, 0, 0, 0], vec![4, 0, 0, 0]]},
                direction: super::MoveDirection::Left,
            },
            MoveSample{
                start:  super::Board{b: vec![vec![0, 0, 2, 0], vec![2, 0, 0, 2], vec![0, 2, 2, 0], vec![0, 2, 2, 2]]},
                end:    super::Board{b: vec![vec![2, 0, 0, 0], vec![4, 0, 0, 0], vec![4, 0, 0, 0], vec![4, 2, 0, 0]]},
                direction: super::MoveDirection::Left,
            },
            MoveSample{
                start:  super::Board{b: vec![vec![2, 2, 2, 2], vec![0, 2, 2, 4], vec![0, 4, 2, 2], vec![4, 0, 2, 2]]},
                end:    super::Board{b: vec![vec![4, 4, 0, 0], vec![4, 4, 0, 0], vec![4, 4, 0, 0], vec![4, 4, 0, 0]]},
                direction: super::MoveDirection::Left,
            }
        ];

        for sample in (&mut samples).into_iter() {
            sample.start.move_(&sample.direction);
            assert_eq!(sample.start, sample.end, "{}", sample.direction.to_string());
        }
    }

    #[test]
    fn board_can_move() {
        let can_moves = [
            super::Board{b: vec![
                vec![2, 4, 2, 4],
                vec![4, 2, 4, 2],
                vec![2, 0, 2, 4],
                vec![4, 2, 4, 2],
            ]},
            super::Board{b: vec![
                vec![2, 4, 2, 4],
                vec![4, 2, 4, 2],
                vec![8, 8, 2, 4],
                vec![4, 2, 4, 2],
            ]},
            super::Board{b: vec![
                vec![2, 4, 2, 4],
                vec![4, 2, 4, 2],
                vec![8, 4, 2, 4],
                vec![8, 2, 4, 2],
            ]},
        ];

        let cannot_moves = [
            super::Board{b: vec![
                vec![2, 4, 2, 4],
                vec![4, 2, 4, 2],
                vec![2, 4, 2, 4],
                vec![4, 2, 4, 2],
            ]},
        ];

        for sample in can_moves {
            assert_eq!(sample.can_move(), true, "{}", sample.to_string());
        }

        for sample in cannot_moves {
            assert_eq!(sample.can_move(), false, "{}", sample.to_string());
        }
    }
}