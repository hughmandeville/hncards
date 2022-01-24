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

    let desc = "";
    let img = "";
    let titleClass = "title";
    if (item.image !== "") {
      titleClass = "title-with-img";
      img = (
        <div>
          <div
            className="img"
            style={{
              backgroundImage: `url(${item.image})`,
            }}
          ></div>
        </div>
      );
    } else if (item.og_description !== "") {
      desc = <div className="desc">{item.og_description}</div>;
    }

    return (
      <div className="box" key={"item-" + item.id}>
        {img}
        <div className={titleClass}>
          <a href={item.url}>{title}</a>
        </div>
        {desc}
      </div>
    );
  });

  return <div className="boxes">{boxes}</div>;
}

export default ItemBoxes;
