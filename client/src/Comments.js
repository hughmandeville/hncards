import React from 'react';
import './Comments.css';

// Comments component
const Comments = ({ kids, id }) => (
  <span>
    <div className="dot">&bull;</div>{' '}
    <a href={'https://news.ycombinator.com/item?id=' + id}>
      <button type="button" className="commentsButtonClass" aria-label={kids.length}>
        <i className="iconClass">
          <span className="textClass">{kids.length}</span>
        </i>
      </button>
    </a>
  </span>
);

export default Comments;
