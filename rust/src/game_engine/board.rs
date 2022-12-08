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