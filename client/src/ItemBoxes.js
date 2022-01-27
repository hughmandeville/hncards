import React from 'react';
import "./ItemBoxes.css";
import { timeSince } from './utils';

// ItemBoxews - Render item boxes.
const ItemBoxes = ({ items }) => {

  if (!items?.length) return null;

  const boxes = items.map((item) => {
    let title = item.title;
    if (title === "") {
      title = item.og_title;
    }

    let pub = "";
    let pubIcon = "";
    let img = "";
    let desc = "";
    let footer = "";
    let titleClass = "title";
    if (item.icon) {
      pubIcon = (
        <img
          className="icon"
          src={item.icon}
          alt={item.publisher}
          onError={(event) => (event.target.style.display = "none")}
        />
      );
    }

    pub = (
      <div className="publisher">
        {pubIcon}
        <div className="publisher-text">
          {item.publisher} <div className="dot">&bull;</div>{" "}
          {timeSince(item.time)}
        </div>
      </div>
    );

    if (item.image !== "") {
      titleClass = "title-with-img";
      img = (
        <div
          className="img"
          style={{
            backgroundImage: `url("${item.image}")`,
          }}
        ></div>
      );
    } else if (item.og_description !== "") {
      desc = <div className="desc">{item.og_description}</div>;
    }

    let comments = "";
    if (item.kids && item.kids.length > 0) {
      // put comment count in a comments icon
      comments = (
        <span>
          <div className="dot">&bull;</div>{" "}
          <a href={"https://news.ycombinator.com/item?id=" + item.id}>
            {item.kids.length} comments
          </a>
        </span>
      );
    }
    footer = (
      <div className="box-footer">
        by{" "}
        <a href={"https://news.ycombinator.com/user?id=" + item.by}>
          {item.by}
        </a>{" "}
        <div className="dot">&bull;</div> {item.score} pts {comments}
      </div>
    );

    return (
      <div className="box" key={"item-" + item.id}>
        {img}
        {pub}
        <div className={titleClass}>
          <a href={item.url}>{title}</a>
        </div>
        {desc}
        {footer}
      </div>
    );
  });

  return <div className="boxes">{boxes}</div>;
}

export default ItemBoxes;
