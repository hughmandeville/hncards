import React from 'react';
import './Comments.css';

// Comments component
const Comments = ({ comment_count, id }) => (
  <span>
    <a href={'https://news.ycombinator.com/item?id=' + id}>
      <button type="button" className="commentsButtonClass" aria-label={comment_count}>
        <i className="iconClass">
          <span className="textClass">{comment_count}</span>
        </i>
      </button>
    </a>
  </span>
);

export default Comments;
