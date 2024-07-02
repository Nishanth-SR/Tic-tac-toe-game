import React from 'react';
import './Board.css';
import Square from './Square';

const Board = ({ board, onClick }) => (
  <div className="board">
    {board.map((value, index) => (
      <Square key={index} value={value} onClick={() => onClick(index)} />
    ))}
  </div>
);

export default Board;
