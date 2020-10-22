import { createSelector } from '@reduxjs/toolkit';
import { selectGame } from '.';

export const selectTable = createSelector(selectGame, (g) => g.table);
export const selectMatch = createSelector(selectGame, (g) => g.match);
export const selectCurrentTableOrThrow = createSelector(selectTable, (t) => {
  if (t === null) {
    throw new Error('no current table');
  }
  return t;
});

export const selectCurrentMatchOrThrow = createSelector(selectMatch, (m) => {
  if (m === null) {
    throw new Error('no current match');
  }
  return m;
});

export const selectPlayers = createSelector(selectCurrentMatchOrThrow, (m) => m && m.players);
