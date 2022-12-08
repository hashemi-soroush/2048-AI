mod game_engine;


fn main() {
    let mut engine = game_engine::GameEngine::new();
    println!("{}", engine.to_string());
    
    engine.board.move_(&game_engine::MoveDirection::Right);
    println!("{}", engine.to_string());

    engine.board.move_(&game_engine::MoveDirection::Up);
    println!("{}", engine.to_string());

    engine.board.move_(&game_engine::MoveDirection::Down);
    println!("{}", engine.to_string());

    engine.board.move_(&game_engine::MoveDirection::Left);
    println!("{}", engine.to_string());
}
