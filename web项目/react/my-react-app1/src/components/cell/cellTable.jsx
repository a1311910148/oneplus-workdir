import CellBox from "./cellBox";

export default function ({ data, handleInput, onselect }) {
  let startEl = null;
  let stopEl = null;
  function handleMouseDown(e) {
    startEl = e.target;
  }
  function handleMouseUp(e) {
    console.log(e);
    stopEl = e.target;
    // if (stopEl === startEl) return false;
    let startxy = { x: startEl.dataset.x, y: startEl.dataset.y };
    let stopxy = { x: stopEl.dataset.x, y: stopEl.dataset.y };
    console.log(startxy, stopxy);
    onselect(startxy, stopxy);
  }
  return (
    <div onMouseDown={handleMouseDown} onMouseUp={handleMouseUp}>
      {data.map((el) => {
        return (
          <CellBox
            key={el.id}
            inputValue={el.content}
            id={el.id}
            select={el.select}
            x={el.x}
            y={el.y}
            width={el.width}
            onInput={handleInput}
          />
        );
      })}
    </div>
  );
}
