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

// Return time since in human readable format (e.g. "2 hrs", "on Sat").
function timeSince(ts) {
  var seconds = Math.floor((new Date() - ts * 1000) / 1000);

  var interval = seconds / 31536000;

  if (interval > 1) {
    if (Math.floor(interval) === 1) {
      return Math.floor(interval) + " yr";
    }
    return Math.floor(interval) + " yrs";
  }
  interval = seconds / 2592000;
  if (interval > 1) {
    if (Math.floor(interval) === 1) {
      return Math.floor(interval) + " mon";
    }
    return Math.floor(interval) + " mons";
  }
  interval = seconds / 86400;
  if (interval > 1) {
    // If 2 to 5 days ago, show day of week (e.g. "on Sat").
    if (Math.floor(interval) >= 2 && Math.floor(interval) <= 5) {
      const d = new Date(ts * 1000);
      return "on " + d.toLocaleDateString("en-US", { weekday: "short" });
    }
    return Math.floor(interval) + " d";
  }
  interval = seconds / 3600;
  if (interval > 1) {
    if (Math.floor(interval) === 1) {
      return "1 hr";
    }
    return Math.floor(interval) + " hrs";
  }
  interval = seconds / 60;
  if (interval > 1) {
    return Math.floor(interval) + " mins";
  }
  return Math.floor(seconds) + " secs";
}

export default ItemBoxes;
