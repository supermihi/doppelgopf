import React from 'react';
import TrickView from './TrickView';
import { Diamond10, Diamond9, DiamondA, DiamondQ, SpadeA } from 'model/cards';
import { Pos } from 'model/players';

interface Props {
  playedCards: number;
}
const Trick = ({ playedCards }: Props) => (
  <div style={{ height: '400px' }}>
    <TrickView
      center={['50%', '50%']}
      cardWidth={150}
      trick={{
        forehand: Pos.left,
        cards: [Diamond9, Diamond10, DiamondA, DiamondQ].slice(0, playedCards),
      }}
    />
  </div>
);
export default {
  'one cards': <Trick playedCards={1} />,
  'two cards': <Trick playedCards={2} />,
  'three cards': <Trick playedCards={3} />,
  'four cards': <Trick playedCards={4} />,
};