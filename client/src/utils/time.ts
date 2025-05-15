export const getTimeElapsedSinceMatch = (
  gameStartTimestamp: number,
  maxTimePlayed: number,
) => {
  const gameEndTimeStamp = gameStartTimestamp + maxTimePlayed * 1000;

  const currentDate = Date.now();

  const timeSinceGame = currentDate - gameEndTimeStamp;

  return timeSinceGame;
};

export const getTimeSinceMatchString = (
  timeSinceGame: number,
  gameEndedTimeStamp: number,
): string => {
  let timeSinceGameString = "";

  const timeDifferenceSecs = Math.floor(timeSinceGame / 1000);
  const timeDifferenceMins = Math.floor(timeSinceGame / 60000);
  const timeDifferenceHours = Math.floor(timeSinceGame / 3600000);
  const timeDifferenceDays = Math.floor(timeSinceGame / 86400000);

  if (timeDifferenceSecs < 60) {
    timeSinceGameString = "just now";
  } else if (timeDifferenceMins < 60) {
    timeSinceGameString =
      timeDifferenceMins < 2
        ? "a minute ago"
        : `${timeDifferenceMins} minutes ago`;
  } else if (timeDifferenceHours < 24) {
    timeSinceGameString =
      timeDifferenceHours < 2
        ? "an hour ago"
        : `${timeDifferenceHours} hours ago`;
  } else if (timeDifferenceDays < 31) {
    timeSinceGameString =
      timeDifferenceDays < 2 ? "a day ago" : `${timeDifferenceDays} days ago`;
  } else {
    timeSinceGameString = `${new Date(
      gameEndedTimeStamp,
    ).toLocaleDateString()}`;
  }

  return timeSinceGameString;
};

export const getGameDurationString = (gameDuration: number) => {
  const gameDurationInMinutes = String(Math.floor(gameDuration / 60)).padStart(
    2,
    "0",
  );
  const remainingSeconds = String(gameDuration % 60).padStart(2, "0");

  return `${gameDurationInMinutes}:${remainingSeconds}`;
};
