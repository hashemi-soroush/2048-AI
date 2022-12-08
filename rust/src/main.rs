mod game_engine;


fn main() {
    let engine = game_engine::GameEngine::new();
    println!("{}", engine.to_string())
}
