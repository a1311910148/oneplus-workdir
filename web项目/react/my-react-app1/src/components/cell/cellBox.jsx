export default function ({ id, x, y, width, select, inputValue, onInput }) {
  function handleInput(e) {
    onInput(id, e.target.value);
  }
  return (
    <>
      <input
        style={{ background: select ? "gray" : "" }}
        type="text"
        data-id={id}
        data-x={x}
        data-y={y}
        value={inputValue}
        onInput={handleInput}
      />
    </>
  );
}
