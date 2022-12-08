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

    pub fn move_(&mut self, direction: MoveDirection) -> u32 {
        match direction {
            MoveDirection::Up => self.move_up(),
            MoveDirection::Down => self.move_down(),
            MoveDirection::Left => self.move_left(),
            MoveDirection::Right => self.move_right(),
        }
    }

    fn move_up(&mut self) -> u32 {
        let mut score: u32 = 0;
        for j in 0..4 {
            let mut cur = 0;
            for i in 0..4 {
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
        for j in 0..4 {
            let mut cur = 3;
            for i in 3..=0 {
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
        for i in 0..4 {
            let mut cur = 0;
            for j in 0..4 {
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
        for i in 0..4 {
            let mut cur = 3;
            for j in 3..=0 {
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