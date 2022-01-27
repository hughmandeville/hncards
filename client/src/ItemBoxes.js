import React from 'react';
import './ItemBoxes.css';
import Card from './Card';
import { timeSince } from './utils';

// ItemBoxews - Render item boxes.
const ItemBoxes = ({ items }) => {
  if (!items?.length) return null;

  return (
    <div className="boxes">
      {items.map(item => (
        <Card item={item} />
      ))}
    </div>
  );
};

export default ItemBoxes;
