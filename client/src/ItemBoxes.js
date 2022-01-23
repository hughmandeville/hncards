import "./ItemBoxes.css";

// ItemBoxews - Render item boxes.
function ItemBoxes(props) {
  if (!props.items) {
    return null;
  }
  console.log("items: ", props.items);
  const boxes = props.items.map((item) => {
    return (
      <div className="box" key={"item-" + item.id}>
        <div className="title">
          <a href={item.url}>{item.title}</a>
        </div>
      </div>
    );
  });

  return <div className="boxes">{boxes}</div>;
}

export default ItemBoxes;
