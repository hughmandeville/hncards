import React from 'react';

// Comments component
const Comments = ({ kids, id }) => (
  <span>
    <div className="dot">&bull;</div>{' '}
    <a href={'https://news.ycombinator.com/item?id=' + id}>{kids.length} comments</a>
  </span>
);

export default Comments;
