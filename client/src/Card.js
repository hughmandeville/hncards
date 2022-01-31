import React from 'react';
import './Card.css';
import Comments from './Comments';
import { timeSince } from './utils';

// Card component
const Card = ({ item }) => {
  const { title, description, icon, publisher, time, image, comment_count, by, points, id, url } =
    item;
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
      {description && <div className="desc">{description}</div>}
      <div className="card-footer">
        by <a href={'https://news.ycombinator.com/user?id=' + by}>{by}</a>{' '}
        <div className="dot">&bull;</div> <a href={'https://news.ycombinator.com/item?id=' + id}>{points} pts</a>{' '}
        {comment_count > 0 && <Comments comment_count={comment_count} id={id} />}
      </div>
    </div>
  );
};

export default Card;
