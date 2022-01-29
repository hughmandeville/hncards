import React from 'react';
import './ItemBoxes.css';
import Card from './Card';

// ItemBoxews - Render item boxes.
const ItemBoxes = ({ items }) => {
  if (!items?.length) return null;

  return (
    <div className="boxes">
      {items.map(item => (
        <Card item={item} key={item.id} />
      ))}
    </div>
  );
};

export default ItemBoxes;
