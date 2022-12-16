mod game_engine;
mod player;


fn main() {
    let player = player::random::RandomPlayer{};
    let mut engine = game_engine::GameEngine::<player::random::RandomPlayer>::new(4, player);
    println!("{}", engine.to_string());
    println!("{}", engine.play());
}
