import "./ItemBoxes.css";

// ItemBoxews - Render item boxes.
function ItemBoxes(props) {
  if (
    (undefined === props.items ||
      undefined === props.items.length ||
      props.items.length < 1) &&
    (undefined === props.error || props.error === "")
  ) {
    return null;
  }

  const boxes = props.items.map((item) => {
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
          {item.publisher} <div className="dot">&bull;</div> time
        </div>
      </div>
    );

    if (item.image !== "") {
      titleClass = "title-with-img";
      img = (
        <div
          className="img"
          style={{
            backgroundImage: `url(${item.image})`,
          }}
        ></div>
      );
    } else if (item.og_description !== "") {
      desc = <div className="desc">{item.og_description}</div>;
    }

    let comments = "";
    if (item.kids && item.kids.length > 0) {
      comments = (
        <span>
          <div className="dot">&bull;</div>{" "}
          <a href={"https://news.ycombinator.com/item?id=" + item.id}>
            {item.kids.length} cmts
          </a>
        </span>
      );
    }
    footer = (
      <div className="box-footer">
        by {item.by} <div className="dot">&bull;</div> {item.score} pts{" "}
        {comments}
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
