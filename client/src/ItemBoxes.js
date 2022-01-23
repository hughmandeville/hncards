import "./ItemBoxes.css";

// ItemBoxews - Render item boxes.
function ItemBoxes(props) {
  if (!props.items) {
    return null;
  }
  const boxes = props.items.map((item) => {
    return (
      <div key={"item-" + item.id}>
        <a href={item.url}>{item.title}</a>
      </div>
    );
  });

  return <div>{boxes}</div>;
}

export default ItemBoxes;
