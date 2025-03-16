export const calculateWinRate = (wins: number, losses: number) => {
  const totalGamesPlayed = wins + losses;

  const winRate = (wins / totalGamesPlayed) * 100;

  return winRate.toFixed(1);
};
