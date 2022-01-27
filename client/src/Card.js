import React from 'react';
import './Card.css';
import Comments from './Comments';
import { timeSince } from './utils';

// Card component
const Card = ({ item }) => {
  const {
    title: titleField,
    og_title,
    og_description,
    icon,
    publisher,
    time,
    image,
    kids,
    by,
    score,
    id,
    url,
  } = item;
  const title = titleField || og_title;
  const titleClass = image ? 'title-with-img' : 'title';

  return (
    <div className="card" key={'item-' + id}>
      {image && (
        <div
          className="img"
          style={{
            backgroundImage: `url("${image}")`,
          }}
        ></div>
      )}
      <div className="publisher">
        {icon && (
          <img
            className="icon"
            src={icon}
            alt={publisher}
            onError={event => (event.target.style.display = 'none')}
          />
        )}
        <div className="publisher-text">
          {publisher} <div className="dot">&bull;</div> {timeSince(time)}
        </div>
      </div>
      <div className={titleClass}>
        <a href={url}>{title}</a>
      </div>
      {og_description && <div className="desc">{og_description}</div>}
      <div className="card-footer">
        by <a href={'https://news.ycombinator.com/user?id=' + by}>{by}</a>{' '}
        <div className="dot">&bull;</div> {score} pts{' '}
        {kids?.length > 0 && <Comments kids={kids} id={id} />}
      </div>
    </div>
  );
};

export default Card;
