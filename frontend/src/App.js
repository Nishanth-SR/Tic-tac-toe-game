import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './App.css';
import Board from './Board';

function App() {
  const [game, setGame] = useState({ board: Array(9).fill(""), player: "X" });

  const startNewGame = () => {
    axios.get('http://localhost:8080/api/newgame').then(response => {
      setGame(response.data);
    });
  };

  useEffect(() => {
    startNewGame();
  }, []);

  const handleClick = (index) => {
    axios.post('http://localhost:8080/api/move', { index }).then(response => {
      setGame(response.data);
    });
  };

  return (
    <div className="App">
      <h1>Tic Tac Toe</h1>
      <Board board={game.board} onClick={handleClick} />
      <p>Next player: {game.player}</p>
      <button onClick={startNewGame}>Start New Game</button>
    </div>
  );
}

export default App;
