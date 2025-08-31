/**
 * @param {character[][]} board
 * @return {boolean}
 */
var isValidSudoku = function (board) {
  const rows = Array.from({ length: 9 }, () => new Set());
  const cols = Array.from({ length: 9 }, () => new Set());
  const squares = Array.from({ length: 3 }, () =>
    Array.from({ length: 3 }, () => new Set())
  );

  console.log(squares);
  for (let x = 0; x < board.length; x++) {
    for (let y = 0; y < board[x].length; y++) {
      const cell = board[x][y];
      if (cell === ".") {
        continue;
      }

      const xSquare = Math.floor(x / 3);
      const ySquare = Math.floor(y / 3);

      if (
        rows[x].has(cell) ||
        cols[y].has(cell) ||
        squares[xSquare][ySquare].has(cell)
      ) {
        return false;
      }

      rows[x].add(cell);
      cols[y].add(cell);
      squares[xSquare][ySquare].add(cell);
    }
  }

  return true;
};

console.log(
  isValidSudoku([
    ["5", "3", ".", ".", "7", ".", ".", ".", "."],
    ["6", ".", ".", "1", "9", "5", ".", ".", "."],
    [".", "9", "8", ".", ".", ".", ".", "6", "."],
    ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
    ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
    ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
    [".", "6", ".", ".", ".", ".", "2", "8", "."],
    [".", ".", ".", "4", "1", "9", ".", ".", "5"],
    [".", ".", ".", ".", "8", ".", ".", "7", "9"],
  ])
);
