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
    let title = item.og_title;
    if (title === "") {
      title = item.title;
    }

    if (item.image !== "") {
      return (
        <div className="box" key={"item-" + item.id}>
          <div
            className="img"
            style={{
              backgroundImage: `url(${item.image})`,
            }}
          ></div>
          <div className="title">
            <a href={item.url}>{title}</a>
          </div>
        </div>
      );
    }

    return (
      <div className="box" key={"item-" + item.id}>
        <div className="title2">
          <a href={item.url}>{title}</a>
        </div>
        <div className="desc">{item.og_description}</div>
      </div>
    );
  });

  return <div className="boxes">{boxes}</div>;
}

export default ItemBoxes;
